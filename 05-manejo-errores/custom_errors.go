package main

import (
	"fmt"
	"time"
)

// 1. Definimos nuestro propio Struct de Error
// Esto es mucho más potente que un simple string
type ErrorDeBaseDeDatos struct {
	Codigo  int
	Mensaje string
	Fecha   time.Time
}

// 2. Implementamos la interfaz 'error'
// Al añadir este método, Go reconoce automáticamente este Struct como un error
func (e *ErrorDeBaseDeDatos) Error() string {
	return fmt.Sprintf("[%v] Error %d: %s", e.Fecha.Format("15:04:05"), e.Codigo, e.Mensaje)
}

// 3. Función que simula una consulta fallida
func consultarUsuario(id int) (string, error) {
	if id != 1 {
		// Retornamos nuestra estructura personalizada
		return "", &ErrorDeBaseDeDatos{
			Codigo:  500,
			Mensaje: "No se pudo conectar con el servidor de base de datos",
			Fecha:   time.Now(),
		}
	}
	return "Usuario: Javier", nil
}

func errores() {
	fmt.Println("--- ERRORES PERSONALIZADOS EN GO ---")

	_, err := consultarUsuario(99)

	if err != nil {
		fmt.Println("Ocurrió un problema:")
		fmt.Println(err) // Imprime el resultado de nuestro método Error()

		// 4. Type Assertion (Casting de errores)
		// Si necesitamos acceder a los campos específicos (como el Codigo)
		if dbErr, ok := err.(*ErrorDeBaseDeDatos); ok {
			fmt.Printf("\n--- Detalles Técnicos ---\n")
			fmt.Printf("Código de error: %d\n", dbErr.Codigo)
			fmt.Printf("Mensaje original: %s\n", dbErr.Mensaje)
		}
	}
}