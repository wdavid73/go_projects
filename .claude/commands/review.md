---
description: Code review idiomático del proyecto actual (R6)
argument-hint: <carpeta del proyecto, opcional>
---

Haz un code review de mi código Go como un Go senior. Alcance: $ARGUMENTS (si está vacío, la rama actual y el proyecto en curso).

1. Ejecuta `go vet ./...`, `gofmt -l .` y `go test ./...` en la carpeta del proyecto y resume los resultados.
2. Lee mi código y señala, ordenado por importancia: lo que no es idiomático, manejo de errores, nombres, estructura de paquetes, concurrencia y posibles bugs.
3. Por cada observación explica el PORQUÉ y qué principio de Go lo respalda. No reescribas el código ni me des el parche; descríbeme el cambio con palabras.
4. Al final, dime qué hice bien.

Yo reescribiré lo que mejoré. No modifiques ningún archivo `.go`.
