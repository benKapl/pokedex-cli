package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations
func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	var data []byte

	// Check cache before making request
	if cacheData, exists := c.cache.Get(url); exists {
		data = cacheData
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()

		apiData, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		c.cache.Add(url, apiData) // Cache response
		data = apiData
	}

	locationsResp := RespShallowLocations{}

	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
