# 00 - Introducción a Go (Golang)

Bienvenido al punto de partida. En este módulo configuraremos el entorno y entenderemos qué hace a Go un lenguaje único para el desarrollo moderno.

## 🚀 ¿Qué es Go?

Go es un lenguaje de programación de **código abierto** creado por Google. Fue diseñado para ser simple, eficiente y confiable.

### Características principales:

- **Compilado:** Se traduce directamente a código máquina (binario), lo que lo hace muy rápido.
- **Tipado Estático:** Los errores de tipo se detectan en tiempo de compilación.
- **Concurrente:** Maneja miles de tareas simultáneas de forma nativa con _Goroutines_.
- **Garbage Collected:** Gestión de memoria automática (adiós a los problemas manuales de C).

---

## 💻 Configuración en Fedora Linux

Como usuario de Fedora, instalar Go es sumamente sencillo:

```bash
# Instalación del SDK
sudo dnf install golang

# Verificar la instalación
go version
```
