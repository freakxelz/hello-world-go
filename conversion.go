package main

import "fmt"

func main(){

	var(

	name = "Asyifa"
	a = name[0]
	y = name[2]
	aString string = string(a)
	yString string = string(y)

	name2 = "Deva"
	d = name2[0]
	v = name2[2]
	dString string = string(d)
	vString string = string(v)

	)

	fmt.Println(name)
	fmt.Println(aString)
	fmt.Println(yString)
	fmt.Println(yString , aString)
	fmt.Println(name2)
	fmt.Println(dString)
	fmt.Println(vString)
	fmt.Println(dString,vString)

}
