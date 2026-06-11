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
