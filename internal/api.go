package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const API_BASE_PATH = "https://pokeapi.co/api/v2"
const LOCATION_AREA_PATH = API_BASE_PATH + "/location-area"

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

func GetLocationArea(pageUrl *string) (LocationArea, error) {
	var locationArea LocationArea
	url := LOCATION_AREA_PATH
	cache := NewCache(5 * time.Second)

	if pageUrl != nil {
		url = *pageUrl
	}

	if res, ok := cache.Get(url); ok {
		err := json.Unmarshal(res, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return LocationArea{}, err
		}

		err = json.Unmarshal(body, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

        cache.Add(url, body)
		return locationArea, nil
	}
}
