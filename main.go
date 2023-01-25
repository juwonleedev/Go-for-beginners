package main

import (
	"fmt"
)

func main () {
	a:=2
	b:=a // 값이 복사됨 
	a=10
	c:=&a
	// & -> 메모리 주소로 안내
	// * -> 메모리가 가리키는 값을 안내 (주소에 담긴 값)
	fmt.Println(a,b,c,*c)
	// 10, 2, 메모리값, 10
	*c = 20
	// *를 사용해서 a 값을 업데이트 
	// *주소 = 없데이트할 값 
	fmt.Println(a,b,c,*c)
	// 20 2 0xc0000160b8 20
}