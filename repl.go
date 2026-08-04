package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
	"time"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	attempts      map[string]int
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var mover = config{
	pokeapiClient: pokeapi.NewClient(5*time.Second, 10*time.Second),
	Next:          nil,
	Previous:      nil,
	attempts:      make(map[string]int),
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays the next 20 maps",
		callback:    getNextMaps,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the previous 20 maps",
		callback:    getPrevMaps,
	},
	"explore": {
		name:        "explore",
		description: "Displays pokemon in area",
		callback:    explore,
	},
	"catch": {
		name:        "catch",
		description: "Attempts to catch pokemon",
		callback:    catch,
	},
}

func start() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		params := words[1:]

		if cmd, ok := commands[commandName]; ok {
			cmd.callback(&mover, params...)

		} else {
			fmt.Println("Unknown command")
		}
	}

}

func cleanInput(text string) []string {
	result := []string{}
	words := strings.Fields(text)
	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}

	return result
}

func commandExit(m *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(m *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	return nil
}

func getNextMaps(m *config, args ...string) error {
	const baseUrl = "https://pokeapi.co/api/v2/location-area/"

	url := baseUrl

	if m.Next != nil {
		url = *m.Next
	}

	res, err := m.pokeapiClient.GetData(&url)

	if err != nil {
		return err
	}

	for _, item := range res.Results {

		fmt.Println(item.Name)
	}

	m.Previous = res.Previous
	m.Next = res.Next

	return nil
}

func getPrevMaps(m *config, args ...string) error {
	const baseUrl = "https://pokeapi.co/api/v2/location-area/"

	if m.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *m.Previous

	res, err := m.pokeapiClient.GetData(&url)

	if err != nil {
		return err
	}

	for _, item := range res.Results {

		fmt.Println(item.Name)
	}

	m.Previous = res.Previous
	m.Next = res.Next

	return nil
}

func explore(m *config, area ...string) error {

	if len(area) == 0 {
		fmt.Println("Enter an area name")
		return nil
	}

	res, err := m.pokeapiClient.GetPokemon(area[0])

	if err != nil {
		return err
	}

	fmt.Printf("Exploring: %s...\n", area[0])
	fmt.Println("Pokemon in the area:")
	for _, item := range res.PokemonEncounters {
		fmt.Println(item.Pokemon.Name)
	}

	return nil

}

func catch(m *config, name ...string) error {

	fmt.Printf("Throwing a Pokeball at %s...\n", name[0])

	res, err := m.pokeapiClient.CatchPokemon(name[0])
	if err != nil {
		return err
	}
	chance := 1.0 - float64(res.Exp)/300.0 + float64(m.attempts[name[0]])*0.15
	m.attempts[name[0]] += 1

	roll := rand.Float64()

	if roll < chance {
		fmt.Printf("%s was caught!\n", name[0])
		m.pokeapiClient.Pokedex[res.Name] = res
	} else {
		fmt.Printf("%s escaped!\n", name[0])
	}

	return nil
}
