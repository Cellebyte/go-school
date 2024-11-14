package main

import (
	"fmt"
	"strings"
)

type Ingredients map[string]int

type Recipe struct {
	Name string
	Ingredients
}

// Implement stringer here
func (i Ingredients) String() string {
	var ingredients []string
	for name, amoount := range i {
		ingredients = append(ingredients, fmt.Sprintf("%v: %vg", name, amoount))
	}
	return strings.Join(ingredients, ", ")
}

func (r Recipe) String() string {
	return fmt.Sprintf("%v: (%s)", r.Name, r.Ingredients.String())
}

func main() {
	cookBook := map[string]Recipe{
		"pizza": {
			Name:        "Pizza",
			Ingredients: map[string]int{"Flour": 500, "Water": 300, "Yeast": 10},
		},
		"steak": {
			Name:        "Steak",
			Ingredients: map[string]int{"Meat": 500, "Salt": 10, "Pepper": 10},
		},
		"pudding": {
			Name:        "Pudding",
			Ingredients: map[string]int{"Milk": 500, "Water": 100, "Sugar": 100, "Vanilla": 10},
		},
	}

	for _, recipe := range cookBook {
		// TODO
		fmt.Println(recipe.String())
	}
}
