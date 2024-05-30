package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Hamziee/Hamzie-API/handlers/gifs"
	"github.com/Hamziee/Hamzie-API/handlers/images"
)

func main() {
	// Define default port
	//port := ":5554" // Dev port
	port := ":5555" // Default port for production

	// Check if any arguments are provided
	if len(os.Args) > 1 {
		// If argument provided, use it as port
		port = ":" + os.Args[1]
	}

	// Define routes
	// Define Images routes
	http.HandleFunc("/v1/images/cats", images.GetRandomCatImage)
	http.HandleFunc("/v1/images/xiaojie", images.GetRandomXiaojieImage)

	// Define Gifs routes
	http.HandleFunc("/v1/gifs/hug", gifs.GetRandomHugGif)
	http.HandleFunc("/v1/gifs/kiss", gifs.GetRandomKissGif)
	http.HandleFunc("/v1/gifs/headpats", gifs.GetRandomHeadpatsGif)
	http.HandleFunc("/v1/gifs/slap", gifs.GetRandomSlapGif)

	// Custom 404 page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `404 - Page Not Found. Make sure to read the docs: https://docs-api.hamzie.site/`)
	})

	// Start the server
	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
