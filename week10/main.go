package main

import (
	"fmt"
	"week10/src/greeting"
	"week10/src/mymath"
)

func main() {
	fmt.Println(mymath.Mypower(2, 9))
	fmt.Println(mymath.Mypower(4, 3))
	greeting.Hi()
	greeting.Hello()
	fmt.Println(mymath.MyAbs(-99))
	fmt.Println(mymath.MyAbs(71))
}
