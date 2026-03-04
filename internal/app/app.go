package internal 

import (
	"fmt"

	"github.com/abdirizak-alaaja/jikan-gin-api/internal/router"
)

func StartApp()error {

	r := internal.SetupRoutes()

	addr := ":8080"
	fmt.Printf("Running app on : %s\n", addr)
	r.Run(addr)
	return nil
}
