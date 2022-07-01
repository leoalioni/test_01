package figuras

type Cuadrado struct {
	Lado float64
}

func (cuad *Cuadrado) area() float64 {
	var area float64 = cuad.Lado * cuad.Lado
	return area
}

func (cuad *Cuadrado) perimetro() float64 {
	perimetro := cuad.Lado * 4
	return perimetro
}
