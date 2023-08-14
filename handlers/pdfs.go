package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const PDF_UPLOAD_BUCKET_NAME string = "coretext-pdfs-general"
const PDF_UPLOAD_ACCESS_POINT string = "coretext-pdfs-genera-meymo4pyxf87dr89ry53a3hzzercgusw1b-s3alias"

func UploadDoc(c *fiber.Ctx) error {

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
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON("AWS configuration failed")
		log.Fatalf("failed to load configuration, %v", err)
	}
	client := s3.NewFromConfig(cfg)

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

	fmt.Println("File uploaded successfully!")
	log.Println("File uploaded successfully!")
	return c.JSON("File uploaded successfully!")
}
