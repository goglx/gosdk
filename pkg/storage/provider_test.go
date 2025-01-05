package storage_test

import (
	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Run("S3", func(t *testing.T) {
		provider, err := storage.NewProvider(types.S3)

		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, provider)
	})
	t.Run("r2", func(t *testing.T) {
		provider, err := storage.NewProvider(types.R2)

		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, provider)
	})
	t.Run("gcs", func(t *testing.T) {
		provider, err := storage.NewProvider(types.GCS)

		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, provider)
	})
	t.Run("local", func(t *testing.T) {
		provider, err := storage.NewProvider(types.Local)

		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, provider)
	})
	t.Run("invalid", func(t *testing.T) {
		_, err := storage.NewProvider("wrong")
		sdktesting.IsNotNull(t, err)
		sdktesting.Ok(t, err.Error(), "unsupported provider type: wrong")
	})
}
