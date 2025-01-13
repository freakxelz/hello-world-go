package main

import (
	"fmt"
)

func main() {

	// simple for looping

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// lebih rapih
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Pengulangan ke ", counter)
	}

	// for menggunakan data slice
	names := []string{
		"Deva Mahendra Ramadhan",
		"Asyifa Adya Maullidya",
		"Nayanika Adya Mahendra",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range menggunakan data slice
	// jika tidak ingin menggunakan variabel i atau
	// yang berisi nomor index, gunakan variabel "_"
	// for _, name := range names

	for _, name := range names {
		fmt.Println(name)
	}

	// menggunakan data array
	persons := make(map[string]string)
	persons["Name"] = "Deva Mahendra Ramadhan"
	persons["Title"] = "DevOps Engineer"

	for key, person := range persons {
		fmt.Println(key, "=", person)
	}
}
