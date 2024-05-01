package images

import (
	"math/rand"
	"net/http"
)

var catImages = []string{
	"https://cdn.hamzie.us.to/Hamzie-API/images/cats/1.webp",
	"https://cdn.hamzie.us.to/Hamzie-API/images/cats/2.webp",
	"https://cdn.hamzie.us.to/Hamzie-API/images/cats/3.webp",
}

func GetRandomCatImage(w http.ResponseWriter, r *http.Request) {
	randomCat := catImages[rand.Intn(len(catImages))]

	http.Redirect(w, r, randomCat, http.StatusFound)
}
