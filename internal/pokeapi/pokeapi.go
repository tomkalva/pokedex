package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type DetailedResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c Client) GetLocationAreas(pageURL *string) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	rawBytes, ok := c.cache.Get(url)

	if !ok {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return RespShallowLocations{}, fmt.Errorf("failed to fetch location areas")
		}

		rawBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		c.cache.Add(url, rawBytes)
	}

	var decoded RespShallowLocations
	err := json.Unmarshal(rawBytes, &decoded)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return decoded, nil
}

func (c Client) GetLocationArea(name string) (DetailedResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + name + "/"

	rawBytes, ok := c.cache.Get(url)

	if !ok {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return DetailedResponse{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return DetailedResponse{}, fmt.Errorf("location area not found")
		}

		rawBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return DetailedResponse{}, err
		}
		c.cache.Add(url, rawBytes)
	}

	var decoded DetailedResponse
	err := json.Unmarshal(rawBytes, &decoded)
	if err != nil {
		return DetailedResponse{}, err
	}

	return decoded, nil
}

func (c Client) GetPokemonStats(name string) (PokemonResponse, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name + "/"

	res, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PokemonResponse{}, fmt.Errorf("pokemon not found")
	}

	rawBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	var decoded PokemonResponse
	err = json.Unmarshal(rawBytes, &decoded)
	if err != nil {
		return PokemonResponse{}, err
	}

	return decoded, nil
}
