package dbcomparer

import (
	"fmt"

	"json-xml-basics/internal/dbreader"
)

// CompareDB сравнивает две базы данных рецептов и выводит отличия.
func CompareDB(oldRecipes, newRecipes *dbreader.Recipes) {
	compareCakes(oldRecipes.Cakes, newRecipes.Cakes)
}

// compareCakes сравнивает списки тортов из старой и новой базы данных.
// Выводит информацию о добавленных, удалённых и изменённых тортах.
func compareCakes(oldCakes, newCakes []dbreader.Cake) {
	oldCakeMap := make(map[string]dbreader.Cake)
	newCakeMap := make(map[string]dbreader.Cake)

	for _, cake := range oldCakes {
		oldCakeMap[cake.Name] = cake
	}
	for _, cake := range newCakes {
		newCakeMap[cake.Name] = cake
	}

	// Проверка удаленных и добавленных тортов
	for newCakeName := range newCakeMap {
		if _, exists := oldCakeMap[newCakeName]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", newCakeName)
		}
	}
	for oldCakeName := range oldCakeMap {
		if _, exists := newCakeMap[oldCakeName]; !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCakeName)
		}
	}

	// Проверка измененных тортов
	for oldCakeName, oldCake := range oldCakeMap {
		if newCake, exists := newCakeMap[oldCakeName]; exists {
			compareCakeDetails(oldCake, newCake)
		}
	}
}

// compareCakeDetails сравнивает детали двух тортов, таких как время готовки и ингредиенты.
// Выводит информацию об изменениях в деталях торта.
func compareCakeDetails(oldCake, newCake dbreader.Cake) {
	if oldCake.Time != newCake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, newCake.Time, oldCake.Time)
	}

	compareIngredients(oldCake.Name, oldCake.Ingredients, newCake.Ingredients)
}

// compareIngredients сравнивает списки ингредиентов для указанного торта из старой и новой базы данных.
// Выводит информацию о добавленных, удалённых и изменённых ингредиентах.
func compareIngredients(cakeName string, oldIngredients, newIngredients []dbreader.Ingredient) {
	oldIngredientMap := make(map[string]dbreader.Ingredient)
	newIngredientMap := make(map[string]dbreader.Ingredient)

	for _, ingredient := range oldIngredients {
		oldIngredientMap[ingredient.Name] = ingredient
	}
	for _, ingredient := range newIngredients {
		newIngredientMap[ingredient.Name] = ingredient
	}

	// Проверка удаленных и добавленных ингредиентов
	for newIngredientName := range newIngredientMap {
		if _, exists := oldIngredientMap[newIngredientName]; !exists {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredientName, cakeName)
		}
	}
	for oldIngredientName := range oldIngredientMap {
		if _, exists := newIngredientMap[oldIngredientName]; !exists {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIngredientName, cakeName)
		}
	}

	// Проверка измененных ингредиентов
	for oldIngredientName, oldIngredient := range oldIngredientMap {
		if newIngredient, exists := newIngredientMap[oldIngredientName]; exists {
			if oldIngredient.Count != newIngredient.Count {
				fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredientName, cakeName, newIngredient.Count, oldIngredient.Count)
			}
			if oldIngredient.Unit != newIngredient.Unit {
				if oldIngredient.Unit == "" {
					fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Unit, oldIngredientName, cakeName)
				} else if newIngredient.Unit == "" {
					fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Unit, oldIngredientName, cakeName)
				} else {
					fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredientName, cakeName, newIngredient.Unit, oldIngredient.Unit)
				}
			}
		}
	}
}
