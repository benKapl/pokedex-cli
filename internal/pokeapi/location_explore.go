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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedLocation{}, fmt.Errorf("Error creating request : %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocation{}, fmt.Errorf("Error making request : %w", err)
	}

	apiData, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocation{}, fmt.Errorf("Error reading response body : %w", err)
	}
	// add cache on this line
	data = apiData

	locationResp := RespDetailedLocation{}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespDetailedLocation{}, fmt.Errorf("Error unmarchalling json data : %w", err)
	}

	return locationResp, nil
}
