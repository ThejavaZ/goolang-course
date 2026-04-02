package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Declaración de Variables
	// Forma larga (útil para variables globales o sin valor inicial)
	var mensaje string = "Hola desde la consola de Go"

	// Forma corta (Solo dentro de funciones, infiere el tipo)
	// Es el estándar en Go para variables locales
	version := 1.21 

	fmt.Println(mensaje)
	fmt.Printf("Estamos usando la versión de Go: %.2f\n", version)
	fmt.Println(strings.Repeat("-", 20))

	// 2. Entrada de datos con bufio (Permite leer frases completas)
	// os.Stdin es el estándar de entrada (teclado)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingresa tu nombre completo: ")
	// ReadString lee hasta encontrar el carácter que le indiquemos
	nombre, _ := reader.ReadString('\n')

	// 3. Limpieza de strings
	// Al leer de consola, el "Enter" se queda pegado al string (\n)
	// strings.TrimSpace elimina espacios y saltos de línea al inicio y final
	nombre = strings.TrimSpace(nombre)

	// 4. Tipos Booleano y Condicional Simple
	esDesarrollador := true

	fmt.Printf("Bienvenido al curso, %s.\n", nombre)

	if esDesarrollador {
		fmt.Println("Veo que ya eres desarrollador. ¡Go te va a encantar por su simplicidad!")
	} else {
		fmt.Println("¡Excelente momento para empezar tu camino en la programación con Go!")
	}

	// 5. Constantes
	const Ciudad = "Hermosillo"
	fmt.Printf("Este curso fue creado en: %s, Sonora.\n", Ciudad)
}