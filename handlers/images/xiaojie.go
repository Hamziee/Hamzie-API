package images

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func generateXiaojieImages() ([]string, error) {
	var images []string
	baseURL := "https://cdn.hamzie.site/Hamzie-API/images/xiaojie/%d.jpg"

	resp, err := http.Get("https://raw.githubusercontent.com/Hamziee/Hamzie-API/main/handlers/images/xiaojie-index.md")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	index := strings.TrimSpace(string(body))
	num, err := strconv.Atoi(index)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= num; i++ {
		imageURL := fmt.Sprintf(baseURL, i)
		images = append(images, imageURL)
	}

	return images, nil
}

var XiaojieImages, _ = generateXiaojieImages()

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
