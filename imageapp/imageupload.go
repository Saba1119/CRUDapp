package main

import (
	"fmt"
	//"io"
	"net/http"
	//"os"
	//"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "<html><body>"+
				"<form method='post' enctype='multipart/form-data'>"+
				"<input type='file' name='file' id='file'/>"+
				"<input type='submit' value='Upload'/>"+
				"</form></body></html>")
		} else {
			// Parse the multipart form
			err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Get the file from the form
			file, handler, err := r.FormFile("file")
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Create a new S3 session
			sess, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-1"), // Replace with your S3 region
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Upload the file to S3
			s3Svc := s3.New(sess)
			_, err = s3Svc.PutObject(&s3.PutObjectInput{
				Bucket: aws.String("sdk-bucket-demo"), // Replace with your S3 bucket name
				Key:    aws.String(handler.Filename),
				Body:   file,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "File uploaded successfully: %v", handler.Filename)
		}
	})

	// Listen on port 8080
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
