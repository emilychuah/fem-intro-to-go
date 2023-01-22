package main

import "fmt"

// Part 1
//func average(scores [5]float64) float64 {
//	total := 0.0
//	for _, value := range scores {
//		total += value
//	}
//	return total / float64(len(scores))
//}
//
//
//func main() {
//	scores := [5]float64{2.5, 8.2, 4.0, 7.3, 5.6}
//	fmt.Println(average(scores))
//}

// Part 2
//func doesPetExist(petName string) bool {
//	_, ok := initialPets[petName]
//	return ok
//}
//
//var initialPets map[string]string = map[string]string{
//	"fido":  "dog",
//	"kitty": "cat",
//}
//
//func main() {
//	pet := "fido"
//	petExists := doesPetExist(pet)
//	fmt.Println(petExists)
//}

// Part 3
var initialGroceries = []string{"beans", "lemons", "chicken", "fruit"}

func addGroceryToList(newGroceries ...string) []string {
	foods := initialGroceries
	for _, g := range newGroceries {
		foods = append(foods, g)
	}
	return foods
}

func main() {
	groceryList := addGroceryToList("beets", "chocolate", "wine")
	fmt.Println(groceryList)
}
