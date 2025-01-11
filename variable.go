package main

import "fmt"

func main(){

	//menggunakan kata kunci var untuk deklarasi variable
	var name string

	name = "Asyifa Adya Maullidya"
	fmt.Println(name)

	name = "Deva Mahendra Ramadhan"
	fmt.Println(name)

	//tanpa menggunakan kata kunci var untuk deklarasi variable
	//gunakan ":="
	address := "Bandung"
	fmt.Println(address)

	//multiple variable
	var(

	firstName = "Asyifa"
	middleName = "Adya"
	lastName = "Maullidya"

	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}

