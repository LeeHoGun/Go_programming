package main

import (
	"fmt" // format 관련 패키지
	//"reflect"
	"strings"
)

func main() { //진입점이 되는 함수
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts) // @가 o로 바뀐 newTexts(Go Money~~)가 출력
	fmt.Println(newTexts)

	// // 변수명은 영문자로 시작해야된다.
	// // 영문 대문자의 경우 다른 패키지에서 접근할 수 있다. | 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능
	// var e string
	// var d bool
	// var c rune // "" : 문자열 / '' : rune
	// var b float64
	// var a int

	// // naming convention // 명명 규칙
	// var studentId string
	// var i int32 // i = index

	// //a := 7
	// fmt.Println(studentId)
	// fmt.Println(i)

	// // 0 = zero value 제로값 | 변수를 선언하고 변수에 대한 값을 할당하지 않았을 때 초기 값
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// // fmt.Printf("%T\n", d) // println은 변환문자열이 아니기 때문에 별도의 패키지가 필요함
	// // fmt.Printf("%T\n", e) // 타입을 알아내는 방법 1

	// fmt.Println(reflect.TypeOf(d)) // 타입을 알아내는 방법 2
	// fmt.Println(reflect.TypeOf(e))
}
