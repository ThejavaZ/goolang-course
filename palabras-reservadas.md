# Declaración de Tipos y Variables

Estas se usan para definir la estructura de tu programa y los datos que maneja.

- **package**: Define a qué paquete pertenece el archivo.

- **import**: Incluye librerías o paquetes externos.

- type: Crea nuevos tipos (como structs o interfaces).

- struct: Define una colección de campos (lo más parecido a una clase).

- interface: Define un conjunto de firmas de métodos.

- var: Declara una variable.

- const: Declara una constante.

- func: Declara una función o método.

1. Control de Flujo (Lógica)

Palabras para tomar decisiones y repetir acciones.

    if: Condicional básico.

    else: Bloque alternativo al if.

    switch: Estructura de decisión múltiple.

    case: Una opción dentro de un switch.

    default: La opción por defecto en un switch.

    for: El único bucle en Go (sirve para todo).

    range: Itera sobre elementos (arrays, slices, maps).

    break: Sale de un bucle o switch.

    continue: Salta a la siguiente iteración de un bucle.

    return: Devuelve un valor desde una función.

    fallthrough: Se usa en switch para pasar a la siguiente caja (poco común).

3. Concurrencia y Manejo Especial

Aquí es donde Go brilla con sus herramientas nativas para procesos paralelos.

    go: Lanza una goroutine (ejecuta una función en segundo plano).

    chan: Define un canal para comunicar datos entre goroutines.

    select: Como un switch pero para operaciones con canales.

    defer: Retrasa la ejecución de una función hasta que la función actual termine (ideal para cerrar archivos o conexiones).

4. Otros

   map: Define una estructura de datos clave-valor.

   goto: Salto incondicional a una etiqueta (casi nunca se recomienda usarlo).

Un detalle curioso:

Notarás que faltan palabras que en otros lenguajes son comunes:

    No existe class ni extends: Go prefiere la composición sobre la herencia.

    No existe public ni private: Como vimos, esto se decide por si la primera letra es mayúscula o minúscula.

    No existe try, catch o throw: El manejo de errores en Go se hace retornando el error como un valor normal.
