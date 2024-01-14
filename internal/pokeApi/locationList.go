package pokeApi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/admiralyeoj/pokedexcli/internal/pokeApi/types"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (types.RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.RespShallowLocations{}, err
	}

	locationsResp := types.RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return types.RespShallowLocations{}, err
	}

	return locationsResp, nil
}
