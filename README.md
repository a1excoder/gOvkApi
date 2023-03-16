# gOvkApi
API library for openvk.su

```bash
# go version 1.19

go get -u github.com/a1excoder/gOvkApi@v1.2.0
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

	errOvk := auth.GetApiToken("a1excoder", "12345678", "")
	if _, err := errOvk.GetError(); err != nil {
		log.Fatalln(err.Error())
	}

	profile, errOvk := auth.GetProfileInfo()
	if _, err := errOvk.GetError(); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(profile.FirstName, profile.LastName, profile.BirthdayDate)

	groups, errOvk := auth.GetGroups(profile.Id, 100, 0, "has_photo,photo_400_orig")
	if _, err := errOvk.GetError(); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("groups count: %d\n", groups.Count)
 
	for _, group := range groups.Items {
		fmt.Printf("id: %d | name: %s | is closed: ", group.Id, group.Name)
		if group.IsClosed {
			fmt.Printf("yes")
		} else {
			fmt.Printf("no")
		}

		if group.HasPhoto != nil && *group.HasPhoto != 0 && group.Photo400Orig != nil {
			fmt.Printf(" | %s", *group.Photo400Orig)
		}

		fmt.Println()
	}

}


// output
// Oleksandr Rudenko 11.11.2004
// groups count: 16
// id: 352 | name: 8-bit highlights | is closed: no | http://ru-spb.openvk.uk/hentai/b8/b8b853f2ceddfda34bb0864ee130f70c6ea8b7e0de409f0fba69c38fa9f9426aad63daf5719c7ed29b50a24e75fac47386c2a0d081a6379
// 6d650cbba419f8829_cropped/normal.jpeg
// id: 1668 | name: Pingvinus OPENVK | is closed: no | http://ru-spb.openvk.uk/hentai/8d/8df1be34b840a90935a9c2f3cd9939f9323b664280dfbb4ca7186fdd23b22dc7a5d89994aefafa4c4929fa581abc95d1715678bf4c8dbb
// 50f28a02d8f971d53e_cropped/normal.jpeg
// id: 1916 | name: Twitch | is closed: no | http://ru-spb.openvk.uk/hentai/dc/dc97c207b66d402992b4e1def81904b2b205a49309051e37798fc9a60f924d8ba2d5427c8d2b8d1207d224cccff58b30acab60fffc1bd799ba97ab9d
// d5ff3e27_cropped/normal.jpeg
// ...
// ...
```
