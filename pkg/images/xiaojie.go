package images

import (
	"fmt"
	"math/rand"
	"net/http"
)

func generateXiaojieImages() []string {
	var images []string
	baseURL := "https://cdn.hamzie.us.to/Hamzie-API/images/xiaojie/%d.jpg"

	for i := 1; i <= 135; i++ {
		imageURL := fmt.Sprintf(baseURL, i)
		images = append(images, imageURL)
	}

	return images
}

var XiaojieImages = generateXiaojieImages()

func GetRandomXiaojieImage(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(XiaojieImages))
	randomCat := XiaojieImages[randomIndex]

	http.Redirect(w, r, randomCat, http.StatusFound)
}
