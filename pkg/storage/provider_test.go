package storage_test

import (
	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage"
	"testing"
)

func TestNewProvider(t *testing.T) {
	tests := []struct {
		name types.ProviderType
	}{
		{
			name: types.S3,
		},
		{
			name: types.R2,
		},
		{
			name: types.GCS,
		},
		{
			name: types.Local,
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			provider, err := storage.NewProvider(tt.name)
			sdktesting.IsNull(t, err)
			sdktesting.IsNotNull(t, provider)
		})
	}

	t.Run("invalid", func(t *testing.T) {
		_, err := storage.NewProvider("wrong")
		sdktesting.IsNotNull(t, err)
		sdktesting.Ok(t, err.Error(), "unsupported provider type: wrong")
	})
}
