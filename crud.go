package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Set up a new AWS session with your credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		// Credentials: credentials.NewStaticCredentials("ACCESS_KEY_ID", "SECRET_ACCESS_KEY", ""),
	})
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		os.Exit(1)
	}

	// Create a new DynamoDB client
	db := dynamodb.New(sess)

	// Create a new S3 client
	s3svc := s3.New(sess)

	// Define the name of the DynamoDB table and S3 bucket you want to use
	tableName := "test"
	bucketName := "service-catalog-pipeline-demo"

	// Define the item you want to insert into DynamoDB
	item := map[string]*dynamodb.AttributeValue{
		"Year":    {S: aws.String("12")},
		"Title":  {S: aws.String("John")},
		//"email": {S: aws.String("john.doe@example.com")},
	}

	// Insert the item into DynamoDB
	_, err = db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		fmt.Println("Failed to insert item into DynamoDB:", err)
		os.Exit(1)
	}

	// Define the path of the file you want to upload to S3
	filePath := "file.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Upload the file to S3
	_, err = s3svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key: aws.String("file.txt"),
		Body: file,
	})
	if err != nil {
		fmt.Println("Failed to upload file to S3:", err)
		os.Exit(1)
	}

	fmt.Println("Item inserted into DynamoDB and file uploaded to S3 successfully!")
}
