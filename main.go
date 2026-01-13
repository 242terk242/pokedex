package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/242terk242/pokedex/repl"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		fmt.Printf("Your command was: %s\n", repl.CleanInput(scanner.Text())[0])

	}

}
