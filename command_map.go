package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func mapCommand(cfg *commandsConfig) error {

	var url string
	if cfg.locationArea.next != "" {
		url = cfg.locationArea.next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	mapResponse := mapResponse{}
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		return err
	}

	for _, result := range mapResponse.Results {
		fmt.Println(result.Name)
	}

	cfg.locationArea.next = mapResponse.Next
	cfg.locationArea.previous = mapResponse.Previous

	return nil
}

type mapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
