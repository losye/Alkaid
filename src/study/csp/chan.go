package main

import (
	"fmt"
	"time"
)

func main() {
	//closeF()
	//pipeline()
	//pipeline1()
	pipeline2()
}

func channel() {

	done := make(chan struct{})

	go func() {
		done <- struct{}{}
	}()

	<-done
}

func pipeline() {
	nature := make(chan int)
	square := make(chan int)

	go func() {
		for x := 1; ; x++ {
			nature <- x
		}
	}()

	go func() {
		for ; ; {
			x := <-nature
			square <- x * x
		}

	}()

	for ; ; {
		fmt.Println(<-square)
	}

}

func pipeline1() {
	nature := make(chan int)
	square := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			nature <- x
		}
		close(nature)
	}()

	go func() {
		for x := range nature {
			square <- x * x
		}
		close(square)
	}()

	for x := range square {
		fmt.Println(x)
	}
}

func pipeline2() {
	nature := make(chan int)
	square := make(chan int)

	go counter(nature)
	go squarer(nature, square)
	printer(square)
}

func counter(out chan<- int) {
	func() {
		for x := 0; x < 100; x++ {
			out <- x
		}
		close(out)
	}()
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func closeF() {
	go func() {
		time.Sleep(time.Minute * 60)
	}()

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // if not close it will block in range c
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("done")

}

func test() {
	foo()
	bar := bar{}

	println(bar.name)
}

func foo() {

}

type bar struct {
	name string
}
