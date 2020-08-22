package admin

import "testing"

func TestCanCreateNewMagicAdmin(t *testing.T) {
	magic, err := NewMagicAdmin("hello")

	if err != nil {
		t.Fatal(err.Error())
	}

	if magic.secretApiKey != "hello" {
		t.Fatal("secretApiKey not set properly")
	}

	if magic.Options.Endpoint != "https://api.magic.link" {
		t.Fatal("expected endpoint to be https://api.magic.link")
	}
}

func TestErrorsWithoutSecretApiKey(t *testing.T) {
	magic, err := NewMagicAdmin("")

	if err != MagicApiKeyMissingError {
		t.Fatal("expected MagicApiKeyMissingError")
	}

	if magic != nil {
		t.Fatal("expected magic admin instance to be nil")
	}
}
