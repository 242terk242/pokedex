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
	cfg := &commands.Config{Next: "", Previous: ""}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		switch repl.CleanInput(scanner.Text())[0] {
		case "exit":
			commands.Exit(cfg)
		case "help":
			commands.Help(cfg)
		case "map":
			commands.Map(cfg)
		case "mapb":
			commands.Mapb(cfg)
		default:
			fmt.Println("Unknown command.")
		}

	}

}
