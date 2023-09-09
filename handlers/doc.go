package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	kagi "github.com/httpjamesm/kagigo"
)

const PDF_UPLOAD_BUCKET_NAME string = "coretext-pdfs-general"
const PDF_UPLOAD_ACCESS_POINT string = "coretext-pdfs-genera-meymo4pyxf87dr89ry53a3hzzercgusw1b-s3alias"

func UploadPDFDoc(sessionStore *session.Store, s3Client *s3.Client) fiber.Handler {
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

		// upload to s3
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(PDF_UPLOAD_ACCESS_POINT),
			Key:    aws.String(file.Filename), // TODO hash filename and store in databse?
			Body:   src,
		})
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON("File upload to S3 failed")
		}

		// // store info in session
		// session, err := sessionStore.Get(c)
		// if err != nil {
		// 	return err
		// }
		// // crete empty []models.PDFDocument{} is doens't exist
		// if session.Get("pdfDocuments") == nil {
		// 	session.Set("pdfDocuments", []models.PDFDocument{})
		// }
		// // create pdfdocument with data from this request
		// pdfdoc := models.PDFDocument{
		// 	Id:        uuid.NewString(),
		// 	Filename:  file.Filename,
		// 	Url:       presignedUrl.URL,
		// 	CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
		// }
		// // add it to array of pdfdocuments
		// pdfDocuments := session.Get("pdfDocuments").([]models.PDFDocument)
		// pdfDocuments = append(pdfDocuments, pdfdoc)
		// session.Set("pdfDocuments", pdfDocuments)

		// session.Save()

		// fmt.Println(session)
		// fmt.Println(session.Get("pdfDocuments"))

		log.Println("File uploaded successfully!")
		return c.Status(200).JSON("File uploaded successfully!")
	}

}

func SummarizePDF(s3PresignClient *s3.PresignClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// data validation & error handling todo here
		filename := c.Query("filename")
		presignedUrl := getPresignedUrl(filename, s3PresignClient)
		fmt.Println("filename: " + filename)
		fmt.Println("presign url: " + presignedUrl)

		// todo move this somewhere else (to app.go and pass in)
		kagiClient := kagi.NewClient(&kagi.ClientConfig{
			APIKey:     os.Getenv("KAGI_API_KEY"),
			APIVersion: "v0",
		})

		response, err := kagiClient.UniversalSummarizerCompletion(kagi.UniversalSummarizerParams{
			URL:         presignedUrl,
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
}

func getPresignedUrl(filename string, s3PresignClient *s3.PresignClient) string {
	presignedUrl, err := s3PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(PDF_UPLOAD_ACCESS_POINT),
		Key:    aws.String(filename),
	}, func(po *s3.PresignOptions) {
		po.Expires = 24 * time.Hour
	})
	if err != nil {
		log.Printf("Filed to generate pre-signed url: %v\n", err)
		return ""
	}
	return presignedUrl.URL
}
