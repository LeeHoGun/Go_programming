package main

import "fmt"

func double(n1 *int, n2 *int) { // 두 수를 교환
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
func main() {
	a := 10
	b := 20

	c := &a // 정수형 타입의 pointer
	fmt.Printf("%d %d %d\n", a, b, c)
}
