package storage_test

import (
	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Parallel()

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

	for _, testCase := range tests {
		t.Run(string(testCase.name), func(t *testing.T) {
			t.Parallel()

			provider, err := storage.NewProvider(testCase.name)

			if testCase.want {
				sdktesting.IsNull(t, err)
				sdktesting.IsNotNull(t, provider)
			}

			if !testCase.want {
				sdktesting.IsNotNull(t, err)
				sdktesting.Ok(t, err.Error(), "unsupported provider type")
			}
		})
	}
}
