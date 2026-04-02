package main

import (
	"fmt"
	"sync" // Paquete fundamental para sincronización
	"time"
)

func tareaPesada(id int, wg *sync.WaitGroup) {
	// 3. Importante: Avisar que la tarea terminó al salir de la función
	// 'defer' asegura que se ejecute incluso si hay un error
	defer wg.Done()

	fmt.Printf("👷 Tarea %d: Iniciando proceso...\n", id)
	time.Sleep(2 * time.Second) // Simulando carga de trabajo
	fmt.Printf("✅ Tarea %d: ¡Finalizada!\n", id)
}

func waitgroups() {
	// 1. Declarar el WaitGroup
	var wg sync.WaitGroup

	fmt.Println("--- INICIO DEL COORDINADOR ---")

	for i := 1; i <= 3; i++ {
		// 2. Incrementar el contador por cada Goroutine que lanzamos
		wg.Add(1)
		
		// IMPORTANTE: Pasar el WaitGroup por PUNTERO (&wg)
		// Si lo pasas por valor, cada función tendrá una copia y el contador original nunca bajará
		go tareaPesada(i, &wg)
	}

	// 4. Bloquear el programa aquí hasta que el contador sea 0
	fmt.Println("Esperando a que todos los trabajadores terminen...")
	wg.Wait()

	fmt.Println("--- TODO LISTO: El programa puede cerrar con seguridad ---")
}