package pokeApi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ResourceList(resourceId string) (ResourceList, error) {

	url := baseURL + "/location-area/" + resourceId

	if resourceId == "" {
		return ResourceList{}, errors.New("Requires a location name")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResourceList{}, err
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := ResourceList{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return ResourceList{}, err
		}

		return locationsResp, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResourceList{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResourceList{}, err
	}

	resourceResp := ResourceList{}
	err = json.Unmarshal(dat, &resourceResp)
	if err != nil {
		return ResourceList{}, err
	}

	c.cache.Set(url, dat)
	return resourceResp, nil
}
