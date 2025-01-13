package main

import "fmt"

func main() {

	// if dan else if
	var name string = "Deva"

	if name == "Deva" {
		fmt.Println("Halo, Deva")
	} else if name == "Asyifa" {
		fmt.Println("Halo Syifa")
	} else if name == "Nayanika" {
		fmt.Println("Halo, Aya")
	} else {
		fmt.Println("Kamu siapa?")
	}

	// switch
	mataPelajaran := "Fisika"

	switch mataPelajaran {
	case "Fisika":
		fmt.Println("Mata Pelajaran Fisika")
	case "Matematika":
		fmt.Println("Mata Pelajaran Matematikaa")
	case "Biologi":
		fmt.Println("Mata Pelajaran Biologi")
	case "Kimia":
		fmt.Println("Mata Pelajaran Kimia")
	default:
		fmt.Println("Mata Pelajaran Tidak Ada")
	}

	var nilai int = 95

	switch {
	case nilai >= 80:
		fmt.Println("Nila Kamu, A")
	case nilai >= 70:
		fmt.Println("Nila Kamu, B")
	case nilai >= 60:
		fmt.Println("Nila Kamu, C")
	case nilai >= 50:
		fmt.Println("Nila Kamu, D")
	default:
		fmt.Println("Nila Kamu, E")
	}
}
