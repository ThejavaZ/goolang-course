package main

import (
	"fmt"
	"strings"
)

// 1. Función básica (Parámetros con tipo)
func saludar(nombre string) {
	fmt.Printf("Hola, %s! Bienvenido al módulo de funciones.\n", nombre)
}

// 2. Función con retorno simple
func sumar(a int, b int) int {
	return a + b
}

// 3. RETORNO MÚLTIPLE (La joya de Go)
// Útil para devolver [Resultado, Error] o como en este caso [Minúscula, Mayúscula]
func transformar(texto string) (string, string, int) {
	min := strings.ToLower(texto)
	may := strings.ToUpper(texto)
	largo := len(texto)
	return min, may, largo
}

// 4. Funciones con "Named Return" (Retornos nombrados)
// Declaramos las variables de salida en la firma de la función
func calcularRectangulo(base, altura float64) (area float64, perimetro float64) {
	area = base * altura
	perimetro = (base + altura) * 2
	// El return solo (naked return) devuelve las variables nombradas
	return
}

func funciones() {
	saludar("Javier")

	// Uso de retorno simple
	resultadoSumar := sumar(10, 20)
	fmt.Println("Suma de 10 + 20 =", resultadoSumar)

	// Uso de RETORNO MÚLTIPLE
	// Recibimos los 3 valores en variables distintas
	textoMin, textoMay, total := transformar("Golang")
	fmt.Printf("Min: %s | May: %s | Letras: %d\n", textoMin, textoMay, total)

	// Si solo queremos un valor, usamos el Blank Identifier (_)
	_, soloMayus, _ := transformar("Fedora")
	fmt.Println("Solo nos importa la mayúscula:", soloMayus)

	// Uso de retornos nombrados
	a, p := calcularRectangulo(10.5, 5.2)
	fmt.Printf("Rectángulo -> Área: %.2f | Perímetro: %.2f\n", a, p)

	// 5. Funciones Anónimas (Lambda style)
	multiplicar := func(x, y int) int {
		return x * y
	}
	fmt.Println("Multiplicación anónima (5x5):", multiplicar(5, 5))
}