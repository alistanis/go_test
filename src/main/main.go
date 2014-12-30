package main

import (
	"fmt"
	"../points"
)

func main() {
	fmt.Println("Hello!")
	var p points.Point = points.NewPoint("CartesianPoint");
	p.SetX(10.12654);
	p.SetY(20.559651651);
	
	p.Print();

	fmt.Println("%F", p.X())
	fmt.Println("%F", p.Y())
}