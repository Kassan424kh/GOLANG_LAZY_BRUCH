package main

import (
	"fmt"

	LB "github.com/Kassan424kh/lazy-brush/lazy-brush"
	LP "github.com/Kassan424kh/lazy-brush/lazy-point"
	P "github.com/Kassan424kh/lazy-brush/point"
)

func main() {
	lb := LB.LazyBrush{
		Radius:   30,
		Pointer:  LP.LazyPoint{X: 0, Y: 0},
		Brush:    LP.LazyPoint{X: 0, Y: 0},
		Angle:    0,
		Distance: 0,
	}
	lb.Enable()

	lb.Update(P.Point{X: 50, Y: 0}, false)
	fmt.Println(lb.GetBrushCoordinates())

	lb.Update(P.Point{X: 200, Y: 50}, false)
	fmt.Println(lb.GetBrushCoordinates())
}
