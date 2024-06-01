package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(pageUrl *string) (LocationArea, error) {
	var val LocationArea
	url := LOCATION_AREA_PATH

	if pageUrl != nil {
		url = *pageUrl
	}

	if res, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(res, &val)
		if err != nil {
			return val, err
		}

		return val, nil
	} else {
		res, err := http.Get(url)
		if err != nil {
			return val, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return val, err
		}

		err = json.Unmarshal(body, &val)
		if err != nil {
			return val, err
		}

		c.cache.Add(url, body)
		return val, nil
	}
}

func (c *Client) GetLocationByArea(area *string) (ExploreArea, error) {
	var val ExploreArea

	if area == nil {
		return val, errors.New("area was not given")
	}

	url := LOCATION_AREA_PATH + "/" + *area

	if res, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(res, &val)
		if err != nil {
			return val, err
		}

		return val, nil
	} else {
		res, err := http.Get(url)
		if err != nil {
			return val, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return val, err
		}

		err = json.Unmarshal(body, &val)
		if err != nil {
			return val, err
		}

		c.cache.Add(url, body)
		return val, nil
	}
}

func (c *Client) GetPokemon(pokemon *string) (Pokemon, error) {
	var val Pokemon

	if pokemon == nil {
		return val, errors.New("pokemon was not given")
	}

	url := POKEMON_PATH + "/" + *pokemon

	if res, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(res, &val)
		if err != nil {
			return val, err
		}

		return val, nil
	} else {
		res, err := http.Get(url)
		if err != nil {
			return val, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return val, err
		}

		err = json.Unmarshal(body, &val)
		if err != nil {
			return val, err
		}

		c.cache.Add(url, body)
		return val, nil
	}
}
