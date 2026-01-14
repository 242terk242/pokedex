package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/242terk242/pokedex/commands"
	"github.com/242terk242/pokedex/repl"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		switch repl.CleanInput(scanner.Text())[0] {
		case "exit":
			commands.Exit()
		case "help":
			commands.Help()
		default:
			fmt.Println("Unknown command.")
		}

	}

}
