package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("You need to provide a pokemon to try and catch")
	}

	fullURL := baseURL + "/pokemon/" + args[0]
	pokemon, err := getPokemonJSON(fullURL, cfg)
	if err != nil {
		return fmt.Errorf("Error at poking: %w", err)
	}
	name := pokemon.Name
	bxp := pokemon.BaseExperience
	threshold := 100
	chance := float64(threshold) / float64(bxp)
	if chance > 0.75 {
		chance = 0.75
	}
	if chance < 0.20 {
		chance = 0.20
	}
	percent := chance * 100
	effectiveThreshold := int(chance * float64(bxp))
	roll := rand.Intn(bxp)
	result := roll + effectiveThreshold
	caught := result >= bxp
	// UX
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	fmt.Printf(
		"Chance to catch: %.0f%% (added threshold: %d)\n - Base EXP: %d\n - Roll: %d\n - Result: %d\n\n",
		percent,
		effectiveThreshold,
		bxp,
		roll,
		result,
	)
	if caught {
		fmt.Printf("%s was caught!\n", name)
		_, exists := cfg.Caught[name]
		if exists {
			fmt.Printf("%s is already registered in your Pokedex.\nInspect it with the inspect command.\n", name)
		} else {
			cfg.Caught[name] = pokemon
			fmt.Printf("%s has been added to your Pokedex!\nYou may now inspect it with the inspect command.\n", name)
		}
		return nil
	}
	fmt.Printf("%s escaped!\n", name)
	return nil
}

func getPokemonJSON(url string, cfg *config) (Pokeman, error) {
	var p Pokeman
	if val, ok := cfg.cache.Get(url); ok {
		// fmt.Printf("\n>> Cached data found at: %s\n", url)
		if err := json.Unmarshal(val, &p); err != nil {
			return p, fmt.Errorf("Error unmarshalling cached data: %w", err)
		}
		return p, nil
	}
	// fmt.Println(">> No cached data. Sendint GET request to PokeAPI...")
	res, err := http.Get(url)
	if err != nil {
		return p, fmt.Errorf("Failed to GET: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return p, fmt.Errorf("Error in body: %w", err)
	}
	if res.StatusCode > 299 {
		return p, fmt.Errorf("Response failed: %w", err)
	}

	if err := json.Unmarshal(body, &p); err != nil {
		return p, fmt.Errorf("Error unmarshalling body: %w", err)
	}

	// Cache pokemon endpoint
	cfg.cache.Add(url, body)
	return p, nil
}