package main

import (
	"github.com/theRealNG/volumes/measurements"
	"github.com/theRealNG/volumes/shapes"
)

func main() {
	c := shapes.Cube{Side: 2.0}
	b := shapes.Box{Length: 1, Width: 2, Height: 3}
	s := shapes.Sphere{Radius: 4}

	measurements.CalculateVolume(c)
	measurements.CalculateVolume(b)
	measurements.CalculateVolume(s)
}
