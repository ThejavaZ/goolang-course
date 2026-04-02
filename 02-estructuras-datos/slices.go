package main

import "fmt"

func main() {
	// 1. Declaración de un Array (Tamaño fijo)
	// Casi no se usan en Go, pero son la base de los Slices.
	var arreglo [3]string
	arreglo[0] = "Linux"
	arreglo[1] = "Windows"
	arreglo[2] = "macOS"

	fmt.Println("Arreglo fijo:", arreglo)

	// 2. Declaración de un Slice (Tamaño dinámico)
	// Nota que NO ponemos número dentro de los corchetes [].
	lenguajes := []string{"Go", "Python", "C#"}
	fmt.Println("Slice inicial:", lenguajes)

	// 3. Agregar elementos (append)
	// IMPORTANTE: append devuelve un NUEVO slice, por eso reasignamos.
	lenguajes = append(lenguajes, "JavaScript", "Kotlin")
	fmt.Println("Después de append:", lenguajes)

	// 4. Capacidad vs Longitud (Concepto Avanzado)
	// len() = cuántos elementos hay.
	// cap() = cuánto espacio tiene reservado en memoria.
	numeros := make([]int, 3, 5) // largo 3, capacidad 5
	fmt.Printf("Largo: %d, Capacidad: %d, Datos: %v\n", len(numeros), cap(numeros), numeros)

	// 5. El "Slicing" (Cortar pedazos de una lista)
	// Sintaxis: slice[inicio:fin] (el fin no se incluye)
	frameworks := []string{"Laravel", "React", "Angular", "Flutter", "Django"}
	preferidos := frameworks[0:3] // Índice 0, 1 y 2 (Laravel, React, Angular)
	
	fmt.Println("Todos los frameworks:", frameworks)
	fmt.Println("Mis preferidos (slice del slice):", preferidos)

	// 6. Copiar Slices (Seguridad de datos)
	// Si igualas sliceA = sliceB, ambos apuntan a la misma memoria.
	// Para duplicar realmente, usamos copy().
	copia := make([]string, len(frameworks))
	copy(copia, frameworks)
	fmt.Println("Copia exacta del slice:", copia)

	// 7. Iteración (Range)
	fmt.Println("\n--- Recorriendo la lista ---")
	for i, valor := range frameworks {
		fmt.Printf("Índice: %d | Framework: %s\n", i, valor)
	}
}