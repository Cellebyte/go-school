package file

import (
	"encoding/json"
	"fmt"
	"goschool/0026-recipe-cli-exercise/api"
	"os"
)

func SaveToFile(meal api.Meal) error {
	fileContent, err := json.MarshalIndent(meal, "", "  ")
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s.json", meal.IDMeal)
	err = os.WriteFile(filename, fileContent, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadFromFile(filename string) (api.Meal, error) {
	var meal api.Meal
	var err error
	data, err := os.ReadFile(filename)
	if err != nil {
		return meal, err
	}
	json.Unmarshal(data, &meal)
	return meal, err
}
