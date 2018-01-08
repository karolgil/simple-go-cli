package gitHub

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/"
)

type User struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	URL               string      `json:"url"`
	Company           string      `json:"company"`
	Email             string      `json:"email"`
}

func GetUser(username string) User {
	resp, err := http.Get(apiURL + userEndpoint + username)

	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var user User
	json.Unmarshal(body, &user)
	return user
}