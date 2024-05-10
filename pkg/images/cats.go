package images

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func generateCatImages() []string {
	var images []string
	baseURL := "https://cdn.hamzie.site/Hamzie-API/images/cats/%d.jpg"

	for i := 1; i <= 135; i++ {
		imageURL := fmt.Sprintf(baseURL, i)
		images = append(images, imageURL)
	}

	return images
}

var CatImages = generateCatImages()

type ImageCatResponse struct {
	Link string `json:"link"`
}

func GetRandomCatImage(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(CatImages))
	randomCat := CatImages[randomIndex]

	response := ImageCatResponse{
		Link: randomCat,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
