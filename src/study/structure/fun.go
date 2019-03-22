package main

import "fmt"

func main() {
	f := Closure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func Closure() func() int {
	var x int
	return func() int {
		x ++
		return x * x
	}
}
