package storage_test

import (
	"context"
	"errors"
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/internal/types"
	"gosdk/pkg/storage"
)

var errFailedToUpload = errors.New("failed to upload")
var errFailedToDelete = errors.New("failed to delete")
var errFailedToDownload = errors.New("failed to download")

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   types.ProviderType
		want   bool
		errMsg string
	}{
		{
			name:   types.S3,
			want:   true,
			errMsg: "",
		},
		{
			name:   types.R2,
			want:   true,
			errMsg: "",
		},
		{
			name:   types.GCS,
			want:   true,
			errMsg: "",
		},
		// {
		//	name:   types.Local,
		//	want:   false,
		//	errMsg: "failed to create local Provider: missing env LOCAL_PATH",
		// },
		{
			name:   "wrong",
			want:   false,
			errMsg: "unsupported Provider type",
		},
	}

	for _, testCase := range tests {
		t.Run(string(testCase.name), func(t *testing.T) {
			t.Parallel()

			provider, err := storage.New(testCase.name)

			if testCase.want {
				sdktesting.IsNull(t, err)
				sdktesting.IsNotNull(t, provider)
			}

			if !testCase.want {
				sdktesting.IsNotNull(t, err)
				sdktesting.Ok(t, err.Error(), testCase.errMsg)
			}
		})
	}
}

func TestNewLocalProvider(t *testing.T) {
	t.Setenv("LOCAL_PATH", "")

	tests := []struct {
		name   types.ProviderType
		want   bool
		errMsg string
	}{
		{
			name:   types.Local,
			want:   true,
			errMsg: "",
		},
	}

	for _, testCase := range tests {
		t.Run(string(testCase.name), func(t *testing.T) {
			t.Setenv("LOCAL_PATH", "/tmp")

			provider, err := storage.New(testCase.name)

			if testCase.want {
				sdktesting.IsNull(t, err)
				sdktesting.IsNotNull(t, provider)
			}

			if !testCase.want {
				sdktesting.IsNotNull(t, err)
				sdktesting.Ok(t, err.Error(), testCase.errMsg)
			}
		})
	}
}

func TestUpload(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		provider := &mockProvider{
			mockUpload: func(ctx context.Context, file *types.File) (*types.File, error) {
				return file, nil
			},
			mockDelete:   nil,
			mockDownload: nil,
		}

		manager := newMock(provider)
		sdktesting.IsNotNull(t, manager)
		sdktesting.IsNotNull(t, manager.Provider)

		upload, err := manager.Upload(context.TODO(), &types.File{
			ID:          "upload-id",
			ContentType: "text/plain",
			Data:        []byte("test-upload"),
		})

		sdktesting.IsNull(t, err)
		sdktesting.IsNotNull(t, upload)
		sdktesting.Equals(t, upload.ID, "upload-id")
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		provider := &mockProvider{
			mockUpload: func(ctx context.Context, file *types.File) (*types.File, error) {
				return nil, errFailedToUpload
			},
			mockDelete:   nil,
			mockDownload: nil,
		}

		manager := newMock(provider)
		sdktesting.IsNotNull(t, manager)

		_, err := manager.Upload(context.TODO(), &types.File{
			ID:          "upload-id",
			ContentType: "text/plain",
			Data:        []byte("test-upload"),
		})

		sdktesting.IsNotNull(t, err)
		sdktesting.Equals(t, err.Error(), "failed to upload file upload-id: failed to upload")
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		provider := &mockProvider{
			mockUpload:   nil,
			mockDelete:   func(ctx context.Context, key string) error { return nil },
			mockDownload: nil,
		}
		manager := newMock(provider)
		sdktesting.IsNotNull(t, manager)
		sdktesting.IsNotNull(t, manager.Provider)

		err := manager.Delete(context.TODO(), "delete-file")

		sdktesting.IsNull(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		provider := &mockProvider{
			mockUpload:   nil,
			mockDelete:   func(ctx context.Context, key string) error { return errFailedToDelete },
			mockDownload: nil,
		}
		manager := newMock(provider)
		sdktesting.IsNotNull(t, manager)

		err := manager.Delete(context.TODO(), "delete-file-id")
		sdktesting.IsNotNull(t, err)
		sdktesting.Equals(t, err.Error(), "failed to delete file delete-file-id: failed to delete")
	})
}

func TestDownload(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		provider := &mockProvider{
			mockUpload: nil,
			mockDelete: nil,
			mockDownload: func(ctx context.Context, key string) ([]byte, error) {
				return []byte("test-download"), nil
			},
		}
		manager := newMock(provider)
		sdktesting.IsNotNull(t, manager)
		sdktesting.IsNotNull(t, manager.Provider)

		download, err := manager.Download(context.TODO(), "test-download")
		sdktesting.IsNull(t, err)
		sdktesting.Equals(t, string(download), "test-download")
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		provider := &mockProvider{
			mockUpload: nil,
			mockDelete: nil,
			mockDownload: func(ctx context.Context, key string) ([]byte, error) {
				return nil, errFailedToDownload
			},
		}
		manager := newMock(provider)
		sdktesting.IsNotNull(t, manager)

		_, err := manager.Download(context.TODO(), "download-id")
		sdktesting.IsNotNull(t, err)
		sdktesting.Equals(t, err.Error(), "failed to download file download-id: failed to download")
	})
}
