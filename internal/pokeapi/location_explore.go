package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ExploreLocation
func (c *Client) ExploreLocation(pageUrl, locationName string) (RespDetailedLocation, error) {
	url := baseUrl + "/location/area" + locationName

	var data []byte

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	apiData, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocation{}, err
	}
	// add cache on this line
	data = apiData

	locationResp := RespDetailedLocation{}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	return locationResp, nil
}
