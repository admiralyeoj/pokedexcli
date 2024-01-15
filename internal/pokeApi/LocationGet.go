package pokeApi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(resourceId string) (Location, error) {

	url := baseURL + "/location-area/" + resourceId

	if resourceId == "" {
		return Location{}, errors.New("Requires a location name")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := Location{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Location{}, err
		}

		return locationsResp, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	resourceResp := Location{}
	err = json.Unmarshal(dat, &resourceResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Set(url, dat)
	return resourceResp, nil
}
