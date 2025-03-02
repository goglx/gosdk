package local

import (
	"context"
	"fmt"

	"gosdk/internal/types"
)

type Provider struct{}

func NewProvider() (*Provider, error) {
	return nil, fmt.Errorf("not implemented")
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
