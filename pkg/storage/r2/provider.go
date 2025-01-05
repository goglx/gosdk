package r2

import (
	"context"
	"fmt"
	"gosdk/internal/types"
	"os"
)

type provider struct {
	bucket    string
	accessKey string
	secretKey string
	region    string
}

func NewProvider() (*provider, error) {
	return &provider{
		bucket:    os.Getenv("BUCKET_NAME"),
		region:    os.Getenv("S3_REGION"),
		accessKey: os.Getenv("S3_ACCESS_KEY"),
		secretKey: os.Getenv("S3_SECRET_KEY"),
	}, nil
}

func (p *provider) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	//TODO implement me
	fmt.Println(file)
	panic("implement me")
}

func (p *provider) Download(ctx context.Context, key string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p *provider) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
