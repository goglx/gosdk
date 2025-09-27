package local_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage/local"
)

var errInvalidFileName = errors.New("invalid file name")

func TestProvider_Upload(t *testing.T) {
	t.Parallel()

	mockFS := newMockFileSystem()
	mockFS.mkdirAllFunc = func(path string, perm os.FileMode) error {
		if path != "." {
			t.Errorf("expected path %s, got %s", "expected/path", path)
		}

		return nil
	}
	mockFS.createFunc = func(name string) (*os.File, error) {
		if name != "./test/test-id" {
			return nil, fmt.Errorf("%w expected name got %s", errInvalidFileName, name)
		}

		return os.CreateTemp(t.TempDir(), "test-id")
	}

	provider := newMockProvider(newMockConfig(), mockFS)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		upload, err := provider.Upload(context.Background(), &types.File{
			ID:          "test-id",
			Data:        []byte("test-id"),
			ContentType: "text/plain",
		})

		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, upload)
		sdktesting.Equals(t, upload.ID, "test-id")
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		_, err := provider.Upload(context.Background(), &types.File{
			ID:          "wrong",
			Data:        []byte("wrong"),
			ContentType: "text/plain",
		})

		sdktesting.IsNotNull(t, err)
		sdktesting.Equals(t, err.Error(), "error creating file, invalid file name expected name got ./test/wrong")
	})
}

func TestNewProvider(t *testing.T) {
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
