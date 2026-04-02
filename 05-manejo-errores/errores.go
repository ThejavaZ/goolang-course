package main

import (
	"errors"
	"fmt"
	"os"
)

// 1. Función que devuelve un error básico
// Por convención, el error siempre es el ÚLTIMO valor de retorno.
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// Creamos un error simple con el paquete errors
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil // nil significa "ningún error" (como null)
}

func main() {
	fmt.Println("--- MANEJO DE ERRORES EN GO ---")

	// 2. Patrón Estándar: Comprobar el error inmediatamente
	resultado, err := dividir(10, 0)

	if err != nil {
		// Aquí manejamos el error (log, retry, stop, etc.)
		fmt.Printf("⚠️ Error detectado: %v\n", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}

	// 3. Errores al trabajar con el Sistema Operativo (Archivos)
	// Muy útil para tus proyectos de inventarios y lectura de configs
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("❌ Error al abrir el archivo:", err)
		// En un programa real, aquí podrías decidir si terminar el programa
	} else {
		fmt.Println("✅ Archivo abierto con éxito")
		defer file.Close() // Buena práctica: cerrar siempre lo que abras
	}

	// 4. "Sentinel Errors" (Errores centinela)
	// Son errores predefinidos que podemos comparar
	errPersonalizado := errors.New("acceso denegado")
	
	if errPersonalizado.Error() == "acceso denegado" {
		fmt.Println("Seguridad: Se intentó un acceso no autorizado.")
	}
}