# altcraft-go

## Introduction
Golang client for [Altcraft API 1.1.](https://docs.altcraft.com/pages/viewpage.action?pageId=2195459)

## Install
Install with `go get`:

```bash
$ go get github.com/edenisn/altcraft-go
```

## Usage
```go
package main

import (
	"fmt"
	"os"

	altcraft "github.com/edenisn/altcraft-go"
)

const (
	token = "YOUR_ALTCRAFT_API_TOKEN"
	limit = 3
)

func main() {
	client := altcraft.New(token)
	// Fetch databases list
	list, err := client.GetDatabasesList(limit)
	if err != nil {
		fmt.Println("Failed to get databases list")
		os.Exit(1)
	}
	fmt.Println(*list)
}
```

### License
[The MIT License (MIT)](LICENSE)
