package local_test

import (
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/pkg/storage/local"
)

func TestITNewProvider(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Setenv("LOCAL_PATH", "/tmp")

		localProvider, err := local.NewProvider()

		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, localProvider)
	})

	t.Run("failed", func(t *testing.T) {
		t.Setenv("LOCAL_PATH", "")

		expectedMsg := "missing env LOCAL_PATH"

		defer func() {
			rec := recover()

			if rec == nil {
				t.Error("expected panic but got none")
			}

			if msg, ok := rec.(string); !ok || msg != expectedMsg {
				t.Errorf("expected panic message %q but got %v", expectedMsg, rec)
			}
		}()

		_, err := local.NewProvider()
		sdktesting.IsNotNull(t, err)
	})
}
