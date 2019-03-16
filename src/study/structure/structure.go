package main

import (
	"fmt"
	"sort"
)

func main() {
	//slice()
	mapFunc()
}

func slice() {
	var s []string
	s = append(s, "1", "2", "a", "b", "c")
	fmt.Println(s)
	fmt.Println(s[:])
	fmt.Println(s[3:])
	fmt.Println(s[:3])

	for i, v := range s {
		fmt.Printf("index:%d, value:%s", i, v)
		fmt.Println()
	}
}

func mapFunc() {
	args := make(map[string]int)
	strs := make([]string, 0, 12)
	args["a"] = 1
	args["b"] = 2
	args["c"] = 3
	args["d"] = 4
	args["e"]++
	for k, v := range args {
		strs = append(strs, k)
		fmt.Printf("key:%s, value:%d", k, v)
		fmt.Println()
	}
	v, ok := args["e"]
	if ok {
		fmt.Println(v)
	}
	for _, v := range strs {
		fmt.Print(v)
	}
	fmt.Println()
	sort.Strings(strs)
	for _, v := range strs {
		fmt.Print(v)
	}

}
