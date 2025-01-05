package file

type Provider string

const (
	S3    Provider = "s3"
	R2    Provider = "r2"
	GCS   Provider = "gcs"
	Local Provider = "local"
)
