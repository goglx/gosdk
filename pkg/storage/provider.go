package storage

import (
	"context"
	"errors"
	"fmt"
	"gosdk/internal/types"
	"gosdk/pkg/storage/gcs"
	"gosdk/pkg/storage/local"
	"gosdk/pkg/storage/r2"
	"gosdk/pkg/storage/s3"
)

var errUnsupportedProviderType = errors.New("unsupported provider type")

type Provider interface {
	Upload(ctx context.Context, file *types.File) (*types.File, error)
	Download(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}

func NewProvider(providerType types.ProviderType) (Provider, error) {
	switch providerType {
	case types.S3:
		provider, err := s3.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create s3 provider: %w", err)
		}

		return provider, nil
	case types.R2:
		provider, err := r2.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create r2 provider: %w", err)
		}

		return provider, nil
	case types.GCS:
		provider, err := gcs.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create gcs provider: %w", err)
		}

		return provider, nil
	case types.Local:
		provider, err := local.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create local provider: %w", err)
		}

		return provider, nil
	default:
		return nil, errUnsupportedProviderType
	}
}
