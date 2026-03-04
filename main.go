package main

import (
	"fmt"

	"github.com/abdirizak-alaaja/jikan-gin-api/internal/app"
)

func main() {
	fmt.Println("#==================================#")
	fmt.Println("#===========Starting Server =======#")
	fmt.Println("#==================================#")

	if err := app.StartApp(); err != nil {
		panic(err)
	}
}
