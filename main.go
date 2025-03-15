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
		return
	}

	provider, err := local.NewProvider()
	if err != nil {
		return
	}

	file, err := provider.Upload(context.TODO(), &types.File{
		ID:          "test-id",
		Data:        []byte("test-id"),
		ContentType: "text/plain",
	})
	if err != nil {
		return
	}

	slog.Info(file.ID)

	manager, err := storage.New(types.Local)
	if err != nil {
		return
	}

	upload, err := manager.Upload(context.TODO(), &types.File{
		ID:          "test-id-2",
		Data:        []byte("test-id"),
		ContentType: "text/plain",
	})
	if err != nil {
		return
	}

	slog.Info(upload.ID)
}
