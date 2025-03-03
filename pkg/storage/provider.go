package storage

import (
	"context"
	"errors"

	"gosdk/internal/types"
	"gosdk/pkg/storage/gcs"
	"gosdk/pkg/storage/local"
	"gosdk/pkg/storage/r2"
	"gosdk/pkg/storage/s3"
)

var errUnsupportedProviderType = errors.New("unsupported provider type")

type Provider struct {
	Upload   Upload
	Download Download
	Delete   Delete
	provider types.ProviderType
}

type Upload func(ctx context.Context, file *types.File) (*types.File, error)
type Download func(ctx context.Context, key string) ([]byte, error)
type Delete func(ctx context.Context, key string) error

func NewProvider(providerType types.ProviderType) (*Provider, error) {
	switch providerType {
	case types.S3:
		return &Provider{
			Upload:   s3.Upload,
			Download: s3.Download,
			Delete:   s3.Delete,
			provider: providerType,
		}, nil

	case types.R2:
		return &Provider{
			Upload:   r2.Upload,
			Download: r2.Download,
			Delete:   r2.Delete,
			provider: providerType,
		}, nil

	case types.GCS:
		return &Provider{
			Upload:   gcs.Upload,
			Download: gcs.Download,
			Delete:   gcs.Delete,
			provider: providerType,
		}, nil

	case types.Local:
		return &Provider{
			Upload:   local.Upload,
			Download: local.Download,
			Delete:   local.Delete,
			provider: providerType,
		}, nil

	default:
		return nil, errUnsupportedProviderType
	}
}
