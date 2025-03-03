package local

import (
	"context"
	"errors"
	"gosdk/internal/types"
	"os"
)

var errMissingLocalPath = errors.New("missing env LOCAL_PATH")

type Provider struct {
	basePath string
}

func NewProvider() (*Provider, error) {
	if os.Getenv("LOCAL_PATH") == "" {
		return nil, errMissingLocalPath
	}

	return &Provider{basePath: os.Getenv("LOCAL_PATH")}, nil
}

func (p *Provider) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	// TODO implement me
	panic("implement me")
}

func (p *Provider) Download(ctx context.Context, key string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (p *Provider) Delete(ctx context.Context, key string) error {
	// TODO implement me
	panic("implement me")
}
