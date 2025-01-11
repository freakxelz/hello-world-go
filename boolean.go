package main

import "fmt"

func main(){

	var(
	nilaiAkhir = 90
	absensi = 80

	lulusUjian = nilaiAkhir > 80
	lulusAbsensi = absensi > 75

	lulus = lulusUjian && lulusAbsensi
	)

	fmt.Println(lulus)

}
