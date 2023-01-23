package main

import "fmt"

func main() {
	var name string
	var city string
	var lengthOfStay int
	var weatherBool bool

	fmt.Print("What's your name? ")
	fmt.Scan(&name)

	fmt.Print("What city do you live in? ")
	fmt.Scan(&city)

	fmt.Print("How long have you lived there? ")
	fmt.Scan(&lengthOfStay)

	fmt.Print("Is the weather there nice? ")
	fmt.Scan(&weatherBool)

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t", name, city, lengthOfStay, weatherBool)
}
