package admin

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLogoutCallsEndpoint(t *testing.T) {
	magic, err := NewMagicAdmin("hello")

	if err != nil {
		t.Fatal(err.Error())
	}

	client := &clientMock{}
	client.res = &http.Response{
		StatusCode: 200,
		Body:       &bodyMock{},
	}
	magic.Client = client

	if err := magic.LogoutByToken(validDid); err != nil {
		t.Fatal(err.Error())
	}

	if client.calls[0].URL.Scheme != "https" {
		t.Fatal("did not call the proper url scheme", client.calls[0].URL.Scheme)
	}
	if client.calls[0].URL.Host != "api.magic.link" {
		t.Fatal("did not call the proper url host", client.calls[0].URL.Host)
	}
	if client.calls[0].URL.Path != "/v2/admin/auth/user/logout" {
		t.Fatal("did not call the proper url path", client.calls[0].URL.Path)
	}

	if client.calls[0].Header.Get("X-Magic-Secret-key") != "hello" {
		t.Fatal("did not call with proper header")
	}

	bytes, err := ioutil.ReadAll(client.calls[0].Body)
	if err != nil {
		t.Fatal(err.Error())
	}

	if string(bytes) != "{\"issuer\":\""+parsedDidToken.ParsedClaim.Iss+"\"}" {
		t.Fatal("unexpected body", string(bytes))
	}
}
