package images

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

var catImages = []string{
	"https://cdn.hamzie.us.to/Hamzie-API/images/cats/1.webp",
	"https://cdn.hamzie.us.to/Hamzie-API/images/cats/2.webp",
	"https://cdn.hamzie.us.to/Hamzie-API/images/cats/3.webp",
}

type CatImageResponse struct {
	Link string `json:"link"`
}

func GetRandomCatImage(w http.ResponseWriter, r *http.Request) {
	randomCat := catImages[rand.Intn(len(catImages))]

	response := CatImageResponse{
		Link: randomCat,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
