# gosu-api-client

This projects aims to easily provide a client for the [osu!api v2](https://osu.ppy.sh/docs/index.html).

## Installation
    
    ```bash
    go get -u github.com/maxkruse/gosu-api-client
    ```

## Usage

```go
package main

import (
    "fmt"
    "github.com/maxkruse/gosu-api-client"
)

func main() {
    client := gosuapi.NewClient()
    user, err := client.GetUser("1199528")
    if err != nil {
        panic(err)
    }
    fmt.Println("%#+v", user)
}
```