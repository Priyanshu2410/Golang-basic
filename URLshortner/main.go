package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}

var urlDb = make(map[string]URL)

func generateShortURL(originalURL string) string {
	haser := md5.New()
	haser.Write([]byte(originalURL))
	data := haser.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	url := URL{
		ID:          shortURL,
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		CreatedAt:   time.Now(),
	}
	urlDb[shortURL] = url
	return shortURL
}

func getURL(shortURL string) (URL, error) {
	url, found := urlDb[shortURL]
	if !found {
		return URL{}, fmt.Errorf("URL not found")
	}
	return url, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		OriginalURL string `json:"original_url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortURL := createURL(data.OriginalURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	url, err := getURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusSeeOther)
}
func main() {
	fmt.Println("url_shortener")

	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect", redirectURLHandler)
	//start server
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
