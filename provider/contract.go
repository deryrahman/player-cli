package provider

import "errors"

var (
	ErrNotFound = errors.New("Couldn't fetch audio URL")
)

type AudioMeta struct {
	Title    string
	MimeType string
	URL      string
}
