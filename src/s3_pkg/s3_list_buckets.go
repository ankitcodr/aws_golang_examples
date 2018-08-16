package s3_pkg

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	logger = log.New(os.Stdout, "S3 List Buckets: ", log.LstdFlags|log.Lshortfile)
)

type S3Client struct {
	Client *s3.S3
}

func (c *S3Client) List_buckets() {
	result, err := c.Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		logger.Println("Failed to list buckets", err)
		return
	}

	for _, bucket := range result.Buckets {
		logger.Printf("%s\n", aws.StringValue(bucket.Name))
	}
}
