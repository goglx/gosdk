package local_test

import (
	"os"

	"gosdk/pkg/storage/local"
)

type mockFileSystem struct {
	mkdirAllFunc func(path string, perm os.FileMode) error
	createFunc   func(name string) (*os.File, error)
}

type mockConfig struct {
	LocalPath string
}

func (fs *mockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return fs.mkdirAllFunc(path, perm)
}

func (fs *mockFileSystem) Create(name string) (*os.File, error) {
	return fs.createFunc(name)
}

func newMockConfig() *mockConfig {
	return &mockConfig{
		LocalPath: "test",
	}
}

func newMockProvider(config *mockConfig, fs *mockFileSystem) *local.Provider {
	return &local.Provider{Config: (*local.Config)(config), FS: fs}
}
