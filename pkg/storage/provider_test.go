package storage_test

import (
	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage"
	"testing"
)

func TestNewProvider_S3(t *testing.T) {
	t.Parallel()
	provider, err := storage.NewProvider(types.S3)

	sdktesting.IsNull(t, err)
	sdktesting.IsNotNull(t, provider)
}

func TestNewProvider_invalidProvider(t *testing.T) {
	t.Parallel()
	_, err := storage.NewProvider("wrong")
	sdktesting.IsNotNull(t, err)
	sdktesting.Ok(t, err.Error(), "unsupported provider type: wrong")
}
