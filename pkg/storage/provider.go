package storage

import (
	"context"
)

type ProviderType string

const (
	S3    ProviderType = "s3"
	R2    ProviderType = "r2"
	GCS   ProviderType = "gcs"
	Local ProviderType = "local"
)

type Provider interface {
	Upload(ctx context.Context, file *File) (*File, error)
	Download(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}

type File struct {
	Id          string `json:"id"`
	ContentType string `json:"content_type"`
	Data        []byte `json:"data"`
}
