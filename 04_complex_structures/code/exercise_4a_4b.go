package main

import "fmt"

//func average(num1, num2, num3 float64) float64 {
//	total := num1 + num2 + num3
//	return total / 3
//}

func average(numbers ...float64) float64 {
	total := 0.0
	for _, value := range numbers {
		total += value
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Println(average(2.5, 6.5, 9.2))
}
