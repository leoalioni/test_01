package figuras

import "fmt"

type Geometria interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometria) {
	fmt.Println("Medidas:", g)
	fmt.Println("Área:", g.area())
	fmt.Println("Perímetro", g.perimetro())
}
