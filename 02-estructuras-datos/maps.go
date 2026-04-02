package main

import "fmt"

func maps() {
	// 1. Declaración e Inicialización
	// make(map[tipo_clave]tipo_valor)
	usuarios := make(map[string]int)

	// Agregar elementos
	usuarios["Javier"] = 23
	usuarios["Paola"] = 22
	usuarios["Mark"] = 1 // El perrito también cuenta

	fmt.Println("Mapa de usuarios:", usuarios)

	// 2. Inicialización directa (Literal)
	precios := map[string]float64{
		"Laptop":    15000.50,
		"Mouse":     350.00,
		"Teclado":   800.00,
	}

	fmt.Println("Inventario de precios:", precios)

	// 3. Acceder a un valor y verificar si existe
	// En Go, si accedes a una clave que no existe, te da el valor por defecto (0, "", etc.)
	// Para saber si realmente existe, usamos el "Comma ok idiom"
	nombre := "Pedro"
	edad, existe := usuarios[nombre]

	if existe {
		fmt.Printf("%s tiene %d años.\n", nombre, edad)
	} else {
		fmt.Printf("El usuario %s no está en la base de datos.\n", nombre)
	}

	// 4. Eliminar elementos
	delete(usuarios, "Mark")
	fmt.Println("Después de eliminar a Mark:", usuarios)

	// 5. Iterar un Mapa (Range)
	// Nota: El orden de los mapas en Go es ALEATORIO cada vez que ejecutas
	fmt.Println("\n--- Lista de Precios ---")
	for producto, precio := range precios {
		fmt.Printf("Producto: %-10s | Precio: $%.2f\n", producto, precio)
	}

	// 6. Tamaño del mapa
	fmt.Println("\nTotal de productos en inventario:", len(precios))
}