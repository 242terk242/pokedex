package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
}

var commands map[string]cliCommand

func Exit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func Help(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func Map(cfg *Config) error {
	url := ""
	if cfg.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = cfg.Next
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	type Response struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous any    `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}
	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	if response.Previous != nil {
		cfg.Previous = response.Previous.(string)
	} else {
		cfg.Previous = ""
	}
	cfg.Next = response.Next

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func Mapb(cfg *Config) error {
	url := ""
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url = cfg.Previous
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	type Response struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous any    `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}
	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	if response.Previous != nil {
		cfg.Previous = response.Previous.(string)
	} else {
		cfg.Previous = ""
	}
	if response.Next != "" {
		cfg.Next = response.Next
	} else {
		cfg.Next = ""
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    Exit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    Help,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    Map,
		},
		"mapb": {
			name:        "map",
			description: "Displays the previous names of 20 location areas in the Pokemon world",
			callback:    Mapb,
		},
	}
}
