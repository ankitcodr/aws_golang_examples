package main

import (
	"com/angupt/aws/examples/s3_pkg"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	s3Object := &s3_pkg.S3Client{
		Client: s3.New(session.New()),
	}

	s3Object.List_buckets()
}
