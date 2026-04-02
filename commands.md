# 🛠 Guía Rápida de Comandos Go (Golang)

Esta es una referencia rápida de los comandos de la CLI de Go que utilizaremos a lo largo de este curso en Fedora.

---

## 📦 Gestión de Módulos y Dependencias

Go utiliza módulos para gestionar el código y las librerías externas.

- **`go mod init <nombre-del-modulo>`**: Inicializa un nuevo proyecto. Crea el archivo `go.mod`.
  - _Ejemplo:_ `go mod init github.com/tu-usuario/curso-go`
- **`go mod tidy`**: Sincroniza las dependencias. Elimina las que no se usan y descarga las faltantes basándose en tus `imports`.
- **`go get <url-paquete>`**: Descarga una librería externa y la añade al proyecto.
  - _Ejemplo:_ `go get github.com/gin-gonic/gin`

---

## 🚀 Ejecución y Compilación

Go permite un flujo de trabajo rápido entre escribir y probar código.

- **`go run <archivo.go>`**: Compila y ejecuta el programa en un solo paso (ideal para desarrollo).
- **`go build`**: Compila el paquete actual y genera un binario ejecutable en la carpeta actual.
- **`go build -o <nombre-salida>`**: Compila y le asigna un nombre específico al ejecutable.
- **`go install`**: Compila e instala el binario en `$GOPATH/bin` para que sea ejecutable desde cualquier parte de tu terminal.

---

## 🧹 Calidad y Formato

Go es muy estricto con el estilo de código para mantener la legibilidad.

- **`go fmt ./...`**: Formatea automáticamente todos los archivos del proyecto siguiendo el estándar oficial.
- **`go vet ./...`**: Examina el código en busca de errores lógicos que el compilador no detecta (como variables sin usar o formatos incorrectos).
- **`go doc <paquete>.<funcion>`**: Muestra la documentación de una función o paquete directamente en la terminal.
  - _Ejemplo:_ `go doc fmt.Println`

---

## 🧪 Testing y Benchmarking

Go tiene herramientas de prueba integradas de primer nivel.

- **`go test ./...`**: Ejecuta todos los tests en el proyecto (archivos que terminan en `_test.go`).
- **`go test -v`**: Ejecuta los tests en modo "verbose" (detallado).
- **`go test -bench=.`**: Ejecuta las pruebas de rendimiento (benchmarks).

---

## ⚙️ Entorno y Utilidades

- **`go env`**: Imprime todas las variables de entorno de Go instaladas en tu sistema Fedora.
- **`go version`**: Muestra la versión actual de Go instalada.
- **`go clean`**: Elimina archivos temporales y binarios generados por compilaciones anteriores.

---

## 🐧 Tip para Fedora (Alias)

Si quieres ser aún más rápido, puedes añadir alias a tu `.bashrc` o `.zshrc`:

```bash
alias gr="go run"
alias gb="go build"
alias fmt="go fmt ./..."
```
