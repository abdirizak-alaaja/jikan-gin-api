# Jikan Gin API

A Go wrapper and REST API server for the [Jikan v4 API](https://jikan.moe/).

This project provides a dual-purpose toolset:

- A **Gin-based REST API server** that mirrors Jikan v4 endpoints
- A **Go module** that allows developers to interact with Jikan directly in their own Go applications

> ⚠️ **Important**
>
> Jikan v4 introduces breaking changes.  
> If you are upgrading to **v1.2.0+**, you must update your code to match the new structure.

---

# Table of Contents

- [Installation](#installation)
- [Running the API Server](#running-the-api-server)
- [API Endpoints](#api-endpoints)
- [Using the Go Wrapper](#using-the-go-wrapper)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

---

# Installation

Install or update the module using `go install`:

```bash
go install github.com/abdirizak-alaaja/jikan-gin-api@latest
```

To use the wrapper in your Go code, import the package:

```go
import "github.com/abdirizak-alaaja/jikan-gin-api/anime"
```

---

# Running the API Server

If you want to host your own **local mirror or proxy of the Jikan API**:

Start the server:

```bash
go run main.go
```

Default Address:

```
http://localhost:8080
```

Base API Path:

```
/api/v4
```

---

# API Endpoints

| Method | Endpoint | Description |
|------|------|------|
| GET | `/api/v4/anime/:id` | Get specific anime details by ID |
| GET | `/api/v4/anime/:id/characters` | Get characters for an anime |
| GET | `/api/v4/anime/:id/episodes` | Get episode list |
| GET | `/api/v4/anime/:id/pictures` | Get image gallery |
| GET | `/api/v4/anime?q=<query>` | Search anime |

Example Request:

```
GET /api/v4/anime/52299
```

Example Response:

```json
{
  "data": {
    "mal_id": 52299,
    "title": "Ore dake Level Up na Ken",
    "score": 8.4
  }
}
```

---

#  Using the  Wrapper

You can use this project as a **Go library** inside your own applications.

## Basic Example

```go
package main

import (
	"fmt"
	"github.com/abdirizak-alaaja/jikan-gin-api/anime"
)

func main() {

	// Get anime by ID
	animeData, err := anime.GetAnimeById(52299)
	if err != nil {
		panic(err)
	}

	fmt.Println(animeData.Data.Title)
	// Output: Ore dake Level Up na Ken
}
```

---

## Search Example

For more complex queries, use `net/url` parameters.

Refer to the official Jikan docs for supported filters.

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

	fmt.Println(result.Data[0].Score)
	// Output: 8.18
}
```

---

# 🛠 Configuration

## Response Format

Most endpoints return the standard **Jikan structure**:

```json
{
  "data": {},
  "pagination": {}
}
```

---

## Modifying the HTTP Client

The default timeout is **60 seconds**.  
You can change it globally:

```go
	// Change timeout to 10 seconds
	config.ClientTimeout = time.Second * 10
```

---

# Contributing

Pull requests are welcome.

Steps:

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Submit a Pull Request

---

# License

MIT License