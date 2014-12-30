package points 

import (
	"fmt"
	"math"
)

type CartesianPoint struct {
	x, y float64
}

type PolarPoint struct {
	r, θ float64
}

func (p CartesianPoint) X() float64 { return p.x }
func (p CartesianPoint) Y() float64 { return p.y }
func (p *CartesianPoint) SetX(x float64) { p.x = x }
func (p *CartesianPoint) SetY(y float64) { p.y = y }

func (p PolarPoint) X() float64 { return p.r * math.Cos(p.θ) }
func (p PolarPoint) Y() float64 { return p.r * math.Sin(p.θ) }
func (p *PolarPoint) SetX(num float64) { p.r = num / math.Cos(p.θ) }
func (p *PolarPoint) SetY(num float64) { p.θ = num / math.Sin(p.θ) }


func (self CartesianPoint) Print() {
	fmt.Printf("(%f, %f)\n", self.x, self.y)
}

func (self PolarPoint) Print() {
	fmt.Printf("(%f, %f)\n", self.r, self.θ)
}

func NewPoint(type_string string) Point {
	var p Point
	if type_string == "CartesianPoint" { 
		p = &CartesianPoint{}
	} else {
		p = &PolarPoint{}
	}
	return p
}

type Point interface {
	Printer
	X() float64
	Y() float64
	SetX(x float64) 
	SetY(y float64) 
}

type Printer interface {
	Print()
}

