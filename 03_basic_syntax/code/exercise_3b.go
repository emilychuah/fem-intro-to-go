package main

import "fmt"

func main() {
	var name string
	var city string
	var lengthOfStay int
	var weather bool

	fmt.Println("What's your name? ")
	fmt.Scan(&name)

	fmt.Println("What city do you live in? ")
	fmt.Scan(&city)

	fmt.Println("How long have you lived there? ")
	fmt.Scan(&lengthOfStay)

	fmt.Println("Is the weather there nice? (true/false) ")
	fmt.Scan(&weather)

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t", name, city, lengthOfStay, weather)
}
