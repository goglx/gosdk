package storage_test

import (
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage"
)

func TestNewProvider(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   types.ProviderType
		want   bool
		errMsg string
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
			name:   types.Local,
			want:   false,
			errMsg: "failed to create local provider: not implemented",
		},
		{
			name:   "wrong",
			want:   false,
			errMsg: "unsupported provider type",
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
				sdktesting.Ok(t, err.Error(), testCase.errMsg)
			}
		})
	}
}

func TestNewProviderManager(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   types.ProviderType
		want   bool
		errMsg string
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
			name:   types.Local,
			want:   false,
			errMsg: "failed to create local provider: not implemented",
		},
		{
			name:   "wrong",
			want:   false,
			errMsg: "unsupported provider type",
		},
	}

	for _, testCase := range tests {
		t.Run(string(testCase.name), func(t *testing.T) {
			t.Parallel()

			provider, err := storage.NewProviderManager(testCase.name)

			if testCase.want {
				sdktesting.IsNull(t, err)
				sdktesting.IsNotNull(t, provider)
			}

			if !testCase.want {
				sdktesting.IsNotNull(t, err)
				sdktesting.Ok(t, err.Error(), testCase.errMsg)
			}
		})
	}
}
