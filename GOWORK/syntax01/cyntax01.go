package main

import (
	"fmt"     // format 관련 패키지
	"math"    // 수학관련 패키지
	"strings" // 문자열 패키지
)

func main() { //진입점이 되는 함수
	var c rune = '가' // 44032 -> 가의 유니코드 출력 | "" : 문자열 / '' : rune
	//var a int16 = 7
	//var a = 7
	//a := 7
	a := 7
	var b float64 = 5.3
	a = int(b) // Type conversion, Casting 형변환 5가 출력
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        // 유니코드(UTF-8) 값 출력
	fmt.Printf("%c\n", c) // 한 글자 출력
	fmt.Printf("%T\n", c) // rune은 결국 int32의 별명

	fmt.Println(math.Round(2.71))
	fmt.Println("Hello GO~")
	fmt.Println(strings.Title("go git github java"))
}
