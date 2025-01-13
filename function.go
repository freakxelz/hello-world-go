package main

import "fmt"

func sayHalo(firstName string, lastName string) {
	fmt.Println("Halo!", firstName, lastName)
}

func whoIs(firstName string, lastName string, umur int) {
	fmt.Println("Ini adalah", firstName, lastName, "dengan umur", umur, "Tahun")
}

func main() {
	sayHalo("Deva", "Mahendra")
	sayHalo("Asyifa", "Adya")

	whoIs("Nayanika", "Adya", 1)
}
