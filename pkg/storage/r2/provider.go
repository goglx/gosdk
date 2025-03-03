package r2

import (
	"context"
	"os"

	"gosdk/internal/types"
)

type Provider struct {
	bucket    string
	accessKey string
	secretKey string
	region    string
}

func NewProvider() (*Provider, error) {
	return &Provider{
		bucket:    os.Getenv("BUCKET_NAME"),
		region:    os.Getenv("S3_REGION"),
		accessKey: os.Getenv("S3_ACCESS_KEY"),
		secretKey: os.Getenv("S3_SECRET_KEY"),
	}, nil
}

func Upload(ctx context.Context, file *types.File) (*types.File, error) {
	// TODO implement me
	panic("implement me")
}

func Download(ctx context.Context, key string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func Delete(ctx context.Context, key string) error {
	// TODO implement me
	panic("implement me")
}
