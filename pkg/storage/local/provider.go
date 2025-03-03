package local

import (
	"context"
	"fmt"
	"os"

	"gosdk/internal/types"
)

type Provider struct {
	Config *Config
	FS     fileSystem
}

type fileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
	Create(name string) (*os.File, error)
}

type realFileSystem struct{}

func NewProvider() (*Provider, error) {
	config := NewConfig()

	return &Provider{Config: config, FS: &realFileSystem{}}, nil
}

func (fs *realFileSystem) MkdirAll(path string, perm os.FileMode) error {
	err := os.MkdirAll(path, perm)
	if err != nil {
		return fmt.Errorf("local storage provider, %w", err)
	}

	return nil
}

func (fs *realFileSystem) Create(name string) (*os.File, error) {
	create, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("local storage provider, %w", err)
	}

	return create, nil
}

func (p *Provider) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	err := p.FS.MkdirAll(p.Config.LocalPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("error creating directory, %w", err)
	}

	localFile, err := p.FS.Create(fmt.Sprintf("%s/%s", p.Config.LocalPath, file.ID))
	if err != nil {
		return nil, fmt.Errorf("error creating file, %w", err)
	}

	defer localFile.Close()

	_, err = localFile.Write(file.Data)

	if err != nil {
		return nil, fmt.Errorf("error writing to file, %w", err)
	}

	return file, nil
}

func (p *Provider) Download(ctx context.Context, key string) ([]byte, error) {
	return []byte(key), nil
}

func (p *Provider) Delete(ctx context.Context, key string) error {
	return nil
}
