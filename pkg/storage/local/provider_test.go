package local_test

import (
	sdktesting "gosdk/internal/testing"
	"gosdk/pkg/storage/local"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Parallel()

	_, err := local.NewProvider()
	sdktesting.IsNotNull(t, err)
}
