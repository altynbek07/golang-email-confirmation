package main

import (
	"fmt"
	"net/http"
)

func App() *http.ServeMux {
	router := http.NewServeMux()

	return router
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
