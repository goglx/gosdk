package s3

import (
	"fmt"
	"os"
)

const s3URL = "https://%s.s3.%s.amazonaws.com/%s"

type Config struct {
	bucket    string
	accessKey string
	secretKey string
	region    string
}

func NewConfig() *Config {
	bucket := os.Getenv("BUCKET_NAME")
	if bucket == "" {
		panic("missing env BUCKET_NAME")
	}

	region := os.Getenv("S3_REGION")
	if region == "" {
		panic("missing env S3_REGION")
	}

	accessKey := os.Getenv("S3_ACCESS_KEY")
	if accessKey == "" {
		panic("missing env S3_ACCESS_KEY")
	}

	secretKey := os.Getenv("S3_SECRET_KEY")
	if secretKey == "" {
		panic("missing env S3_SECRET_KEY")
	}

	return &Config{
		bucket,
		region,
		accessKey,
		secretKey,
	}
}

func debug() string {
	return fmt.Sprintf(s3URL, "1", "2", "3")
}
