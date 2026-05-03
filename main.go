package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputText := scanner.Text()
		cleaned := cleanInput(inputText)
		result := fmt.Sprintf("Your command was: %s", cleaned[0])
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}