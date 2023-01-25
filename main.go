package main

import (
	"fmt"
	"strings"
)

// naked return
func lenAndUpper (name string) (length int, uppercase string) {

	// defer: function 내용이 다 실행되고 나서 실행된다
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase=strings.ToUpper(name)
	// =를 쓰는 이유 : 값 업데이트 (생성x)
	return
	// return할 값을 명시하지 않아도 된다 (선언 문에서 작성했으므로 )
}

func main() {
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength,upperName)
}