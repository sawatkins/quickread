package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	kagi "github.com/httpjamesm/kagigo"
	"github.com/lithammer/shortuuid/v3"
)

const PDF_UPLOAD_ACCESS_POINT string = "quickread-pdf-upload-qa4gp8jcnk5orz4qtrnpmorqd9ogrusw1a-s3alias"

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
		key := shortuuid.New() + ".pdf"
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

func SummarizeDoc(s3PresignClient *s3.PresignClient, kagiClient *kagi.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// data validation & error handling todo here
		presignedUrl := c.Query("presignedUrl")

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
		// fmt.Println(response.Data.Output)
		return c.Status(200).JSON(response.Data.Output)
	}
}

func getPresignedUrl(key string, s3PresignClient *s3.PresignClient) string {
	presignedUrl, err := s3PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(PDF_UPLOAD_ACCESS_POINT),
		Key:    aws.String(key),
	}, func(po *s3.PresignOptions) {
		po.Expires = 24 * time.Hour
	})
	if err != nil {
		log.Printf("Filed to generate pre-signed url: %v\n", err)
		return ""
	}
	// for testing...
	//testurl := "https://coretext-pdfs-genera-meymo4pyxf87dr89ry53a3hzzercgusw1b-s3alias.s3.us-west-1.amazonaws.com/jE92nGcDGtTPMmpe2Qb64c.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIA4HM6FL2ZIUNVZEV3%2F20230920%2Fus-west-1%2Fs3%2Faws4_request&X-Amz-Date=20230920T061735Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=9c723f3fd240417f9239af83c0a026af0786448771f0b111851d282cc9059018"
	//fmt.Println(presignedUrl.Method)
	//return testurl
	return presignedUrl.URL
}
