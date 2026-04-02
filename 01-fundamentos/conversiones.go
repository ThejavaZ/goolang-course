package main

import (
	"fmt"
	"strconv" // Paquete clave: String Conversion
)

func EjemploConversiones() {
	fmt.Println("--- EJERCICIOS DE CONVERSIÓN EN GO ---")

	// 1. De String a Entero (int)
	// Muy común al recibir datos de formularios o consola
	textoNumero := "100"
	numero, err := strconv.Atoi(textoNumero) // Atoi = ASCII to Integer

	if err != nil {
		fmt.Printf("Error al convertir: %v\n", err)
	} else {
		fmt.Printf("Tipo original: string | Nuevo valor: %d (Tipo: int)\n", numero + 50)
	}

	// 2. De Entero a String
	// Útil para concatenar números en mensajes
	edad := 23
	edadString := strconv.Itoa(edad) // Itoa = Integer to ASCII
	fmt.Println("Mi edad es: " + edadString)

	// 3. De String a Flotante (float64)
	// ParseFloat requiere el string y los bits de precisión (32 o 64)
	textoPrecio := "1250.45"
	precio, _ := strconv.ParseFloat(textoPrecio, 64)
	fmt.Printf("El precio con impuestos es: %.2f\n", precio * 1.16)

	// 4. Conversiones Numéricas Directas (Casting)
	// Entre tipos numéricos se usa la función del tipo destino
	var entero32 int32 = 42
	var entero64 int64 = int64(entero32) // Conversión explícita obligatoria
	
	var pi float64 = 3.14159
	var piEntero int = int(pi) // Perderá los decimales (truncado)
	
	fmt.Printf("Float a Int: %f -> %d\n", pi, piEntero)
	fmt.Printf("Castings: int32 a int64: %d | Float a Int: %d\n", entero64, piEntero)

	// 5. El "Truco" de los Booleans
	textoBool := "true"
	valorBool, _ := strconv.ParseBool(textoBool)
	if valorBool {
		fmt.Println("El valor booleano es verdadero")
	}
}