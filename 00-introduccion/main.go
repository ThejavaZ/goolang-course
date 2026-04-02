package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 1. Declaración de variables
	// En Go, usamos := para declarar e inicializar rápido (inferencia de tipos)
	// Es el equivalente a 'var' en C#
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- BIENVENIDO AL CURSO DE GO ---")

	// 2. Entrada de datos (Strings)
	fmt.Print("Ingresa tu nombre: ")
	nombre, _ := reader.ReadString('\n') // Leemos hasta el salto de línea
	
	// Limpieza de datos: strings.TrimSpace es como .Trim() en C# o trim() en PHP
	nombre = strings.TrimSpace(nombre)

	// 3. Entrada de datos (Números)
	fmt.Print("Ingresa tu año de nacimiento: ")
	anioInput, _ := reader.ReadString('\n')
	anioInput = strings.TrimSpace(anioInput)

	// Conversión de tipos: strconv.Atoi (ASCII to Integer)
	// Go no tiene excepciones (try/catch), así que las funciones devuelven el error
	anioNacimiento, err := strconv.Atoi(anioInput)

	if err != nil {
		fmt.Println("Error: El año ingresado no es válido.")
		return // Salimos de la función si hay error
	}

	// 4. Lógica básica
	anioActual := time.Now().Year()
	edad := anioActual - anioNacimiento

	// 5. Salida con formato (Printf)
	// %s = string, %d = int (decimal)
	fmt.Printf("\nHola %s, estamos en el año %d.\n", nombre, anioActual)
	fmt.Printf("Basado en tu año de nacimiento (%d), tienes %d años.\n", anioNacimiento, edad)

	// 6. Comparación simple (If/Else)
	if edad >= 18 {
		fmt.Println("Eres mayor de edad en México. ¡Ya puedes votar!")
	} else {
		fmt.Println("Eres menor de edad.")
	}
}