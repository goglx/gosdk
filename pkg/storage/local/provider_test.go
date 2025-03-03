package local_test

import (
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/pkg/storage/local"
)

func TestNewProviderSuccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Setenv("LOCAL_PATH", "/tmp")

		localProvider, err := local.NewProvider()
		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, localProvider)
	})
}

func TestNewProvider(t *testing.T) {
	t.Parallel()

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		_, err := local.NewProvider()
		sdktesting.IsNotNull(t, err)
		sdktesting.Equals(t, err.Error(), "missing env LOCAL_PATH")
	})
}
