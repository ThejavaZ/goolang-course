package main

import (
	"fmt"
	"strings"
	"time"
)

// 1. Una función normal que simula un proceso lento
func mostrarProceso(nombre string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("[%s] Paso %d...\n", strings.ToUpper(nombre), i)
		// Simulamos una tarea que tarda (como una petición a una API)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("--- INICIO DEL PROGRAMA ---")

	// 2. EJECUCIÓN SÍNCRONA (Secuencial)
	// El programa se detiene aquí hasta que termine 'Tarea A'
	// mostrarProceso("Tarea A")

	// 3. EJECUCIÓN ASÍNCRONA (Goroutines)
	// Al agregar la palabra 'go', la función se ejecuta en segundo plano
	go mostrarProceso("Tarea A")
	go mostrarProceso("Tarea B")

	// 4. El problema del Main Thread
	// Si el programa llega al final de main(), se cierra TODO, 
	// incluso si las goroutines no han terminado.
	fmt.Println("Las tareas están corriendo en segundo plano...")
	
	// Por ahora usamos un Sleep para esperar (forma "sucia")
	// En el siguiente ejercicio veremos WaitGroups (la forma correcta)
	time.Sleep(3 * time.Second)

	fmt.Println("--- FIN DEL PROGRAMA ---")
}