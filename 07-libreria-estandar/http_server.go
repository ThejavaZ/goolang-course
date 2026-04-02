package main

import (
	"fmt"
	"net/http"
)

// 1. Definimos un controlador (Handler)
// w: donde escribimos la respuesta (ResponseWriter)
// r: donde leemos lo que el cliente envió (Request)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>¡Bienvenido a mi servidor en Go!</h1><p>Hola Javier, esto corre en Fedora.</p>")
}

func contactoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Puedes contactarme en mi perfil de GitHub: @javiersg")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado para que el navegador sepa que es JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "success", "message": "API de inventario activa"}`)
}

func https() {
	// 2. Mapear las rutas (Ruteo básico)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contacto", contactoHandler)
	http.HandleFunc("/api/v1", apiHandler)

	// 3. Iniciar el servidor
	puerto := ":8080"
	fmt.Println("🚀 Servidor iniciado en http://localhost" + puerto)
	
	// ListenAndServe bloquea el programa y mantiene el servidor escuchando
	err := http.ListenAndServe(puerto, nil)
	
	if err != nil {
		fmt.Printf("No se pudo iniciar el servidor: %v\n", err)
	}
}