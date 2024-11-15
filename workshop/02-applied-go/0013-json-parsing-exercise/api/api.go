package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type mealsResponse struct {
	Meals []Meal `json:"meals"`
}

func GetRandomMeal() (Meal, error) {
	var meal Meal
	const api = "https://themealdb.com/api/json/v1/1/random.php"
	resp, err := http.Get(api)
	if err != nil {
		return meal, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return meal, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var meals mealsResponse
	// TODO parse json into &meal
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return meal, fmt.Errorf("cannot read data from body")
	}
	err = json.Unmarshal(data, &meals)
	if err != nil {
		return meal, fmt.Errorf("cannot parse data from body")
	}
	mealIndex := rand.Intn(len(meals.Meals))

	meal = meals.Meals[mealIndex]

	return meal, nil
}
