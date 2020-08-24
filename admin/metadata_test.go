package admin

import (
	"net/http"
	"strings"
	"testing"
)

func TestMetadataCallsEndpoint(t *testing.T) {
	magic, err := NewMagicAdmin("hello")

	if err != nil {
		t.Fatal(err.Error())
	}

	client := &clientMock{}
	client.res = &http.Response{
		StatusCode: 200,
		Body: &bodyMock{
			reader: strings.NewReader("{\"data\": {\"email\": \"test@test.com\",\"issuer\": \"issuervalue\", \"public_address\": \"public_addressvalue\"}, \"status\": \"ok\"}"),
		},
	}
	magic.Client = client

	user, err := magic.GetMetadataByToken(validDid)
	if err != nil {
		t.Fatal(err.Error())
	}

	if client.calls[0].URL.Scheme != "https" {
		t.Fatal("did not call the proper url scheme", client.calls[0].URL.Scheme)
	}
	if client.calls[0].URL.Host != "api.magic.link" {
		t.Fatal("did not call the proper url host", client.calls[0].URL.Host)
	}
	if client.calls[0].URL.Path != "/v1/admin/auth/user/get" {
		t.Fatal("did not call the proper url path", client.calls[0].URL.Path)
	}

	if client.calls[0].Header.Get("X-Magic-Secret-key") != "hello" {
		t.Fatal("did not call with proper header")
	}

	if user.Email != "test@test.com" {
		t.Fatal("email not expected value")
	}
	if user.Issuer != "issuervalue" {
		t.Fatal("issuer not expected value")
	}
	if user.PublicAddress != "public_addressvalue" {
		t.Fatal("public_address not expected value")
	}
}
