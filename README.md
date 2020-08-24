# magic
Magic Auth Admin API for Golang

## Usage

```go
package main

import (
	"fmt"

	"github.com/jacob-ebey/magic/admin"
)

var userProvidedDid = "someDidTokenFromAUser"

func main() {
	magic, err := admin.NewMagicAdmin("your_magic_secret")
	if err != nil {
		panic(err.Error())
	}

	user, err := magic.GetMetadataByToken(userProvidedDid)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(user.Email, user.Issuer, user.PublicAddress)

	if err := magic.LogoutByIssuer(user.Issuer); err != nil {
		panic(err.Error())
	}
}

```
