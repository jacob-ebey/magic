package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var MagicInvalidDidError = errors.New("invalid did token")
var MagicMetadataFailedError = errors.New("retrieving metadata failed")
var MagicMetadataDecodeFailedError = errors.New("decoding metadata failed")

type UserMetadata struct {
	Email         string
	Issuer        string
	PublicAddress string
}

func (magic *MagicAdmin) GetMetadataByIssuer(issuer string) (UserMetadata, error) {
	fmt.Println(issuer)
	result := UserMetadata{}

	client := &http.Client{}

	req, err := http.NewRequest("GET", magic.Options.Endpoint+"/v1/admin/auth/user/get", nil)
	if err != nil {
		return result, err
	}
	req.Header.Add("X-Magic-Secret-key", magic.secretApiKey)
	query := req.URL.Query()
	query.Add("issuer", issuer)
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		return result, MagicMetadataFailedError
	}
	defer res.Body.Close()

	var decodedResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&decodedResult); err != nil {
		return result, MagicMetadataDecodeFailedError
	}

	if decodedResult["status"].(string) == "ok" {
		data := decodedResult["data"].(map[string]interface{})
		result.Email = data["email"].(string)
		result.Issuer = data["issuer"].(string)
		result.PublicAddress = data["public_address"].(string)
	} else {
		return result, MagicMetadataDecodeFailedError
	}

	return result, nil
}

func (magic *MagicAdmin) GetMetadataByToken(didToken string) (UserMetadata, error) {
	token, err := ParseDidToken(didToken)
	if err != nil {
		return UserMetadata{}, err
	}

	return magic.GetMetadataByIssuer(token.Parsed.ParsedClaim.Iss)
}
