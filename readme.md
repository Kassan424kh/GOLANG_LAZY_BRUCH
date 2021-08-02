### بسم الله الرحمن الرحيم


Cloned this nice tool from https://github.com/dulnan/lazy-brush, 🌴 Thanks alot to [dulnan](https://github.com/dulnan)


Sample example how to use it:
```GOLANG

package main

import (
	"fmt"

	LB "github.com/Kassan424kh/golang-lazy-brush/lazy-brush"
	LP "github.com/Kassan424kh/golang-lazy-brush/lazy-point"
	P "github.com/Kassan424kh/golang-lazy-brush/point"
)

func main() {
	lb := LB.Init()

	lb.Update(P.Point{X: 50, Y: 0}, false)
	fmt.Println(lb.GetBrushCoordinates())

	lb.Update(P.Point{X: 200, Y: 50}, false)
	fmt.Println(lb.GetBrushCoordinates())
}
```
