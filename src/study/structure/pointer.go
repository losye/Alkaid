package main

import "fmt"

func main() {
	test()
	x := 5
	zero1(x) // References cannot be reassigned
	fmt.Println(x)
	zero2(&x)
	fmt.Println(x)
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

func test() {

	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(p)  // memory address
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)
	fmt.Println(p)
}

func zero1(x int) {
	x = 0
}

func zero2(xPtr *int) {
	*xPtr = 0
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
