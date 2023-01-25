package main

import (
	"fmt"
	"strings"
)

// 작성할 패키지 명
func multiply(a int,b int) int {
	return a*b
	// 필수적으로 타입 명시
}

func multiply2(a, b int) int{
	return a*b
}

// go는 리턴값을 여러개 반환 가능 
func lenAndUpper(name string) (int,string){
	return len(name),strings.ToUpper(name)
	// go는 표준 라이브러리가 잘 되어 있다
}

func repeatMe(words ...string){
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2,2))
	
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength,upperName)
	// go 는 사용되지 않는 변수에 대해 에러 발생
	// totalLength, _ := lenAndUppper("nico")로 리턴값 무시 가능 

	repeatMe("nico","love","cute")
}