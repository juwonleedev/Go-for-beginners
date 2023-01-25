package main

import (
	"fmt"
)

func main () {
	// Arrays 
	names := [5]string{"nico","lynn","dal"}
	names[3] = "alalal"
	names[4] = "fdfdfd"

	//Slices = arrays without length 
	names2 := []string{"nico", "lynn","dal"}
	names2= append(names2,"irene")
	// append (slice, 추가할 값) 을 이용해 slice에 값 추가
	fmt.Println(names, names2)
}