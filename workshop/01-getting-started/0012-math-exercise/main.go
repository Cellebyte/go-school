package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// calculate the Euclidean distance between two points
	// (x1, y1) and (x2, y2)
	// x1, y1 := 0.0, 0.0
	// x2, y2 := 3.0, 4.0

	// dist := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	// fmt.Println(dist)
	// ref: https://tutorialedge.net/golang/reading-console-input-golang/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the radius of the circle: ")
	fmt.Println("-----------------------------")

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	r, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	d := r * 2
	c := 2 * math.Pi * r
	a := math.Pi * math.Pow(r, 2)

	fmt.Printf("Diameter: %.2f\n", d)
	fmt.Printf("Circumference: %.2f\n", c)
	fmt.Printf("Area: %.2f\n", a)
}
