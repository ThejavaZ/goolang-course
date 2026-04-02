package main

import (
	"fmt"
	"time"
)

// 1. Función que envía datos a un canal
func enviarMensaje(canal chan string) {
	fmt.Println("🚀 Empezando a procesar el mensaje...")
	time.Sleep(2 * time.Second)
	
	// El operador <- a la derecha del canal SIGNIFICA ENVIAR
	canal <- "¡Hola desde la Goroutine! Procesamiento completado."
}

func channels() {
	fmt.Println("--- INICIO DE COMUNICACIÓN ---")

	// 2. Crear un canal de tipo string
	// Los canales deben tener un tipo de dato específico
	mensajeria := make(chan string)

	// 3. Lanzamos la función en segundo plano
	go enviarMensaje(mensajeria)

	// 4. RECIBIR el mensaje
	// El operador <- a la izquierda de la variable SIGNIFICA RECIBIR
	// Importante: Esta línea BLOQUEA el programa principal hasta que llegue el dato
	fmt.Println("Esperando respuesta del canal...")
	respuesta := <-mensajeria

	fmt.Println("Mensaje recibido:", respuesta)
	fmt.Println("--- FIN DEL PROGRAMA ---")
}