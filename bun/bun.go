package bun

import (
	"encoding/json"
	"net/http"
)

type Release struct {
	TagName     string `json:"tag_name"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
}

var httpClient = &http.Client{}
var url = "https://api.github.com/repos/oven-sh/bun/releases"

func ListVersions() ([]Release, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set User-Agent header

	req.Header.Set("User-Agent", "BVM (https://github.com/gaurishhs/bvm)")

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Parse the response
	var releases []Release
	err = json.NewDecoder(response.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}
	return releases, nil
}
