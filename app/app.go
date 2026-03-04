package app

import (
	"fmt"

	"github.com/abdirizak-alaaja/jikan-gin-api/internal/router"
)

func StartApp() {

	r := router.SetupRoutes()

	addr := ":8080"
	fmt.Printf("Running app on : %s\n", addr)
	r.Run(addr)

}
