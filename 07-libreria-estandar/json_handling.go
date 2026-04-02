package main

import (
	"encoding/json"
	"fmt"
)

// 1. Definimos el Struct con Tags de JSON
// Los tags `json:"nombre"` mapean la clave del JSON al campo del Struct
type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

func main() {
	fmt.Println("--- MANEJO DE JSON EN GO ---")

	// --- PARTE A: De Struct a JSON (Marshaling) ---
	p1 := Producto{
		ID:     101,
		Nombre: "Teclado Mecánico RGB",
		Precio: 1250.50,
		Stock:  15,
	}

	// json.MarshalIndent lo hace "bonito" para humanos (con espacios)
	jsonData, err := json.MarshalIndent(p1, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Println("JSON Generado:")
	fmt.Println(string(jsonData))

	// --- PARTE B: De JSON a Struct (Unmarshaling) ---
	jsonEntrada := `{"id": 102, "nombre": "Mouse Gamer", "precio": 450.00, "stock": 50}`
	
	var p2 Producto
	// Pasamos el string convertido a bytes y el PUNTERO del struct destino
	err = json.Unmarshal([]byte(jsonEntrada), &p2)
	if err != nil {
		fmt.Println("Error al leer JSON:", err)
		return
	}

	fmt.Printf("\nStruct Restaurado:\nID: %d | Producto: %s | Precio: $%.2f\n", p2.ID, p2.Nombre, p2.Precio)

	// --- PARTE C: JSON Dinámico (Sin Struct previo) ---
	// Si no conoces la estructura, usas un mapa de tipo [string]any
	jsonDesconocido := `{"web": "google.com", "ranking": 1, "activo": true}`
	var datosDinamicos map[string]any

	json.Unmarshal([]byte(jsonDesconocido), &datosDinamicos)
	
	fmt.Println("\nLectura Dinámica (Map):")
	for clave, valor := range datosDinamicos {
		fmt.Printf("Clave: %-8s | Valor: %v (%T)\n", clave, valor, valor)
	}
}