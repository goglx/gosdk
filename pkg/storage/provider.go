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

type ProviderManager struct {
	provider Provider
}

func NewProviderManager(providerType types.ProviderType) (*ProviderManager, error) {
	var provider Provider

	var err error

	switch providerType {
	case types.S3:
		provider, err = s3.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create s3 provider: %w", err)
		}

	case types.R2:
		provider, err = r2.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create r2 provider: %w", err)
		}

	case types.GCS:
		provider, err = gcs.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create gcs provider: %w", err)
		}

	case types.Local:
		provider, err = local.NewProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to create local provider: %w", err)
		}

	default:
		return nil, errUnsupportedProviderType
	}

	return &ProviderManager{provider: provider}, nil
}

func (pm *ProviderManager) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	uploadedFile, err := pm.provider.Upload(ctx, file)
	if err != nil {
		return nil, fmt.Errorf("provider manager: failed to upload file: %w", err)
	}

	return uploadedFile, nil
}
