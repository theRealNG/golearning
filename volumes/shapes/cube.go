package shapes

import (
	"fmt"
	"math"
)

type Cube struct {
	Side float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

func (c Cube) String() string {
	return fmt.Sprintf("Cube ( side: %v )", c.Side)
}
