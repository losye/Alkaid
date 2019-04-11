package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println(toString(errors.New("a error type val")))
}

func toString(in interface{}) string {
	switch val := in.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", val)
	case bool:
		if val {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case string:
		return val
	default:
		panic(fmt.Sprintf("unexcepted type %T:  %v", in, val))
	}
}
