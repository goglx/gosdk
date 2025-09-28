package s3

import (
	"context"
	"errors"

	"gosdk/internal/types"
)

type Provider struct {
	config *Config
}

func NewProvider() (*Provider, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	return &Provider{
		config: config,
	}, nil
}

var ErrNotImplemented = errors.New("not implemented")

// ... (rest of the file is the same)

func (p *Provider) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	// TODO implement me
	return nil, ErrNotImplemented
}

func (p *Provider) Download(ctx context.Context, key string) ([]byte, error) {
	// TODO implement me
	return nil, ErrNotImplemented
}

func (p *Provider) Delete(ctx context.Context, key string) error {
	// TODO implement me
	return ErrNotImplemented
}
