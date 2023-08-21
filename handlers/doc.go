package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sawatkins/quickread/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	kagi "github.com/httpjamesm/kagigo"
)

const PDF_UPLOAD_BUCKET_NAME string = "coretext-pdfs-general"
const PDF_UPLOAD_ACCESS_POINT string = "coretext-pdfs-genera-meymo4pyxf87dr89ry53a3hzzercgusw1b-s3alias"

func UploadPDFDoc(sessionStore *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// upload size validation
		// pdf filetype validation

		file, err := c.FormFile("input-upload-doc")
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON("File upload failed")
		}

		src, err := file.Open()
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON("File open failed")
		}
		defer src.Close()

		// TODO make these in a seperate function so i can just call once to get the cfg (or client)
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON("AWS configuration failed")
			log.Printf("failed to load configuration, %v", err)
		}
		client := s3.NewFromConfig(cfg)
		presignClient := s3.NewPresignClient(client)

		// upload to s3
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(PDF_UPLOAD_ACCESS_POINT),
			Key:    aws.String(file.Filename), // TODO hash filename and store in databse?
			Body:   src,
		})
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON("File upload to S3 failed")
		}

		// get presign url
		presignedUrl, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(PDF_UPLOAD_ACCESS_POINT),
			Key:    aws.String(file.Filename),
		}, func(po *s3.PresignOptions) {
			po.Expires = 24 * time.Hour
		})
		if err != nil {
			log.Printf("Filed to generate pre-signed url: %v\n", err)
			return c.Status(500).JSON("Filed to generate pre-signed url")

		}
		fmt.Println("presigned url generated successfully")
		fmt.Println(presignedUrl.URL)

		// store info in session
		session, err := sessionStore.Get(c)
		if err != nil {
			return err
		}
		// crete empty []models.PDFDocument{} is doens't exist
		if session.Get("pdfDocuments") == nil {
			session.Set("pdfDocuments", []models.PDFDocument{})
		}
		// create pdfdocument with data from this request
		pdfdoc := models.PDFDocument{
			Id:        uuid.NewString(),
			Filename:  file.Filename,
			Url:       presignedUrl.URL,
			CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
		}
		// add it to array of pdfdocuments
		pdfDocuments := session.Get("pdfDocuments").([]models.PDFDocument)
		pdfDocuments = append(pdfDocuments, pdfdoc)
		session.Set("pdfDocuments", pdfDocuments)

		session.Save()

		log.Println("File uploaded successfully!")
		return c.JSON("File uploaded successfully!")
	}

}

func SummarizePDF(c *fiber.Ctx) error {
	// data validation & error handling todo here
	url := c.Query("url")

	// todo move this somewhere else
	kagiClient := kagi.NewClient(&kagi.ClientConfig{
		APIKey:     os.Getenv("KAGI_API_KEY"),
		APIVersion: "v0",
	})

	response, err := kagiClient.UniversalSummarizerCompletion(kagi.UniversalSummarizerParams{
		URL:         url,
		SummaryType: kagi.SummaryTypeSummary,
		Engine:      kagi.SummaryEngineCecil,
	})
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON("Summarize PDF failed")
	}
	// todo log the meta data response
	fmt.Println(response.Data.Output)
	return c.Status(200).JSON(response.Data.Output)
}
