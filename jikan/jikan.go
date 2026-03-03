// Package jikan provides Jikan API bindings for Go.
// Please note that all functions/structs are named in accordance to Jikan's endpoints/responses
package jikan

import (
	"net/http"
	"time"

	"github.com/abdirizak-alaaja/jikan-gin-api/config"
)

// Client is a *http.Client
var Client *http.Client

func init() {
	Client = &http.Client{Timeout: config.ClientTimeout * time.Second}
}
