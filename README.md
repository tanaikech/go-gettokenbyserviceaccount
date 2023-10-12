# go-gettokenbyserviceaccount

[MIT License](LICENCE)

<a name="top"></a>

# Overview

This is a Golang library to retrieve access token from [Service Account of Google](https://developers.google.com/identity/protocols/OAuth2ServiceAccount) without using [Google's OAuth2 package](https://github.com/golang/oauth2).

# Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/tanaikech/go-gettokenbyserviceaccount"
)

func main() {
	privateKey := "-----BEGIN PRIVATE KEY-----\n###-----END PRIVATE KEY-----\n"
	clientEmail := "###"
	scopes := "https://www.googleapis.com/auth/drive https://www.googleapis.com/auth/spreadsheets"
	impersonateEmail := ""
	res, err := gettokenbyserviceaccount.Do(privateKey, clientEmail, impersonateEmail, scopes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.AccessToken) // In this case, the access token is retrieved.
}
```

- `privateKey`, `clientEmail`, `impersonateEmail` and `scopes` are string values.

- When you want to use multiple scopes, please put them separated by a space like `"https://www.googleapis.com/auth/drive https://www.googleapis.com/auth/spreadsheets"`.

You can obtain the access token like below.

```
{
  "access_token": "#####",
  "expires_in": 3600,
  "token_type": "Bearer",
  "start_time": 1234567890,
  "end_time": 1234567890
}
```

[You can also retrieve this result using Google's OAuth2 package.](https://gist.github.com/tanaikech/4b4cb27ece27573b3f4df0e050b52330) I created this library to study the JWT process.

---

<a name="licence"></a>

# Licence

[MIT](LICENCE)

<a name="author"></a>

# Author

[Tanaike](https://tanaikech.github.io/about/)

If you have any questions and commissions for me, feel free to tell me.

<a name="Update_History"></a>

# Update History

- v1.0.0 (December 11, 2018)

  1. Initial release.

- v1.0.1 (October 12, 2023)

  1. Initial release.

[TOP](#top)
