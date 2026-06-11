package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.Caught) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some pokemon!")
		return nil
	}
	fmt.Println("Registered Pokemon:")
	for n, _ := range cfg.Caught {
		fmt.Printf(" - %s\n", n)
	}
	return nil
}