package main

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"net/url"
)

type PinboardClient struct {
	APIUrl   string
	APIToken string
}

func NewPinboardClient(apiUrl string, apiToken string) *PinboardClient {
	return &PinboardClient{
		APIToken: apiToken,
		APIUrl:   apiUrl,
	}
}

// doRequest sends an HTTP GET request to the Pinboard API.
func (c *PinboardClient) doRequest(endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := c.APIUrl
	params["auth_token"] = c.APIToken
	params["format"] = "json"

	query := url.Values{}
	for key, value := range params {
		query.Set(key, value)
	}

	reqURL := baseURL + endpoint + "?" + query.Encode()
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.New("API request failed with status " + resp.Status)
	}
	return resp, nil
}

func (p *PinboardClient) GetRecentBookmarks(tag string, count int) ([]Bookmark, error) {
	params := map[string]string{}

	if tag != "" {
		params["tag"] = tag
	}
	if count > 0 {
		params["count"] = fmt.Sprintf("%d", count)
	}

	endpoint := "posts/recent"
	req, err := p.doRequest(endpoint, params)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	type recentResponse struct {
		Posts []Bookmark `json:"posts"`
	}
	resp := &recentResponse{}
	err = json.NewDecoder(req.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp.Posts, nil
}
