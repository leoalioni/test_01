package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (cir *Circulo) area() float64 {
	area := math.Pi * cir.Radio * cir.Radio
	return area
}

func (cir *Circulo) perimetro() float64 {
	perimetro := 2 * math.Pi * cir.Radio
	return perimetro
}
