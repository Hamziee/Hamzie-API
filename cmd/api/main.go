package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Hamziee/Hamzie-API/pkg/images"
)

func main() {
	// Define default port
	port := ":8080"

	// Check if any arguments are provided
	if len(os.Args) > 1 {
		// If argument provided, use it as port
		port = ":" + os.Args[1]
	}

	// Define routes
	http.HandleFunc("/v1/images/cats", images.GetRandomCatImage)
	http.HandleFunc("/v1/images/xiaojie", images.GetRandomXiaojieImage)

	// Custom 404 page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 - Page Not Found. Make sure to read the docs at: docs.hamzie.us.to\n")
	})

	// Start the server
	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
