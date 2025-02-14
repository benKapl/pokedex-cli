package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name

	var data []byte

	// Check cache before making request
	if cacheData, exists := c.cache.Get(url); exists {
		data = cacheData
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error creating request : %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error making request : %w", err)
		}
		defer resp.Body.Close()

		apiData, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error reading response body : %w", err)
		}
		c.cache.Add(url, apiData)
		data = apiData
	}

	pokemon := Pokemon{}

	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error unmarshalling json data : %w", err)
	}

	return pokemon, nil
}
