package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (*MapResponse, error) {

	url := BaseURL + "/location-area/"
	if pageURL != nil && *pageURL != "" {
		url = *pageURL
	}

	if c.cache != nil {
		if cachedResponse, found := c.cache.Get(url); found {
			mapResponse := MapResponse{}
			err := json.Unmarshal(cachedResponse, &mapResponse)
			if err != nil {
				return nil, err
			}
			return &mapResponse, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)

	mapResponse := MapResponse{}
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		return nil, err
	}

	return &mapResponse, nil

}
