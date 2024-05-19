package s3

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFile(key string, file *multipart.FileHeader) (string, error) {
	bucket := os.Getenv("AWS_S3_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		)),
	)
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("unable to open file, %v", err)
	}
	defer src.Close()

	// Upload the file
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   src,
	})
	if err != nil {
		return "", fmt.Errorf("unable to upload file, %v", err)
	}

	log.Println("Successfully uploaded to", bucket)

	// Construct the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)

	return url, nil
}
