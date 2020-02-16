package octocat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Octocat type represents the Octocat user
type Octocat struct {
	Login     string
	ID        int
	URL       string `json:"html_url"`
	Name      string
	Followers int
	Following int
	Company   string
	Location  string
}

// Get retrieves the @octocat profile from GitHub.com
func Get() (*Octocat, error) {
	url := "https://api.github.com/users/octocat"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code, %d, when fetching %s", resp.StatusCode, url)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var octocat Octocat
	err = json.Unmarshal(body, &octocat)
	if err != nil {
		return nil, err
	}

	return &octocat, nil
}
