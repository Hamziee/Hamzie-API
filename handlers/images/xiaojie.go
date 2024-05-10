package images

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func generateXiaojieImages() []string {
	var images []string
	baseURL := "https://cdn.hamzie.site/Hamzie-API/images/xiaojie/%d.jpg"

	for i := 1; i <= 135; i++ {
		imageURL := fmt.Sprintf(baseURL, i)
		images = append(images, imageURL)
	}

	return images
}

var XiaojieImages = generateXiaojieImages()

type ImageXiaojieResponse struct {
	Link string `json:"link"`
}

func GetRandomXiaojieImage(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(XiaojieImages))
	randomCat := XiaojieImages[randomIndex]

	response := ImageXiaojieResponse{
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
