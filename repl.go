package main

// strings.ToLower() to convert a string to lowercase 
// and strings.TrimSpace() to remove leading and trailing whitespace
// strings.Fields() function, which divides the string into substrings
// by removing any spaces, including newlines, and returns a slice of the resulting substrings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"pokedexler/internal/pokecache"
)

func startRepl(cfg *config) {
	// Create the cache
	cfg.cache = pokecache.NewCache(2 * time.Minute)
	// Set up the input scanner
	scanner := bufio.NewScanner(os.Stdin)
	// Create Pokedex REP(Loop)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// If user presses enter with nothing:
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		
		proper := input[0]
		command, ok := getCommands()[proper]
		if ok {
			args := input[1:]
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command was invalid")
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: " Display 20 location areas in the world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Discover pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"keys": {
			name:        "DEBUG: keys",
			description: "Returns cached urls",
			callback:    commandKeys,
		},
		"dexler": {
			name:        "DEBUG: dexler",
			description: "Returns registered pokemon from Pokedex",
			callback:    commandDexler,
		},
	}
}

type config struct {
	cache      *pokecache.Cache
	Next       *string
	Previous   *string
	Caught     map[string]Pokeman
}
