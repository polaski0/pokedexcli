package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/polaski0/pokedexcli/internal/cache"
)

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

var locationCache cache.Cache = cache.NewCache(10 * time.Second)
var exploreCache cache.Cache = cache.NewCache(10 * time.Second)

func getRequest[T any](pageUrl string, c cache.Cache) (*T, error) {
	var t *T

	if res, ok := c.Get(pageUrl); ok {
		err := json.Unmarshal(res, &t)
		if err != nil {
			return t, err
		}

		return t, nil
	} else {
		res, err := http.Get(pageUrl)
		if err != nil {
			return t, err
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return t, err
		}

		err = json.Unmarshal(body, &t)
		if err != nil {
			return t, err
		}

		c.Add(pageUrl, body)
		return t, nil
	}
}

func GetLocation(pageUrl *string) (LocationArea, error) {
	url := LOCATION_AREA_PATH

	if pageUrl != nil {
		url = *pageUrl
	}

	val, err := getRequest[LocationArea](url, locationCache)

	return *val, err
}

type ExploreArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationByArea(area *string) (ExploreArea, error) {
    if area == nil {
        return ExploreArea{}, nil
    }

    url := LOCATION_AREA_PATH + "/" + *area

	val, err := getRequest[ExploreArea](url, exploreCache)

	if err != nil {
		return ExploreArea{}, err
	}

	return *val, err
}
