package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		println("Pokedex >")
		for scanner.Scan() {

			inputTxT := scanner.Text()
			words := cleanInput(inputTxT)
			fmt.Println("Your command was:", words[0])
		}
	}
}
