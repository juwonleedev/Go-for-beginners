package main

import "fmt"

// 작성할 패키지 명

func main() {
	const name string = "nico"
	// go 언어는 type 언어여서 타입 명시
	
	var name2 string = "nico"
	name2="lynn"
	fmt.Println(name2)

	// 간편 작성 : go가 type 추론 후 추론한 타입에서 변경은 불가 
	name3:="nico"
	//name3:=false 이렇게 수정 불가 
	// 축약형: func 안에서만 + 변수에 한정해서 가능 
	fmt.Println(name3)
}