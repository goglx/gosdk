package local

import (
	"context"
	"fmt"
	"gosdk/internal/types"
	"os"
)

type Provider struct {
	Config *Config
	FS     FileSystem
}

type FileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
	Create(name string) (*os.File, error)
}

type RealFileSystem struct{}

func (fs *RealFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs *RealFileSystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func NewRealFileSystem() *RealFileSystem {
	return &RealFileSystem{}
}

func NewProvider(fs FileSystem) (*Provider, error) {
	config := NewConfig()
	return &Provider{Config: config, FS: fs}, nil
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
