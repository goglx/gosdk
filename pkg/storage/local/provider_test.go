package local_test

import (
	sdktesting "gosdk/internal/testing"
	"gosdk/pkg/storage/local"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		localProvider, err := local.NewProvider()
		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, localProvider)
	})

	t.Run("failed", func(t *testing.T) {
		_, err := local.NewProvider()
		sdktesting.IsNotNull(t, err)
	})
}
