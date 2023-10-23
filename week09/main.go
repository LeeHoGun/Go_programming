package main

import "fmt"

func double(n1 *int, n2 *int) { // 두 수를 교환
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
func main() {
	a := 1
	b := 2
	c := &a // 정수형 타입의 pointer | var c *int = &a
	// fmt.Printf("%d %d %d\n", a, b, c)
	fmt.Printf("%d %d %d\n", a, b, *c)
}
