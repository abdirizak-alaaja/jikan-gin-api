package main

import (
	"fmt"

	internal "github.com/abdirizak-alaaja/jikan-gin-api/internal/app"
)

func main() {
	fmt.Println("#==================================#")
	fmt.Println("#===========Starting Server =======#")
	fmt.Println("#==================================#")

	if err := internal.StartApp(); err != nil {
		panic(err)
	}}
