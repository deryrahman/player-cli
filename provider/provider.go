package provider

type Provider interface {
	GetAudio(query string) (*AudioMeta, error)
}
