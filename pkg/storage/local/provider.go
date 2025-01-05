package local

import (
	"context"
	"gosdk/internal/types"
)

type Provider struct{}

func NewProvider() (*Provider, error) {
	return &Provider{}, nil
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
