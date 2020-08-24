package admin

import (
	"errors"
	"net/http"
)

var MagicApiKeyMissingError = errors.New("api key missing")

type MagicAdminOptions struct {
	Endpoint string
}

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type MagicAdmin struct {
	secretApiKey string
	Options      MagicAdminOptions
	Client       Client
}

func NewMagicAdmin(secretApiKey string) (*MagicAdmin, error) {
	if secretApiKey == "" {
		return nil, MagicApiKeyMissingError
	}

	return &MagicAdmin{
		secretApiKey,
		MagicAdminOptions{
			Endpoint: "https://api.magic.link",
		},
		&http.Client{},
	}, nil
}
