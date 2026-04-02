package main

import "fmt"

// 1. Definición del "Objeto" (Struct)
// Imagina que esto es una clase en C# o Laravel
type Usuario struct {
	Nombre string
	Edad   int
	Email  string
}

// 2. MÉTODO CON RECEPTOR DE VALOR (Copia)
// (u Usuario) es el "receptor". Permite usar u.Presentarse()
// Funciona como el 'this' o 'self' pero es una copia de los datos.
func (u Usuario) Presentarse() {
	fmt.Printf("Hola, mi nombre es %s y tengo %d años.\n", u.Nombre, u.Edad)
}

// 3. MÉTODO CON RECEPTOR DE PUNTERO (Referencia)
// Usamos *Usuario para poder MODIFICAR los datos originales.
// Si no usamos *, Go crea una copia y el cambio se pierde al salir.
func (u *Usuario) CumplirAnios() {
	u.Edad++
	fmt.Printf("¡Felicidades %s! Ahora tienes %d años.\n", u.Nombre, u.Edad)
}

// 4. Los métodos no solo son para Structs. 
// Puedes crear tipos nuevos y darles métodos.
type Dinero float64

func (d Dinero) Formatear() string {
	return fmt.Sprintf("$%.2f MXN", d)
}

func main() {
	// Instanciar el struct
	persona := Usuario{
		Nombre: "Javier",
		Edad:   23,
		Email:  "javier@example.com",
	}

	// Llamar a un método de lectura (Valor)
	persona.Presentarse()

	// Llamar a un método de escritura (Puntero)
	// Go es tan amable que no necesitas poner (&persona).CumplirAnios()
	// lo hace por ti automáticamente.
	persona.CumplirAnios()

	// Probando el método del tipo personalizado
	miSaldo := Dinero(1500.50)
	fmt.Println("Saldo actual:", miSaldo.Formatear())
}