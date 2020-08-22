package admin

import (
	"encoding/base64"
	"encoding/json"
)

type DidClaim struct {
	Iat int    // Issued At Timestamp
	Ext int    // Expiration Timestamp
	Iss string // Issuer of DID Token
	Sub string // Subject
	Aud string // Audience
	Nbf int    // Not Before Timestamp
	Tid string // DID Token ID
	Add string // Encrypted signature of arbitrary data
}

type ParsedDidToken struct {
	Proof       string
	ParsedClaim DidClaim
}

type ParseDidTokenResult struct {
	Raw    []string
	Parsed ParsedDidToken
}

func ParseDidToken(didToken string) (ParseDidTokenResult, error) {
	result := ParseDidTokenResult{}

	decodedBase64, err := base64.StdEncoding.DecodeString(didToken)
	if err != nil {
		return result, MagicInvalidDidError
	}

	var decodedJson []string
	if err := json.Unmarshal(decodedBase64, &decodedJson); err != nil {
		return result, MagicInvalidDidError
	}

	result.Raw = decodedJson

	var decodedClaim DidClaim
	if err := json.Unmarshal([]byte(decodedJson[1]), &decodedClaim); err != nil {
		return result, err
	}

	result.Parsed.ParsedClaim = decodedClaim
	result.Parsed.Proof = decodedJson[0]

	return result, nil
}
