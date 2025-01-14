package main

import "fmt"

// basic function
func sayHalo(firstName string, lastName string) {
	fmt.Println("Halo!", firstName, lastName)
}

func whoIs(firstName string, lastName string, umur int) {
	fmt.Println("Ini adalah", firstName, lastName, "dengan umur", umur, "Tahun")
}

// variadic function
func sumAll(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// function as parameter
func helloWithFilter(name string, filter func(string) string) {
	nameFilltered := filter(name)
	fmt.Println("helo", nameFilltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "****"
	} else {
		return name
	}
}

// anonymous function
type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You're Blocked", name)
	} else {
		fmt.Println("Welcome Back", name)
	}
}

func main() {
	sayHalo("Deva", "Mahendra")
	sayHalo("Asyifa", "Adya")

	whoIs("Nayanika", "Adya", 1)

	sum := sumAll(19, 29, 99, 00)
	fmt.Println(sum)

	helloWithFilter("Deva", spamFilter)
	helloWithFilter("anjing", spamFilter)

	blackList := func(name string) bool {
		return name == "Admin"
	}

	registerUser("Admin", blackList)
	registerUser("asyifa", blackList)
	registerUser("deva", blackList)
}
