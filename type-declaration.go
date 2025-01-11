package main

import "fmt"

func main(){

	type(
	nama string
	jenisKelamin bool
	noHp string
	status bool
	)

	var(
	nama1 nama = "Asyifa Adya Maullidya"
	JKelamin1 jenisKelamin = false
	no1 noHp = "081234567"
	status1 status = true

	nama2 nama ="Deva Mahendra Ramadhan"
	JKelamin2 jenisKelamin = true
	no2 noHp = "082384908"
	status2 status = true
	)

	fmt.Println(nama1, JKelamin1, no1, status1)
	fmt.Println(nama2, JKelamin2, no2, status2)
}
