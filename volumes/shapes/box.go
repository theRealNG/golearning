package shapes

import "fmt"

type Box struct {
	Length float64
	Width  float64
	Height float64
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}

func (b Box) String() string {
	return fmt.Sprintf("Box ( %v * %v *%v)", b.Length, b.Width, b.Height)
}
