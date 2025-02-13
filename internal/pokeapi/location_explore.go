package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ExploreLocation
func (c *Client) ExploreLocation(locationName string) (RespDetailedLocation, error) {
	url := baseUrl + "/location-area/" + locationName

	var data []byte

	// Check cache before making request
	if cacheData, exists := c.cache.Get(url); exists {
		data = cacheData
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespDetailedLocation{}, fmt.Errorf("Error creating request : %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespDetailedLocation{}, fmt.Errorf("Error making request : %w", err)
		}
		defer resp.Body.Close()

		apiData, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespDetailedLocation{}, fmt.Errorf("Error reading response body : %w", err)
		}
		c.cache.Add(url, apiData)
		data = apiData
	}

	locationResp := RespDetailedLocation{}

	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespDetailedLocation{}, fmt.Errorf("Error unmarshalling json data : %w", err)
	}

	return locationResp, nil
}
