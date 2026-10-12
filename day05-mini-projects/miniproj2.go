//mini project 2
//building a simple CLI tool to fetch user data from github API, parse JSON to struct and display a clean terminal dashboard
//use struct to parse the json fields
//construct the base github api url with the captured username
//memeory management to ensure that the network connection closes securely to preserve system resources
//data parsing: extracting and prepping the obtained data from the api call and translating it

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//struct to define the data we are extracting from the user profile (using struct tag to attach metadata for json translation)
//since encoding/json uses exact names from struct field as json key, we need to translate is into the appropriate json key
//the struct fields must begin with captial for it to be public
type GitHubProfile struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Company string `json:"company"`
	PublicRepos int `json:"public_repos"`
	Followers int `json:"followers"`
	Following int `json:"following"`
}

func main() {
	//capturing terminal input (github username)
	if len(os.Args) < 2 {
		fmt.Println("Provide valid github username")
		fmt.Println("usage: go run <filename> <username>")
		os.Exit(1)  //exit and forcefully term
	}

	username := os.Args[1]

	//constructing the api url for the github api call 
	githubAPI := "https://api.github.com/users/" + username

	//send an HTTP GET request to the githubapi url
	//define a client with a timeout (avoid using the default client in go which has no timeout)
	client := &http.Client{
		Timeout: 10*time.Second,
	}
    
	//creating a new request context
	req, err := http.NewRequest(http.MethodGet, githubAPI, nil) 
	if err != nil {
		log.Fatalf("Error creating the request: %v", err)
	}

	//sending the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	//closing the body to prevent resource leaks  using defer statement (so the network connection is guaranteed to close)
	defer resp.Body.Close()

	//handling 404 errors if the user does not exist
	if resp.StatusCode == 404 {
		fmt.Printf("Error: Github user '%s' does not exist.\n", username)
		os.Exit(1)
	}

	var profile GitHubProfile

	//reading all the data from the http response into a byte slice
	bodyBytes, err := io.ReadAll((resp.Body))
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	//unmarshal the raw bytes into the go struct
	err = json.Unmarshal(bodyBytes, &profile)
	if err != nil {
		fmt.Printf("Error parsing JSON data: %v\n", err)
		os.Exit(1)
	}

	//handling empty fields so the dasboard looks clean
	if profile.Name == "" {
		profile.Name = "N/A"
	}
	if profile.Bio == "" {
		profile.Bio = "N/A"
	}
	if profile.Company == "" {
		profile.Company = "N/A"
	}

	fmt.Println("\n========================================")
    fmt.Printf("       GITHUB PROFILE ANALYZER        \n")
    fmt.Println("========================================")
    fmt.Printf("Username      : %s\n", profile.Login)
    fmt.Printf("Name          : %s\n", profile.Name)
    fmt.Printf("Company       : %s\n", profile.Company)
    fmt.Printf("Bio           : %s\n", profile.Bio)
    fmt.Println("----------------------------------------")
    fmt.Printf("Public Repos  : %d\n", profile.PublicRepos)
    fmt.Printf("Followers     : %d\n", profile.Followers)
    fmt.Printf("Following     : %d\n", profile.Following)
    fmt.Println("========================================")

}

