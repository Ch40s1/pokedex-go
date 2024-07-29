package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	// create a new scanner
	reader := bufio.NewScanner(os.Stdin)
	// infinite for loop
	for {
		// this is the prompt
		fmt.Print("Pokedex > ")
		// scan the input
		reader.Scan()

		// clean the input
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		// get the command name
		commandName := words[0]

		// check if the command exists in the map
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// this function will clean the input from the user and return a slice of words
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show locations",
			callback:    getPokeApi,
		},
	}
}
