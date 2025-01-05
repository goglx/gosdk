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

type Provider interface {
	Upload(ctx context.Context, file *types.File) (*types.File, error)
	Download(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}

func NewProvider(providerType types.ProviderType) (Provider, error) {
	switch providerType {
	case types.S3:
		return s3.NewProvider()
	case types.R2:
		return r2.NewProvider()
	case types.GCS:
		return gcs.NewProvider()
	case types.Local:
		return local.NewProvider()
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", providerType)
	}
}
