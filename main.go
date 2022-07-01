package goarea

import "math"

// Pi é uma proporção numérica pela relação entre
// o perímetro de uma circunferência e seu diametro
const Pi = 3.1416

// Circ é responsável por calcular a área do circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * raio
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
