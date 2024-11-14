package main

import "fmt"

type Recipe struct {
	Name        string
	Ingredients map[string]int
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

	totalIngredients := make(map[string]int)

	// TODO: get the total number of ingredients
	for _, recipe := range cookBook {
		for ingredient, amount := range recipe.Ingredients {
			if storedAmount, ok := totalIngredients[ingredient]; ok {
				totalIngredients[ingredient] = amount + storedAmount
			} else {
				totalIngredients[ingredient] = amount
			}
		}
	}
	// TODO: print the total number of in
	fmt.Printf("%v", totalIngredients)
}
