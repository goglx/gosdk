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

type provider interface {
	Upload(ctx context.Context, file *types.File) (*types.File, error)
	Download(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}

type Manager struct {
	provider provider
}

func New(providerType types.ProviderType) (*Manager, error) {
	var provider provider

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

	return &Manager{provider: provider}, nil
}

func (pm *Manager) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	result, err := pm.provider.Upload(ctx, file)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file %s: %w", file.ID, err)
	}

	return result, nil
}

func (pm *Manager) Download(ctx context.Context, key string) ([]byte, error) {
	result, err := pm.provider.Download(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("failed to download file %s: %w", key, err)
	}

	return result, nil
}

func (pm *Manager) Delete(ctx context.Context, key string) error {
	err := pm.provider.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", key, err)
	}

	return nil
}
