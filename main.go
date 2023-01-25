package main

import (
	"fmt"
)

// struct : object와 비슷한데 class 개념 
// map = [key] value 
// map도 key-value와 사용 가능
// map에 대한 function들은 나중에 (추가, 조회 등)

func main () {
	nico := map[string]string{"name":"nico","age":"12"}
	fmt.Println(nico)

	for _,value := range nico {
		fmt.Println(value)
	}
}