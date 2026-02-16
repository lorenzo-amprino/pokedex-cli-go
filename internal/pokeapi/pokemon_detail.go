package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {

	url := BaseURL + "/pokemon/" + pokemonName
	res, err := http.Get(url)

	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("error: received status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	response := Pokemon{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return response, nil
}
