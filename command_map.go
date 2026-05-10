//////////////////////////////////////
/*
1. Make an HTTP GET request to https://pokeapi.co/api/v2/location-area/.
2. Read the JSON response.
3. Unmarshal it into a Go struct.
4. Loop through the results and print each name. 
*/
//////////////////////////////////////

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	locAreaURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
)

type Heartattack struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// --------------------------------
// Map Forward
func commandMap(cfg *config) error {
	url := ""
	if cfg.Next == nil {
		url = locAreaURL
	} else {
		url = *cfg.Next
	}
	// Get response from URL
	heart, err := locationHelper(url, cfg)
	if err != nil {
		return err
	}
	// Print the map results
	for _, result := range heart.Results {
		fmt.Println(result.Name)
	}
	// Update the config data through the pointer in the parameter
	cfg.Next = heart.Next
	cfg.Previous = heart.Previous
	return nil
}

// --------------------------------
// Map Backward
func commandMapb(cfg *config) error {
	url := ""
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	} else {
		url = *cfg.Previous
	}
	heart, err := locationHelper(url, cfg)
	if err != nil {
		return err
	}
	// Print the map results
	for _, result := range heart.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = heart.Next
	cfg.Previous = heart.Previous
	return nil
}

// --------------------------------
// location-area page fetcher... helper
func locationHelper(url string, cfg *config) (Heartattack, error) {
	// Cache check hit = unmarshal and return
	fmt.Println("cache key:", url)
	if val, ok := cfg.cache.Get(url); ok {
		fmt.Println(">> Cached data found! <<")
		var h Heartattack
		if err := json.Unmarshal(val, &h); err != nil {
		return Heartattack{}, fmt.Errorf("Error unmarshalling body: %w", err)
		}
		return h, nil
	}
	// No cache hit, send a request to the url
	fmt.Println("<< No cached data. Sending request to PokeAPI... >>")
	res, err := http.Get(url)
	if err != nil {
		return Heartattack{}, fmt.Errorf("Failed to GET: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Heartattack{}, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return Heartattack{}, fmt.Errorf("Error: %w", err)
	}
	// Store cache
	cfg.cache.Add(url, body)
	
	var h Heartattack
	if err := json.Unmarshal(body, &h); err != nil {
		return Heartattack{}, fmt.Errorf("Error unmarshalling body: %w", err)
	}
	return h, nil
}