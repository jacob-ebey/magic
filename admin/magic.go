package admin

import "errors"

var MagicApiKeyMissingError = errors.New("api key missing")

type MagicAdminOptions struct {
	Endpoint string
}

type MagicAdmin struct {
	secretApiKey string
	Options      MagicAdminOptions
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
	}, nil
}
