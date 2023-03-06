# gOvkApi
API library for openvk.su

```bash
# go version 1.19

go get -u github.com/a1excoder/gOvkApi
# and use
```

### example
```go
package main

import (
	"fmt"
	"github.com/a1excoder/gOvkApi"
	"log"
)

func main() {
	auth := gOvkApi.GetAuthData("http://openvk.co")

	errOvk, err := auth.GetApiToken("a1excoder@openvk.uk", "12345678", "", false)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if errOvk != nil {
		log.Fatalln(errOvk.ErrorMsg)
	}

	profile, errOvk, err := auth.GetProfileInfo()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if errOvk != nil {
		log.Fatalln(errOvk.ErrorMsg)
	}
	
	fmt.Println(profile.FirstName, profile.LastName, profile.BirthdayDate)
}

// output
// Oleksandr Rudenko 11.11.2004
```
