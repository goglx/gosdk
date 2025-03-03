package local

import (
	"context"
	"errors"

	"gosdk/internal/types"
)

var errNotImplemented = errors.New("not implemented")

type Provider struct{}

func NewProvider() (*Provider, error) {
	return nil, errNotImplemented
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
