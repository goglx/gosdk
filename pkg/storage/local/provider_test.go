package local_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage/local"
)

var errInvalidFileName = errors.New("invalid file name")

type mockFileSystem struct {
	MkdirAllFunc func(path string, perm os.FileMode) error
	CreateFunc   func(name string) (*os.File, error)
}

func (fs *mockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return fs.MkdirAllFunc(path, perm)
}

func (fs *mockFileSystem) Create(name string) (*os.File, error) {
	return fs.CreateFunc(name)
}

type mockConfig struct {
	LocalPath string
}

func mockNewConfig() *mockConfig {
	return &mockConfig{
		LocalPath: "test",
	}
}

func mockNewProvider(config *mockConfig, fs *mockFileSystem) *local.Provider {
	return &local.Provider{Config: (*local.Config)(config), FS: fs}
}

func TestUpload(t *testing.T) {
	t.Parallel()

	mockFS := &mockFileSystem{
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			if path != "test" {
				t.Errorf("expected path %s, got %s", "expected/path", path)
			}

			return nil
		},

		CreateFunc: func(name string) (*os.File, error) {
			if name != "test/test-id" {
				return nil, fmt.Errorf("%w expected name test/test-id, got %s", errInvalidFileName, name)
			}

			return os.NewFile(1, name), nil
		},
	}

	provider := mockNewProvider(mockNewConfig(), mockFS)
	upload, err := provider.Upload(context.Background(), &types.File{
		ID:          "test-id",
		Data:        []byte("test-id"),
		ContentType: "text/plain",
	})

	sdktesting.IsNull(t, err)
	sdktesting.IsNotNull(t, upload)
	sdktesting.Equals(t, upload.ID, "test-id")
}
