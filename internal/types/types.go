package types

type File struct {
	Id          string `json:"id"`
	ContentType string `json:"content_type"`
	Data        []byte `json:"data"`
}
