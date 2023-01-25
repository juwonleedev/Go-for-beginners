package main

import (
	"fmt"
)

type person struct {
	name string
	age int
	favFood [] string 
}

func main () {
	favFood:=[]string{"kimchi","cake","ramen"}
	nico := person {"nico",18,favFood}
	fmt.Println(nico.name)

	// 혹은 이렇게 field:value로 표현
	nico2 := person {name:"nico",age:18,favFood: favFood}
	fmt.Println(nico2.age)
}