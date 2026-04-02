package main

import "fmt"

// 1. Definimos un Struct base (Como una clase padre)
type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años.\n", p.Nombre, p.Edad)
}

// 2. COMPOSICIÓN: Embeber Persona dentro de Desarrollador
type Desarrollador struct {
	Persona // <-- Campo anónimo: Esto es el "embebido"
	Especialidad string
	Lenguaje     string
}

// 3. Podemos sobrescribir métodos o añadir nuevos
func (d Desarrollador) Programar() {
	fmt.Printf("%s está programando en %s (Especialidad: %s).\n", d.Nombre, d.Lenguaje, d.Especialidad)
}

// Sobrescribimos el saludo original de Persona
func (d Desarrollador) Saludar() {
	fmt.Printf("Hola! Soy %s, desarrollador %s.\n", d.Nombre, d.Especialidad)
}

func composicion() {
	// 4. Inicialización
	// Nota: Para llenar los datos de Persona, debemos instanciarla dentro
	dev := Desarrollador{
		Persona: Persona{
			Nombre: "Javier",
			Edad:   23,
		},
		Especialidad: "Backend",
		Lenguaje:     "Go",
	}

	// 5. Acceso directo a campos embebidos
	// No necesitas hacer dev.Persona.Nombre, puedes hacer dev.Nombre directamente
	fmt.Println("Nombre del dev:", dev.Nombre)

	// 6. Llamada a métodos
	dev.Saludar()   // Llama al método de Desarrollador (Sobrescrito)
	dev.Programar() // Método propio de Desarrollador
	
	// Si quisiéramos el saludo original de Persona:
	dev.Persona.Saludar()
}