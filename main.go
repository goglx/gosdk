package main

import (
	"context"
	"log/slog"
	"os"

	"gosdk/internal/types"
	"gosdk/pkg/storage"
	"gosdk/pkg/storage/local"
)

func main() {
	err := os.Setenv("LOCAL_PATH", "./")
	if err != nil {
		slog.Error("failed to set LOCAL_PATH environment variable", "error", err)
		os.Exit(1)
	}

	provider, err := local.NewProvider()
	if err != nil {
		slog.Error("failed to create local provider", "error", err)
		os.Exit(1)
	}

	file, err := provider.Upload(context.TODO(), &types.File{
		ID:          "test-id",
		Data:        []byte("test-id"),
		ContentType: "text/plain",
	})
	if err != nil {
		slog.Error("failed to upload file with local provider", "error", err)
		os.Exit(1)
	}

	slog.Info(file.ID)

	manager, err := storage.New(types.Local)
	if err != nil {
		slog.Error("failed to create storage manager", "error", err)
		os.Exit(1)
	}

	upload, err := manager.Upload(context.TODO(), &types.File{
		ID:          "test-id-2",
		Data:        []byte("test-id"),
		ContentType: "text/plain",
	})
	if err != nil {
		slog.Error("upload failed", "error", err)
		os.Exit(1)
	}

	slog.Info(upload.ID)

	upload, err = manager.Upload(context.TODO(), &types.File{
		ID:          "./upload-dir/test-id-3",
		Data:        []byte("test-id"),
		ContentType: "text/plain",
	})
	if err != nil {
		slog.Error("upload failed", "error", err)
		os.Exit(1)
	}

	slog.Info(upload.ID)
}
