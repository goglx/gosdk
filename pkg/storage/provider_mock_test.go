package storage_test

import (
	"context"

	"gosdk/internal/types"
	"gosdk/pkg/storage"
)

type mockProvider struct {
	mockUpload   func(ctx context.Context, file *types.File) (*types.File, error)
	mockDownload func(ctx context.Context, key string) ([]byte, error)
	mockDelete   func(ctx context.Context, key string) error
}

func (mp *mockProvider) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	return mp.mockUpload(ctx, file)
}

func (mp *mockProvider) Download(ctx context.Context, key string) ([]byte, error) {
	return mp.mockDownload(ctx, key)
}

func (mp *mockProvider) Delete(ctx context.Context, key string) error {
	return mp.mockDelete(ctx, key)
}

// NewMock is a helper function to create a Manager with a mock Provider.
func newMock(mockProvider storage.Provider) *storage.Manager {
	return &storage.Manager{Provider: mockProvider}
}
