# Json Client
A Json-input / JsonObject-output client.

# how to use

get

```
$ go get github.com/mokelab-go/jsonclient
```

write

```go
package main

import (
        "fmt"
        jc "github.com/mokelab-go/jsonclient/http"
)

func main() {
        client := jc.NewClient()
        client.SetURL("https://api-jp.kii.com/api/oauth2/token")
        client.SetMethod("POST")
        client.SetHeader("x-kii-appid", "aa")
        client.SetHeader("x-kii-appkey", "bb")
        resp, err := client.Send(map[string]interface{}{
                "username": "fkm",
                "password": "123456",
        })
        if err != nil {
                fmt.Printf("Error : %s\n", err)
                return
        }
        fmt.Printf("Status = %d\n", resp.Status)
        fmt.Printf("Response headers\n")
        for key, value := range resp.Headers {
                fmt.Printf("%s = %s\n", key, value)
        }
        fmt.Printf("Response body\n")
        // resp.Body is map[string]interface{}
        fmt.Printf("%s", resp.Body)
}
```
