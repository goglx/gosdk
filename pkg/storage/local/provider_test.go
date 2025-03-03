package local_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/pkg/storage/local"
)

var errMissingEnvironmentVariable = errors.New("missing environment variable")

func checkEnvVariables(provider string, variables []string) error {
	for _, v := range variables {
		if os.Getenv(v) == "" {
			return fmt.Errorf("%w: %s provider, variable %s", errMissingEnvironmentVariable, provider, v)
		}
	}

	return nil
}

func setupEnv() error {
	return checkEnvVariables("local", []string{"LOCAL_PATH"})
}

func TestNewProvider(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("LOCAL_PATH", "/tmp")
		sdktesting.IsNull(t, err)

		localProvider, err := local.NewProvider()
		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, localProvider)
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		err := os.Unsetenv("LOCAL_PATH")
		sdktesting.IsNull(t, err)

		_, err = local.NewProvider()
		sdktesting.IsNotNull(t, err)
		sdktesting.Equals(t, err.Error(), "missing env LOCAL_PATH")
	})
}
