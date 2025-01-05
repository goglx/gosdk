package local

import (
	"context"
	"fmt"
	"gosdk/internal/types"
)

type provider struct{}

func NewProvider() (*provider, error) {
	return &provider{}, nil
}

func (p *provider) Upload(ctx context.Context, file *types.File) (*types.File, error) {
	//TODO implement me
	fmt.Println(file)
	panic("implement me")
}

func (p *provider) Download(ctx context.Context, key string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p *provider) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
