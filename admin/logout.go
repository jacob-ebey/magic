package admin

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var MagicLogoutFailedError = errors.New("logout failed")

func (magic *MagicAdmin) LogoutByIssuer(issuer string) error {
	client := &http.Client{}

	body, err := json.Marshal(map[string]string{
		issuer: issuer,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", magic.Options.Endpoint+"/v2/admin/auth/user/logout", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("X-Magic-Secret-key", magic.secretApiKey)

	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		return MagicLogoutFailedError
	}
	defer res.Body.Close()

	return nil
}

func (magic *MagicAdmin) LogoutByToken(didToken string) error {
	token, err := ParseDidToken(didToken)
	if err != nil {
		return err
	}

	return magic.LogoutByIssuer(token.Parsed.ParsedClaim.Iss)
}
