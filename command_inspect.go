package main

import (
	"errors"
	"fmt"
)


func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Provide a pokemon to inspect: inspect pidgey")
	}

	pokemon := args[0]
	p, exists := cfg.Caught[pokemon]
	if exists {
		fmt.Printf("Name: %s\n", p.Name) // Name
		fmt.Printf("Weight: %d\n", p.Weight) // Weight
		fmt.Printf("Height: %d\n", p.Height) // Height
		fmt.Println("Stats:") // Stats: ->
		for _, stat := range(p.Stats) {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:") // Types: ->
		for _, t := range(p.Types) {
			fmt.Printf(" - %s\n", t.Type.Name)
		}
	} else {
		fmt.Println("You haven't caught that pokemon yet")
	}
	return nil
}