package storage

import (
	"context"
	"fmt"
	"gosdk/internal/types"
	"gosdk/pkg/storage/gcs"
	"gosdk/pkg/storage/local"
	"gosdk/pkg/storage/r2"
	"gosdk/pkg/storage/s3"
)

type ProviderType string

const (
	S3    ProviderType = "s3"
	R2    ProviderType = "r2"
	GCS   ProviderType = "gcs"
	Local ProviderType = "local"
)

type Provider interface {
	Upload(ctx context.Context, file *types.File) (*types.File, error)
	Download(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}

func NewProvider(providerType ProviderType) (Provider, error) {
	switch providerType {
	case S3:
		return s3.NewProvider()
	case R2:
		return r2.NewProvider()
	case GCS:
		return gcs.NewProvider()
	case Local:
		return local.NewProvider()
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", providerType)
	}
}
