package pokeApi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationType, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationType{}, err
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationType{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationType{}, err
		}

		return locationsResp, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationType{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationType{}, err
	}

	locationsResp := LocationType{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationType{}, err
	}

	c.cache.Set(url, dat)
	return locationsResp, nil
}
