package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"gosdk/internal/types"
	"gosdk/pkg/storage"
	"gosdk/pkg/storage/local"
	"gosdk/pkg/storage/s3"
)

func main() {
	fileProvider := flag.String("file-provider", "", "storage provider to use (local or s3)")
	s3FileProvider := flag.Bool("s3-file-provider", false, "run s3 file provider")
	localFileProvider := flag.Bool("local-file-provider", false, "run local file provider")
	flag.Parse()

	switch *fileProvider {
	case "local":
		localProvider()
	case "s3":
		s3Provider()
	default:
		if *s3FileProvider {
			s3Provider()
		}

		if *localFileProvider {
			localProvider()
		}

		if !*s3FileProvider && !*localFileProvider && *fileProvider == "" {
			slog.Error("Unknown file provider")
			os.Exit(1)
		}
	}
}

func s3Provider() {
	slog.Info("s3 provider")

	slog.Info(os.Getenv("BUCKET_NAME"))

	provider, err := s3.NewProvider()
	if err != nil {
		slog.Error("failed to create s3 provider", "error", err)
		os.Exit(1)
	}

	file, err := provider.Upload(context.TODO(), &types.File{
		ID:          "test-s3-id",
		Data:        []byte("test-s3-id"),
		ContentType: "text/plain",
	})
	if err != nil {
		slog.Error("failed to upload file with s3 provider", "error", err)
		os.Exit(1)
	}

	slog.Info(file.ID)
}

func localProvider() {
	slog.Info("local provider")

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
