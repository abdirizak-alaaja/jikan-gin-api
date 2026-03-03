A Go wrapper for the [Jikan v4 API](https://jikan.moe/).

> **Notice:** Jikan v4 introduces breaking changes. You **will** need to make
> changes to your code when upgrading to v1.2.0+.


### Installation

To install/update: `go install github.com/abdirizak-alaaja/jikan-gin-api@latest`

To import: `import "github.com/abdirizak-alaaja/jikan-gin-api"` and use as `jikan`


### Usage

##### Basic Example

```go
package main

import (
	"fmt"
	"github.com/abdirizak-alaaja/jikan-gin-api"
)

func main() {
	// Get anime
	anime, err := anime.GetAnimeById(52299)
	if err != nil {
		panic(err)
	}
	fmt.Println(anime.Data.Title)
}
```

```
Output:

 Ore dake Level Up na Ken
```

##### Search Query Example

Please refer to the
official [Jikan API documentation](https://docs.api.jikan.moe/) for
clarification on query parameters.

```go
package main

import (
	"fmt"
	"github.com/abdirizak-alaaja/jikan-gin-api"
	"net/url"
)

func main() {
	// Setup query
	query := url.Values{}
	query.Set("q", "solo leveling")
	query.Set("type", "tv")

	// Search anime
	search, err := anime.GetAnimeSearch(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(search.Data[0].Score)
}
```

```
Output:

8.18
```


##### Troubleshooting

If it is necessary to modify the http client (eg. modify timeout), you can
access the client via `jikan.Client`.

Example: setting the client timeout from 60 seconds (default) to 10 seconds.

```go
config.ClientTimeout = time.Second * 10
```