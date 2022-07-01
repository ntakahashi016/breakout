package main

import "math"

type Vector interface {
	X() float64
	SetX(x float64)
	Y() float64
	SetY(y float64)
	Magnitude() float64
	Normalize() Vector
	Diff(other Vector) Vector
	Add(other Vector) Vector
}

func NewVector(x,y float64) Vector {
	return &vector{x: x, y: y}
}

type vector struct {
	x float64
	y float64
}

func (v *vector) X() float64 { return v.x }
func (v *vector) SetX(x float64) { v.x = x }
func (v *vector) Y() float64 { return v.y }
func (v *vector) SetY(y float64) { v.y = y }

func (v *vector) Normalize() Vector {
	absX := math.Abs(v.x)
	absY := math.Abs(v.y)
	if absX > absY {
		return NewVector(v.x/absX, v.y/absX)
	} else {
		return NewVector(v.x/absY, v.y/absY)
	}
}
func (v *vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *vector) Diff(other Vector) Vector {
	return NewVector(other.X()-v.x, other.Y()-v.y)
}

func (v *vector) Add(other Vector) Vector {
	return NewVector(other.X()+v.x, other.Y()+v.y)
}
