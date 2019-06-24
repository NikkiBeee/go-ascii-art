package main

import (
	"fmt"
	"net/http"

	"./routes"
)

func main() {
	// fmt.Println("main is running")
	routes.SetupRoutes()
	http.ListenAndServe(":8000", nil)
}
