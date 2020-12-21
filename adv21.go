package main

import (
	"fmt"
	"io/ioutil"
	// "strconv"
	// "regexp"
	"sort"
	"strings"
)
import "github.com/golang-collections/collections/set"

var ingredientToAllergen map[string]string

type ByAllergen []string

func (a ByAllergen) Len() int      { return len(a) }
func (a ByAllergen) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAllergen) Less(i, j int) bool {
	return ingredientToAllergen[a[i]] < ingredientToAllergen[a[j]]
}

func main() {
	reportLine, err := ioutil.ReadFile("adv21.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(reportLine)), "\n")

	allergensTranslations := make(map[string]*set.Set)
	safeIngredients := make(map[string]int)
	for _, line := range lines {
		parts := strings.Split(line, " (contains ")
		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(parts[1][:len(parts[1])-1], ", ")

		for _, ingredient := range ingredients {
			value := 1
			if v, exists := safeIngredients[ingredient]; exists {
				value = v + 1
			}
			safeIngredients[ingredient] = value
		}

		for _, allergen := range allergens {
			ingredientSet := set.New()
			for _, ingredient := range ingredients {
				ingredientSet.Insert(ingredient)
			}
			if v, exists := allergensTranslations[allergen]; !exists {
				allergensTranslations[allergen] = ingredientSet
			} else {
				allergensTranslations[allergen] = ingredientSet.Intersection(v)
			}
		}
	}

	translatedAllergens := make(map[string]string)
	ingredientToAllergen = make(map[string]string)
	unsafeIngredients := make([]string, 0)
	for translated := false; !translated; {
		translated = true
		for allergen, ingredients := range allergensTranslations {
			if ingredients.Len() == 1 {
				var ingredient string
				ingredients.Do(func(data interface{}) {
					ingredient = data.(string)
				})
				translatedAllergens[allergen] = ingredient
				unsafeIngredients = append(unsafeIngredients, ingredient)
				ingredientToAllergen[ingredient] = allergen
				delete(allergensTranslations, allergen)
				set := set.New(ingredient)
				delete(safeIngredients, ingredient)
				for k, v := range allergensTranslations {
					allergensTranslations[k] = v.Difference(set)
				}
			} else {
				translated = false
			}
		}
	}
	sort.Sort(ByAllergen(unsafeIngredients))

	var partOne int
	for _, v := range safeIngredients {
		partOne += v
	}
	// fmt.Println(safeIngredients)
	partTwo := strings.Join(unsafeIngredients, ",")
	/*
		for _, ingredient := range unsafeIngredients {
			fmt.Printf("%s contains %s\n", ingredient, ingredientToAllergen[ingredient])
		}
	*/

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
