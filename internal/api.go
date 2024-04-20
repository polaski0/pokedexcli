package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

const API_BASE_PATH = "https://pokeapi.co/api/v2"
const LOCATION_AREA_PATH = API_BASE_PATH + "/location-area"

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(pageUrl *string) (LocationArea, error) {
	url := LOCATION_AREA_PATH
	if pageUrl != nil {
		url = *pageUrl
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	err = json.Unmarshal(body, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, err
}
