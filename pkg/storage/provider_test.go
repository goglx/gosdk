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
		want bool
	}{
		{
			name: types.S3,
			want: true,
		},
		{
			name: types.R2,
			want: true,
		},
		{
			name: types.GCS,
			want: true,
		},
		{
			name: types.Local,
			want: true,
		},
		{
			name: types.Local,
			want: true,
		},
		{
			name: "wrong",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			provider, err := storage.NewProvider(tt.name)

			if tt.want {
				sdktesting.IsNull(t, err)
				sdktesting.IsNotNull(t, provider)
			}

			if !tt.want {
				sdktesting.IsNotNull(t, err)
				sdktesting.Ok(t, err.Error(), "unsupported provider type: wrong")
			}
		})
	}
}
