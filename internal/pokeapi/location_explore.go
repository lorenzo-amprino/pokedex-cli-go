package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (ExploreResponse, error) {

	url := BaseURL + "/location-area/" + locationName

	if c.cache != nil {
		if cachedResponse, found := c.cache.Get(url); found {
			exploreResponse := ExploreResponse{}
			err := json.Unmarshal(cachedResponse, &exploreResponse)
			if err != nil {
				return ExploreResponse{}, err
			}
			return exploreResponse, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return ExploreResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ExploreResponse{}, fmt.Errorf("error: received status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return ExploreResponse{}, err
	}

	c.cache.Add(url, body)

	exploreResponse := ExploreResponse{}
	err = json.Unmarshal(body, &exploreResponse)
	if err != nil {
		return ExploreResponse{}, err
	}

	return exploreResponse, nil
}
