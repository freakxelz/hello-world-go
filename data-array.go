package main

import "fmt"

func main(){

	var name [3]string
	name[0] = "Deva Mahendra Ramadhan"
	name[1] = "Asyifa Adya Maullidya"
	name[2] = "Nayanika Adya Mahendra"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])


	//array langsung 1 baris
	var sureName = [3]string{
	"Deva",
	"Syifa",
	"Aya",
	}

	fmt.Println(sureName)

	//cek panjang array
	fmt.Println(len(sureName))
}
