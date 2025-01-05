package types

type File struct {
	Id          string `json:"id"`
	ContentType string `json:"contentType"`
	Data        []byte `json:"data"`
}

type ProviderType string

const (
	S3    ProviderType = "s3"
	R2    ProviderType = "r2"
	GCS   ProviderType = "gcs"
	Local ProviderType = "local"
)
