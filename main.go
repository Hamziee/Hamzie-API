package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type CatLinks struct {
	Links []string `json:"links"`
}

func loadCatLinks(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var catLinks CatLinks
	err = json.Unmarshal(data, &catLinks)
	if err != nil {
		return nil, err
	}

	return catLinks.Links, nil
}

func handleRequests(catLinks []string) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cats", func(w http.ResponseWriter, r *http.Request) {
		randomCatLink(w, catLinks)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	catLinks, err := loadCatLinks("links/cats.json")
	if err != nil {
		log.Fatal(err)
	}
	handleRequests(catLinks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Found, please follow the docs to use the API correctly.")
}

func randomCatLink(w http.ResponseWriter, catLinks []string) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(catLinks))
	randomCatLink := catLinks[randomIndex]
	fmt.Fprintf(w, "Random Cat Link: %s", randomCatLink)
}
