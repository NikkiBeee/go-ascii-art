package main

import (
	"net/http"
	"./routes"
)

func main() {
	routes.SetupRoutes()
	http.ListenAndServe(":3000", nil)
}
