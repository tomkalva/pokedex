package pokeapi

import (
	"encoding/json"
	"io"
)

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c Client) GetLocationAreas(pageURL *string) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var data RespShallowLocations
	err = json.Unmarshal(body, &data)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return data, nil
}
