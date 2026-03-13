# Jikan Go SDK

A lightweight Go SDK for interacting with the Jikan v4 API.

This library provides a simple and idiomatic Go interface for retrieving
anime data from MyAnimeList through Jikan.

It is designed for developers who want to integrate anime data into
their Go applications without dealing directly with HTTP requests.

> Important
>
> Jikan v4 introduced structural changes to the API. If you are
> upgrading from versions prior to v1.2.0, your code may require updates
> to match the new response models.

---

# Features

- Fetch anime details by ID
- Search anime using query parameters
- Retrieve anime characters, episodes, and media
- Simple and minimal dependencies
- Fully compatible with Go modules

---

# Installation

Install the module using `go get`:

```bash
go get github.com/abdirizak-alaaja/jikan-gin-api
```

Then import the package in your Go application:

```go
import "github.com/abdirizak-alaaja/jikan-gin-api/anime"
```

---

# Quick Start

## Get Anime by ID

```go
package main

import (
    "fmt"

    "github.com/abdirizak-alaaja/jikan-gin-api/anime"
)

func main() {

    animeData, err := anime.GetAnimeById(52299)
    if err != nil {
        panic(err)
    }

    fmt.Println(animeData.Data.Title)
}
```

Example output:

    Ore dake Level Up na Ken

---

# Searching Anime

Search queries can be built using `net/url` parameters.

```go
package main

import (
    "fmt"
    "net/url"

    "github.com/abdirizak-alaaja/jikan-gin-api/anime"
)

func main() {

    query := url.Values{}
    query.Set("q", "solo leveling")
    query.Set("type", "tv")

    result, err := anime.GetAnimeSearch(query)
    if err != nil {
        panic(err)
    }

    fmt.Println(result.Data[0].Title)
    fmt.Println(result.Data[0].Score)
}
```

---

# Response Structure

Responses follow the standard structure used by Jikan:

```json
{
  "data": {},
  "pagination": {}
}
```

The `data` field contains the requested resource, while `pagination` is
included for endpoints returning lists.

---

# Configuration

## HTTP Client Timeout

The default HTTP timeout is 60 seconds.

You can modify it globally:

```go
config.ClientTimeout = time.Second * 10
```

---

# Error Handling

All SDK functions return standard Go errors.

```go
animeData, err := anime.GetAnimeById(1)
if err != nil {
    log.Fatal(err)
}
```

Always check errors before accessing response data.

---

# Contributing

Contributions are welcome.

If you would like to improve the SDK:

1.  Fork the repository
2.  Create a feature branch
3.  Commit your changes
4.  Open a Pull Request

Please ensure code is formatted and documented before submitting.

---

# License

This project is released under the MIT License.
