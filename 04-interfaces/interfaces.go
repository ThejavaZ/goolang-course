package main

import (
	"fmt"
	"math"
)

// 1. Definición de la Interfaz
// Por convención, si solo tiene un método, se suele llamar con el sufijo "er" (ej. Reader, Writer)
type Figura interface {
	Area() float64
	Perimetro() float64
}

// 2. Implementación 1: Rectángulo
type Rectangulo struct {
	Ancho, Alto float64
}

// Rectangulo implementa Figura porque tiene Area() y Perimetro()
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2*r.Ancho + 2*r.Alto
}

// 3. Implementación 2: Círculo
type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

// 4. POLIMORFISMO: Esta función no sabe qué es 'f', solo sabe que "sabe calcular su área"
func medir(f Figura) {
	fmt.Printf("Figura: %T | Área: %.2f | Perímetro: %.2f\n", f, f.Area(), f.Perimetro())
}

func main() {
	r := Rectangulo{Ancho: 10, Alto: 5}
	c := Circulo{Radio: 3}

	fmt.Println("--- Calculadora de Figuras ---")
	
	// Podemos pasar ambos a la misma función gracias a la interfaz
	medir(r)
	medir(c)

	// 5. La Interfaz Vacía: interface{} o 'any'
	// Es el equivalente a 'object' en C#. Puede guardar CUALQUIER cosa.
	var listaMagica []any
	listaMagica = append(listaMagica, "Hola", 42, true, r)

	fmt.Println("\nContenido de la lista 'any':")
	for _, elemento := range listaMagica {
		fmt.Printf("- Valor: %v | Tipo: %T\n", elemento, elemento)
	}
}