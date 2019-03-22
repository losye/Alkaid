package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	var wheel Wheel
	wheel.X = 5
	wheel.Y = 10
	wheel.Radius = 15
	wheel.Spokes = 20
	jsonFuc()
}

func jsonFuc() {
	var err error

	data, err := json.Marshal([]Point{Point{1, 2}, Point{3, 4}, Point{5, 6}})
	if err != nil {
		log.Fatal("json Marshal error")
	}
	fmt.Printf("data: %s", data)

	data, err = json.Marshal(Point{1, 2})
	if err != nil {
		log.Fatal("json Marshal error")
	}
	fmt.Printf("data: %s \n", data)

	var p = Point{}
	if err != json.Unmarshal(data, &p){
		log.Fatal("json Unmarshal error")
	}
	fmt.Print(p)

}


type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
