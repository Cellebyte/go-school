package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseApi = "https://themealdb.com/api/json/v1/1/"
const api = "https://www.themealdb.com/api/json/v1/1/lookup.php?i=52772"

func URL(path string) url.URL {
	api, _ := url.Parse(baseApi)
	api = api.JoinPath(path)
	return *api
}

func Get[T any](uri url.URL, object *T) error {
	fmt.Println(uri.String())
	resp, err := http.Get(uri.String())
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(object); err != nil {
		return err
	}
	fmt.Println(object)
	return nil
}

func GetRandomMeal() (Meal, error) {
	var meal Meal
	var mealsResponse mealsResponse
	err := Get(URL("random.php"), &mealsResponse)
	if err != nil {
		return meal, err
	}

	if len(mealsResponse.Meals) == 0 {
		return meal, fmt.Errorf("no meals found")
	}

	return mealsResponse.Meals[0], nil
}

func GetMealById(id string) (Meal, error) {
	var mealsResponse mealsResponse
	uri := URL("lookup.php")
	values := uri.Query()
	values.Set("i", id)
	uri.RawQuery = values.Encode()
	err := Get(uri, &mealsResponse)
	if err != nil {
		return Meal{}, err
	}
	return mealsResponse.Meals[0], nil
}

func SearchMealsByName(name string) ([]Meal, error) {
	var mealsResponse mealsResponse
	uri := URL("search.php")
	values := uri.Query()
	values.Set("s", name)
	uri.RawQuery = values.Encode()
	err := Get(uri, &mealsResponse)
	if err != nil {
		return nil, err
	}
	return mealsResponse.Meals, nil
}
