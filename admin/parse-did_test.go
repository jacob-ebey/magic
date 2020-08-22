package admin

import (
	"testing"

	"github.com/go-test/deep"
)

var validDid = "WyIweGUwMjQzNTVlNDI5ZGNhZDM1MTdhZDk5ZWEzNDEwYWJmZDQ1YjBiNjM5OGIwNjY1NGRiYTQxNzljODdlMTYyNzgxNTc1YjA5ODFjNjU4ZjcwMjYwZTQ5MjMwZGE5NDg4YTA0ZDk5NzBlYjM4ZTZmZGRlY2Q2NTA5YTAyN2IwOGI5MWIiLCJ7XCJpYXRcIjoxNTg1MDExMjA0LFwiZXh0XCI6MTkwMDQxMTIwNCxcImlzc1wiOlwiZGlkOmV0aHI6MHhCMmVjOWI2MTY5OTc2MjQ5MWI2NTQyMjc4RTlkRkVDOTA1MGY4MDg5XCIsXCJzdWJcIjpcIjZ0RlhUZlJ4eWt3TUtPT2pTTWJkUHJFTXJwVWwzbTNqOERReWNGcU8ydHc9XCIsXCJhdWRcIjpcImRpZDptYWdpYzpmNTQxNjhlOS05Y2U5LTQ3ZjItODFjOC03Y2IyYTk2YjI2YmFcIixcIm5iZlwiOjE1ODUwMTEyMDQsXCJ0aWRcIjpcIjJkZGY1OTgzLTk4M2ItNDg3ZC1iNDY0LWJjNWUyODNhMDNjNVwiLFwiYWRkXCI6XCIweDkxZmJlNzRiZTZjNmJmZDhkZGRkZDkzMDExYjA1OWI5MjUzZjEwNzg1NjQ5NzM4YmEyMTdlNTFlMGUzZGYxMzgxZDIwZjUyMWEzNjQxZjIzZWI5OWNjYjM0ZTNiYzVkOTYzMzJmZGViYzhlZmE1MGNkYjQxNWU0NTUwMDk1MmNkMWNcIn0iXQ=="

var parsedDidToken = ParsedDidToken{
	Proof: "0xe024355e429dcad3517ad99ea3410abfd45b0b6398b06654dba4179c87e162781575b0981c658f70260e49230da9488a04d9970eb38e6fddecd6509a027b08b91b",
	ParsedClaim: DidClaim{
		Iat: 1585011204,
		Ext: 1900411204,
		Iss: "did:ethr:0xB2ec9b61699762491b6542278E9dFEC9050f8089",
		Sub: "6tFXTfRxykwMKOOjSMbdPrEMrpUl3m3j8DQycFqO2tw=",
		Aud: "did:magic:f54168e9-9ce9-47f2-81c8-7cb2a96b26ba",
		Nbf: 1585011204,
		Tid: "2ddf5983-983b-487d-b464-bc5e283a03c5",
		Add: "0x91fbe74be6c6bfd8ddddd93011b059b9253f10785649738ba217e51e0e3df1381d20f521a3641f23eb99ccb34e3bc5d96332fdebc8efa50cdb415e45500952cd1c",
	},
}

func TestDecodesDidToken(t *testing.T) {
	decoded, err := ParseDidToken(validDid)

	if err != nil {
		t.Fatal("did not expect an error")
	}

	if decoded.Raw[0] != decoded.Parsed.Proof {
		t.Fatal("raw[0] should equal the parsed proof")
	}

	if diff := deep.Equal(decoded.Parsed, parsedDidToken); diff != nil {
		t.Fatal(diff)
	}
}
