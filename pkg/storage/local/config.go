package local

import "os"

type Config struct {
	LocalPath string
}

func NewConfig() *Config {
	localPath := os.Getenv("LOCAL_PATH")

	if localPath == "" {
		panic("missing env LOCAL_PATH")
	}

	return &Config{
		LocalPath: localPath,
	}
}
