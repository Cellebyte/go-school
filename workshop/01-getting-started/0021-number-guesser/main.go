package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// number guessing game
func main() {
	// TODO: 1. generate a random number
	randNumber := int64(rand.Intn(100))
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number: ")
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
