package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Location name required to explore!")
	}
	if len(args) > 1 {
		return errors.New("Too many args for explore command")
	}

	fullURL := baseURL + "/location-area/" + args[0]
	pokemans, err := exploreHelper(fullURL, cfg)
	if err != nil {
		return fmt.Errorf("Error at chumps: %w", err)
	}
	fmt.Println("Found Pokemon:")
	for _, i := range pokemans.PokemonEncounters {
		chump := " - " + i.Pokemon.Name
		fmt.Println(chump)
	}
	return nil
}

func exploreHelper(url string, cfg *config) (Pokechumps, error) {
	if val, ok := cfg.cache.Get(url); ok {
		// fmt.Printf(">> Cached data found at: %s\n", url)
		var p Pokechumps
		if err := json.Unmarshal(val, &p); err != nil {
			return Pokechumps{}, fmt.Errorf("Error unmarshalling cached data: %w", err)
		}
		return p, nil
	}
	// fmt.Println(">> No cached data. Sending request to PokeAPI...")
	res, err := http.Get(url)
	if err != nil {
		return Pokechumps{}, fmt.Errorf("Failed to GET: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokechumps{}, fmt.Errorf("Error in body: %w", err)
	}
	if res.StatusCode > 299 {
		return Pokechumps{}, errors.New("Response failed")
	}

	var p Pokechumps
	if err := json.Unmarshal(body, &p); err != nil {
		return Pokechumps{}, fmt.Errorf("Error unmarshalling body: %w", err)
	}
	// Cache url
	cfg.cache.Add(url, body)
	return p, nil
}
