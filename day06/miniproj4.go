// url shortener
// building a backend application that accepts a long url via an API request, generates a short code, stores the mapping in memory and redirects
// we create the random code for the short url by randomly selecting characters from a slice of bytes
// obtain the url as a raw json data stream and verify it is not empty and generate the new shortcode
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

//in-memory database using map
var urlStore = make(map[string]string)

//struct to define the json input and output

//json input
type ShortenRequest struct {
	OriginalURL string `json:"originalURL"`
}

//json output
type ShortenResponse struct {
    ShortCode string `json:"shortCode"`
}

//function to create a random 6 character string
func generateCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	//creating an empty slize of 6 bytes that will generate and store the shortcode
	shortCode := make([]byte, 6)

	//loop 6 times to obtain a random byte from the charset and convert thefinal byte array to a printable string and return it
	for i := 0; i < 6; i++ {
		randomIndex := rand.Intn(len(charset))
		shortCode = append(shortCode, charset[randomIndex])
	}

	codeString := string(shortCode)

	return codeString
}

//function to take a long url and create a short code
func shortenHandler(w http.ResponseWriter, r *http.Request) {

	//creating an instance of thejson input struct
	var req ShortenRequest

	//a decoder that reads from the incoming request body using json package
	decoder := json.NewDecoder(r.Body)

	//decoder translates the json and injects it to the req struct
	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
    
	//validating the input (if the request payload is empty)
	if req.OriginalURL == "" {
		http.Error(w, "originalURL is required", http.StatusBadRequest)
		return
	}

	//now generating a random 6 character code
	code := generateCode()

	//saving the mapping between original url and the code in the map
	urlStore[code] = req.OriginalURL

	//prepping the response data as the output json struct
	res := ShortenResponse{
		ShortCode: code,
	}

	//set the headers and the status code as success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//will create an encoder to send the json back to the user, call .encode on the res struct
	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}

//function to catch the short link and send the browser to real website
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	//extract the code from the url reuest path 
	code := r.PathValue("code")

	//look up the code in the mapto obtain the long url
	longURL, exists := urlStore[code]
	if !exists {
		http.Error(w, "404 - Short URL not found", http.StatusNotFound)
		return
	}
    
	//redirect the user to the original website
	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	//register the shortenhandler to handle post reuests
	http.HandleFunc("POST /api/shorten", shortenHandler)

	//register the redirecthandler to handle get requests 
	http.HandleFunc("GET /{code}", redirectHandler)

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Server crashed: %v\n", err)
	}
}
