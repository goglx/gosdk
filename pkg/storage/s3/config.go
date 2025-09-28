package s3

import (
	"errors"
	"fmt"
	"os"
)

var ErrMissingEnv = errors.New("missing env")

const s3URL = "https://%s.s3.%s.amazonaws.com/%s"

type Config struct {
	bucket    string
	accessKey string
	secretKey string
	region    string
}

func NewConfig() (*Config, error) {
	bucket, err := getEnv("BUCKET_NAME")
	if err != nil {
		return nil, err
	}

	region, err := getEnv("S3_REGION")
	if err != nil {
		return nil, err
	}

	accessKey, err := getEnv("S3_ACCESS_KEY")
	if err != nil {
		return nil, err
	}

	secretKey, err := getEnv("S3_SECRET_KEY")
	if err != nil {
		return nil, err
	}

	debug()

	return &Config{
		bucket,
		accessKey,
		secretKey,
		region,
	}, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%w: %s", ErrMissingEnv, key)
	}

	return value, nil
}

func debug() string {
	return fmt.Sprintf(s3URL, "1", "2", "3")
}
