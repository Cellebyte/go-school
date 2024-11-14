package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const upperBoundary = 100

// number guessing game
func main() {
	// TODO: 1. generate a random number
	randNumber := int64(rand.Intn(upperBoundary))
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the number guess between (0,%d): \n", upperBoundary)
	fmt.Println("-----------------------------")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		userValue, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Printf("Your input is not a number %q \n", text)
			continue
		}
		if userValue < randNumber {
			fmt.Println("Too low")
		} else if userValue > randNumber {
			fmt.Println("Too high")
		} else {
			fmt.Println("You win")
			break
		}

		// TODO: 2. read the user input

		// TODO: 3. convert the user input to an integer

		// TODO: 4. compare the user input with the random number
	}
}
