package measurements

import "fmt"

type OfStructure interface {
	Volume() float64
}

func CalculateVolume(kind OfStructure){
	fmt.Printf("The volume calculated for our %s is: %f \n", kind, kind.Volume())
}
