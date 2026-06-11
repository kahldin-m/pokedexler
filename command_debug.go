package main

import (
	"fmt"
)

// Prints the cached urls for your pleasure
func commandKeys(cfg *config, args []string) error {
	if len(cfg.cache.Keys()) == 0 {
		fmt.Println("The pokecache is currently empty")
	}
	for _, k := range cfg.cache.Keys() {
		fmt.Printf("Cached key: %s\n", k)
	}
	return nil
}

func commandDexler(cfg *config, args []string) error {
	if len(cfg.Caught) == 0 {
		fmt.Println("The Pokedex is empty. Go catch some pokemon!")
		return nil
	}
	fmt.Println("Registered Pokemon:")
	for n, _ := range cfg.Caught {
		fmt.Printf(" - %s\n", n)
	}
	return nil
}
