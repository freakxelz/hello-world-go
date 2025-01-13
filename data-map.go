package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Asyifa Adya Maullidya",
	}

	fmt.Println(person)

	person["umur"] = "24 Tahun"

	fmt.Println(person)

	person["umur"] = "25 Tahun"

	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["Title"] = "Belajar Go lang"
	book["author"] = "Deva Mahendra"

	fmt.Println(book)

	book["Title"] = "The Power of Habit"

	fmt.Println(book)
}
