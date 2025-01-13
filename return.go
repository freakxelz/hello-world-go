package main

import "fmt"

// return single value
func getName(name string) string {
	return "hallo " + name
}

// return multiple value
func getFullName() (string, string) {
	return "Deva", "Mahendra"
}

// named return value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Asyifa"
	middleName = "Adya"
	lastName = "Maullidya"
	return
}

func main() {
	result := getName("Deva")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
