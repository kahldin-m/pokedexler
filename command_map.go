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
	locAreaURL = "https://pokeapi.co/api/v2/location-area/"
)

type Heartattack struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Map Forward
func commandMap(cfg *config) error {
	url := ""
	if cfg.Next == nil {
		url = locAreaURL
	} else {
		url = *cfg.Next
	}
	// Get response from URL
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to GET: %w", err)
	}
	// Turn the bytes into json data
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	var heart Heartattack
	// Put the json data into a struct
	if err := json.Unmarshal(body, &heart); err != nil {
		return fmt.Errorf("Error unmarshalling body: %w", err)
	}
	for _, result := range heart.Results {
		fmt.Println(result.Name)
	}
	// Update the config data through the pointer in the parameter
	cfg.Next = heart.Next
	cfg.Previous = heart.Previous
	return nil
}

// Map Backward
func commandMapb(cfg *config) error {
	url := ""
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	} else {
		url = *cfg.Previous
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to GET: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	var heart Heartattack
	if err := json.Unmarshal(body, &heart); err != nil {
		return fmt.Errorf("Error unmarshalling body: %w", err)
	}
	for _, result := range heart.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = heart.Next
	cfg.Previous = heart.Previous
	return nil
}