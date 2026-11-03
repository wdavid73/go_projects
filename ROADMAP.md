# Ruta de aprendizaje de Go

Cada proyecto va en su propia carpeta. Un proyecto está terminado solo cuando las 8 casillas R1-R8 están marcadas. Las reglas completas están en `CLAUDE.md`.

## Nivel 1: Fundamentos (1-2 semanas)

### 01-conversor: CLI de conversión de unidades o calculadora
- Conceptos: variables, funciones, `switch`, múltiples retornos, `error` como valor.
- Entregable: CLI que convierte longitud, peso y temperatura (o calculadora) con errores claros.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 02-adivina-numero: Adivina el número o Ahorcado
- Conceptos: bucles, `math/rand`, `stdin` con `bufio`, strings y runes.
- Entregable: juego de terminal con validación, contador de intentos y fin de partida.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 03-todo-cli: Gestor de tareas en terminal
- Conceptos: `struct`, slices, maps, métodos, `encoding/json`.
- Entregable: CLI que crea, lista, completa y elimina tareas, persistidas en JSON.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

## Nivel 2: Idiomas propios de Go (2-3 semanas)

### 04-grep: Contador de palabras o grep simplificado
- Conceptos: archivos, `io.Reader`, `bufio.Scanner`, `flag` (luego `cobra` como comparación).
- Entregable: CLI que cuenta palabras, líneas y caracteres, o busca un patrón en archivos.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 05-banco-preguntas: Banco de preguntas con interfaces
- Conceptos: interfaces, implementaciones intercambiables (memoria y archivo), composición, embedding.
- Entregable: app con una interfaz `Store` y dos implementaciones sin cambiar la lógica principal.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 06-paquete-tests: Paquete reutilizable con tests
- Conceptos: `go mod`, paquetes y visibilidad, table-driven tests, benchmarks, `go test -cover`.
- Entregable: paquete propio con cobertura alta y al menos un benchmark.
- Boilerplate IA: tests repetitivos, solo después de que el usuario escribió los primeros.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

## Nivel 3: Concurrencia (2-3 semanas)

### 07-url-checker: Checker de URLs concurrente
- Conceptos: goroutines, channels, `sync.WaitGroup`, `go run -race`.
- Entregable: programa que consulta una lista de URLs en paralelo y reporta estado y latencia.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 08-worker-pool: Worker pool
- Conceptos: canales con buffer, `select`, `context` (cancelación y timeouts), cierre ordenado.
- Entregable: pool de M workers para N tareas, con cancelación y apagado limpio.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 09-rate-limiter: Rate limiter propio
- Conceptos: `time.Ticker`, `sync.Mutex`, estado compartido seguro, comparación con `golang.org/x/time/rate`.
- Entregable: limiter propio (token bucket) con tests concurrentes y nota de diferencias con la librería oficial.
- Boilerplate IA: tests repetitivos, solo después de los primeros del usuario.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

## Nivel 4: HTTP y APIs (3-4 semanas)

### 10-api-rest: API REST con net/http
- Conceptos: handlers, middleware (logging, recover), JSON, códigos de estado, mux de la stdlib (Go 1.22+).
- Entregable: API CRUD en memoria con middleware propio y tests con `httptest`.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 11-acortador-urls: Acortador de URLs
- Conceptos: `database/sql` (SQLite o Postgres) o `pgx`, redirecciones, códigos cortos, colisiones.
- Entregable: servicio que acorta y redirige, primero en memoria y luego con base de datos, sin cambiar los handlers.
- Boilerplate IA: docker-compose de la base de datos.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 12-cliente-http: Cliente HTTP robusto
- Conceptos: `http.Client` con timeouts, reintentos con backoff, `context`, errores con `%w`, `errors.Is/As`.
- Entregable: cliente para una API pública que reintenta fallos transitorios y respeta cancelación.
- Boilerplate IA: ninguno.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

## Nivel 5: Backend real (4-6 semanas)

### 13-api-postgres: API de tareas o notas con Postgres
- Conceptos: capas handler, service y repository; `golang-migrate`; `sqlc` o `pgx`; validación; config por variables de entorno.
- Entregable: API REST con CRUD en Postgres, migraciones versionadas y capas separadas.
- Boilerplate IA: docker-compose de Postgres.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 14-auth-jwt: Autenticación JWT con refresh tokens
- Conceptos: middleware de auth, `bcrypt` o `argon2`, access y refresh tokens, roles, expiración y revocación.
- Entregable: registro, login, refresh y logout sobre la API del proyecto 13, con rutas protegidas por rol.
- Boilerplate IA: ninguno. Todo el código de seguridad lo escribe y lo entiende el usuario.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 15-tests-integracion: Tests de integración
- Conceptos: `testcontainers-go`, mocks y fakes por interfaces, limpieza de datos entre tests.
- Entregable: suite que prueba repositorios y endpoints contra Postgres real, ejecutable con un comando.
- Boilerplate IA: tests repetitivos, solo después de los primeros del usuario.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

## Nivel 6: Producción (continuo)

### 16-fintrack-backend: Backend completo tipo FinTrack Pro
- Conceptos: modelado financiero, transacciones de base de datos, reportes agregados, paginación, validación de montos y monedas.
- Entregable: backend con usuarios, cuentas, transacciones y reportes para la app Flutter FinTrack Pro.
- Boilerplate IA: docker-compose y archivos de configuración.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 17-observabilidad: Observabilidad
- Conceptos: `log/slog`, métricas Prometheus, health checks, `pprof`, OpenTelemetry.
- Entregable: backend con logs estructurados, `/healthz`, `/metrics` y al menos una traza de extremo a extremo.
- Boilerplate IA: configuración de Prometheus y del collector.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 18-docker-cicd: Docker y CI/CD
- Conceptos: Dockerfile multi-stage, GitHub Actions (lint, test, build), graceful shutdown con señales y `context`.
- Entregable: imagen Docker, pipeline que corre tests y construye la imagen, apagado ordenado.
- Boilerplate IA: Dockerfile y workflow de GitHub Actions, declarados de antemano. El graceful shutdown lo escribe el usuario.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 19-extras: Extras de producción
- Conceptos: caché con Redis, colas (NATS o RabbitMQ), gRPC entre dos servicios, rate limiting, paginación por cursor.
- Entregable: cuatro mini-implementaciones, una por concepto. Cada extra pasa por el ciclo completo; no se marca una casilla hasta cumplirla en los cuatro.
- Boilerplate IA: docker-compose de Redis y de la cola.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado

### 20-final-fintech: Backend de fintech y loyalty (proyecto final)
- Conceptos: arquitectura limpia con criterio, puntos y transacciones, webhooks, idempotencia en pagos, Postgres, Redis, JWT, pruebas, Docker, despliegue.
- Entregable: servicio desplegado con README, tests de integración y pipeline de CI/CD.
- Boilerplate IA: solo lo declarado en el spec, nunca la lógica de negocio.

- [ ] R1 Me atasqué 15-20 min antes de pedir ayuda
- [ ] R2 La IA solo dio pistas
- [ ] R3 Consulté la documentación oficial primero
- [ ] R4 Explico cada línea sin mirar
- [ ] R5 Cero copiar y pegar de la IA
- [ ] R6 Code review idiomático hecho y reescrito
- [ ] R7 Evaluación de cierre aprobada
- [ ] R8 NOTES.md actualizado
