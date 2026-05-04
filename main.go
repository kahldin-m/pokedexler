package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands()
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}


func getCommands() map[string]cliCommand {
	commandList := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commandList
}


func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputText := scanner.Text()
		cleaned := cleanInput(inputText)
		if len(cleaned) == 0 {
			continue
		}
		proper := cleaned[0]
		
		_, ok := commands[proper]
		if ok {
			cmd := commands[proper]
			cmd.callback()
		} else {
			fmt.Println("Command was invalid")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}