package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	kagi "github.com/httpjamesm/kagigo"
	"github.com/lithammer/shortuuid/v3"
)

const PDF_UPLOAD_BUCKET_NAME string = "coretext-pdfs-general"
const PDF_UPLOAD_ACCESS_POINT string = "coretext-pdfs-genera-meymo4pyxf87dr89ry53a3hzzercgusw1b-s3alias"

func UploadDoc(s3Client *s3.Client, s3PresignClient *s3.PresignClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// upload size and filetype validation
		file, err := c.FormFile("input-upload-doc")
		if err != nil {
			c.Status(500).JSON("Filaure to select document from form")
		}
		src, err := file.Open()
		if err != nil {
			c.Status(500).JSON("File to open file")
		}
		defer src.Close()

		key := shortuuid.New()
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(PDF_UPLOAD_ACCESS_POINT),
			Key:    aws.String(key),
			Body:   src,
		})
		if err != nil {
			fmt.Println(err)
			return c.Status(500).JSON("File upload to S3 failed")
		}

		presignedUrl := getPresignedUrl(key, s3PresignClient)
		if presignedUrl == "" {
			return c.SendStatus(500)
		}

		log.Println("File uploaded successfully!")
		return c.Status(200).JSON(fiber.Map{
			"presignedUrl": presignedUrl,
		})

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

func getPresignedUrl(key string, s3PresignClient *s3.PresignClient) string {
	presignedUrl, err := s3PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(PDF_UPLOAD_ACCESS_POINT),
		Key:    aws.String(key),
	}, func(po *s3.PresignOptions) {
		po.Expires = 2 * time.Hour
	})
	if err != nil {
		log.Printf("Filed to generate pre-signed url: %v\n", err)
		return ""
	}
	return presignedUrl.URL
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
