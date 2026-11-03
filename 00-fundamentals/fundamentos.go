// =============================================================================
// MIS NOTAS DE "UN TOUR POR GO"
// =============================================================================
// Archivo vivo: aquí voy guardando lo que aprendo en cada capítulo del tour.
//
// PROGRESO
//   [x] Conceptos básicos
//       [x] Paquetes, variables y funciones
//           - Paquetes, imports y nombres exportados
//           - Funciones: parámetros, múltiples retornos, retornos con nombre
//           - Variables, valores cero, tipos básicos, conversiones
//           - Constantes (incluidas las numéricas "sin tipo")
//       [x] Sentencias de control de flujo
//           - for (clásico, tipo "while", infinito)
//           - if (con declaración corta) / else
//           - switch (con valor, sin condición, orden de evaluación)
//           - defer
//       [ ] Más tipos: punteros, structs, arrays, slices, maps, closures
//   [ ] Métodos e interfaces
//   [ ] Genéricos
//   [ ] Concurrencia (goroutines, channels, select, mutex)
//
// Para ejecutarlo:  go run main.go
// =============================================================================

// Todo programa ejecutable de Go vive en el paquete `main`
// y arranca en la función `main()`.
package main

// Imports agrupados ("factored import"). Es la forma idiomática.
// Go NO compila si importas un paquete que no usas.
import (
	"fmt"        // Formateo e impresión: Print, Println, Printf, Sprint...
	"math"       // Funciones matemáticas: Sqrt, Pow, Pi...
	"math/cmplx" // Matemáticas para números complejos
	"math/rand"  // Números pseudoaleatorios
	"runtime"    // Información del entorno de ejecución (SO, arquitectura...)
	"time"       // Fechas, horas y duraciones
)

// -----------------------------------------------------------------------------
// FUNCIONES
// -----------------------------------------------------------------------------

// sum recibe dos enteros y devuelve su suma.
// Cuando varios parámetros consecutivos comparten tipo, puedes escribir
// el tipo una sola vez:  (a, b int)  equivale a  (a int, b int).
// El tipo de retorno va DESPUÉS de los parámetros.
func sum(a, b int) int {
	return a + b
}

// swap demuestra MÚLTIPLES VALORES DE RETORNO.
// Se declaran entre paréntesis: (string, string).
// Quien llama los recibe así:  a, b := swap("x", "y")
// Este patrón es la base del manejo de errores en Go:  valor, err := algo()
func swap(x, y string) (string, string) {
	return y, x
}

// split demuestra RETORNOS CON NOMBRE ("named results").
// x e y se declaran en la firma como variables, inicializadas en su valor cero (0).
// Un `return` sin argumentos ("naked return") devuelve sus valores actuales.
// ⚠️ Úsalo solo en funciones cortas: en funciones largas empeora la legibilidad.
func split(sum int) (x, y int) {
	x = sum * 4 / 9 // División entera: 400 / 9 = 44 (se trunca)
	y = sum - x
	return // Devuelve x e y
}

// needInt y needFloat se usan más abajo con constantes numéricas
// para mostrar que una constante "sin tipo" se adapta al tipo que se necesite.
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// sqrt muestra un IF simple y RECURSIVIDAD.
// - En Go el `if` NO lleva paréntesis, pero las llaves {} son obligatorias.
// - Si x es negativo, calcula la raíz del positivo y le agrega "i" (imaginario).
// - fmt.Sprint convierte cualquier valor a string (no imprime, devuelve el texto).
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// pow muestra el IF CON DECLARACIÓN CORTA.
// Sintaxis:  if inicialización; condición { ... }
// La variable `v` solo existe dentro del if y de sus bloques else / else if.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// `v` también es visible aquí, dentro del else.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// Aquí `v` YA NO existe: su alcance (scope) terminó con el if/else.
	return lim
}

// -----------------------------------------------------------------------------
// FUNCIÓN PRINCIPAL
// -----------------------------------------------------------------------------

func main() {
	// ===== IMPRESIÓN =====
	// fmt.Print    → imprime tal cual; agrega espacio solo entre operandos que NO son strings.
	// fmt.Println  → agrega espacios entre todos los operandos y un salto de línea al final.
	// fmt.Printf   → usa verbos de formato (%v, %T, %g, %s, %d...). No agrega salto de línea.
	fmt.Print("Hello world!", "\n")

	fmt.Print("La hora es:", time.Now(), "\n")

	// rand.Intn(n) devuelve un entero en [0, n).
	// Desde Go 1.20 la semilla es aleatoria automáticamente (antes siempre daba lo mismo).
	fmt.Print("Mi numero favorito es:", rand.Intn(100), "\n")

	// %g → formato compacto para flotantes.
	fmt.Printf("Ahora tienes %g problema. \n", math.Sqrt(7))

	// Nombres EXPORTADOS: en Go algo es público si empieza con MAYÚSCULA.
	// math.Pi funciona; math.pi NO compilaría.
	fmt.Println(math.Pi)

	fmt.Println(sum(23, 56))

	// Recibir múltiples retornos con declaración corta `:=`.
	a, b := swap("hello", "world")
	fmt.Println(a, b) // world hello

	// Println acepta directamente los múltiples retornos de una función.
	fmt.Println(split(100)) // 44 56

	// ===== VARIABLES =====
	// `var` declara variables. Sin valor inicial, reciben su VALOR CERO:
	//   bool → false | números → 0 | string → "" | punteros, slices, maps... → nil
	var c bool
	var python bool
	var java string

	fmt.Println(c, python, java) // false false ""

	// Asignación múltiple (se usa `=` porque ya están declaradas).
	c, python, java = true, false, "no!"

	fmt.Println(c, python, java)

	// `var` con inicializador: el tipo puede omitirse y se infiere.
	var i, j int = 1, 2
	// Declaración corta `:=` → declara e infiere el tipo.
	// ⚠️ Solo funciona DENTRO de funciones. A nivel de paquete hay que usar `var`.
	k := 3
	fmt.Println(i, j, k)

	/*
		===== TIPOS BÁSICOS =====

		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias para uint8

		rune // alias para int32
		     // representa un punto de código Unicode (un "carácter")

		float32 float64

		complex64 complex128

		Consejo: usa `int` para enteros salvo que tengas una razón concreta
		(tamaño fijo, sin signo). int/uint/uintptr miden 32 bits en sistemas
		de 32 bits y 64 bits en sistemas de 64 bits.
	*/

	// Bloque `var ( ... )`: agrupa varias declaraciones, igual que los imports.
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1            // Máximo valor de uint64 (18446744073709551615)
		z      complex128 = cmplx.Sqrt(-5 + 12i) // Raíz de un número complejo → (2+3i)
	)

	// %T → imprime el TIPO del valor.
	// %v → imprime el VALOR en su formato por defecto.
	fmt.Printf("Tipo: %T Valor: %v\n", ToBe, ToBe)
	fmt.Printf("Tipo: %T Valor: %v\n", MaxInt, MaxInt)
	fmt.Printf("Tipo: %T Valor: %v\n", z, z)

	// INFERENCIA DE TIPOS con literales:
	//   42      → int
	//   3.14    → float64
	//   0.5i    → complex128
	g := 0.867 + 0.5i // complex128

	fmt.Println("Complex128", g)

	// CONVERSIONES (no hay en este archivo, pero es clave recordarlo):
	// Go NO convierte tipos implícitamente. Siempre es explícito: T(v)
	//   var n int = 42
	//   var f float64 = float64(n)
	//   var u uint = uint(f)

	// ===== CONSTANTES =====
	// Se declaran con `const`. Pueden ser string, bool, numéricas o caracteres.
	// ⚠️ NO se pueden declarar con `:=`.
	const work = "¡Trabajando!"
	fmt.Println(work)

	// CONSTANTES NUMÉRICAS "SIN TIPO":
	// Son valores de alta precisión. No tienen tipo fijo hasta que se usan,
	// y entonces toman el tipo que exija el contexto.
	const (
		// Crear un número enorme desplazando 1 bit a la izquierda 100 lugares.
		// Es decir, el número binario que es 1 seguido de 100 ceros.
		Big = 1 << 100
		// Moverlo a la derecha de nuevo 99 lugares, por lo que terminamos con:
		// 1<<1 o 2.
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))   // Small (2) cabe en int → 21
	fmt.Println(needFloat(Small)) // Small se usa como float64 → 0.2
	fmt.Println(needFloat(Big))   // Big cabe en float64 → 1.2676506002282295e+29
	// fmt.Println(needInt(Big))  // ❌ NO compila: Big desborda int (overflow)

	// ===== BUCLES =====
	// Go tiene UNA sola palabra para bucles: `for`. Sirve para todo.

	// 1) for clásico: inicialización; condición; post
	//    Sin paréntesis, pero llaves obligatorias.
	//
	// ⚠️ OJO: aquí `sum` es una VARIABLE que "sombrea" (shadowing) a la función sum().
	//    A partir de esta línea, dentro de main, ya no podrías llamar a sum(1, 2).
	//    Funciona, pero en código real conviene evitar reutilizar nombres así.
	sum := 0
	for i := 0; i < 10; i++ { // Este `i` es nuevo y solo vive dentro del for
		sum += i
	}
	fmt.Println("Suma:", sum) // 45

	// 2) for como "while": solo con condición.
	//    Go no tiene la palabra `while`.
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2 // Se duplica: 1, 2, 4, 8... hasta 1024
	}
	fmt.Println("Suma 2:", sum2)

	// 3) Bucle infinito: `for` sin nada. Se sale con break o return.
	//for {
	//}

	// 4) (Más adelante en el tour) for con range, para recorrer slices, maps, strings:
	//    for indice, valor := range coleccion { ... }

	// ===== CONDICIONALES =====
	fmt.Println(sqrt(2), sqrt(-4)) // 1.4142135623730951 2i

	// pow(3, 2, 10) → 9 < 10, devuelve 9
	// pow(3, 3, 20) → 27 >= 20, imprime "27 >= 20" y devuelve 20
	// Nota: el Printf de dentro de pow se ejecuta ANTES de que este Println imprima,
	// porque Go evalúa todos los argumentos antes de llamar a la función.
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// ===== SWITCH =====
	// Diferencias clave con otros lenguajes (C, Java, JS, Dart):
	//   - Solo se ejecuta el caso que coincide: NO hace falta `break`.
	//   - Si quieres que siga al siguiente caso, usa `fallthrough` explícitamente.
	//   - Los casos no tienen que ser constantes ni enteros.

	// 1) switch con declaración corta (igual que el if).
	//    `os` solo existe dentro del switch.
	fmt.Print("Go se ejecuta en ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// 2) Orden de evaluación: los casos se evalúan de ARRIBA hacia ABAJO
	//    y se detiene en el primero que coincide (los siguientes ni se evalúan).
	//    Aquí se compara time.Saturday contra cada expresión.
	fmt.Println("¿Cuándo es sábado?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Hoy.")
	case today + 1:
		fmt.Println("Mañana.")
	case today + 2:
		fmt.Println("En dos días.")
	default:
		fmt.Println("Falta mucho.")
	}

	// 3) switch SIN condición = `switch true`.
	//    Es la forma limpia de escribir cadenas largas de if / else if / else.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buenos días!")
	case t.Hour() < 17:
		fmt.Println("Buenas tardes.")
	default:
		fmt.Println("Buenas noches.")
	}

	// ===== DEFER =====
	// `defer` pospone la ejecución de una llamada hasta que la función que la
	// contiene (aquí, main) termine.
	//   - Los ARGUMENTOS se evalúan en el momento del defer, no al final.
	//   - Varios defer se apilan (LIFO): el último declarado se ejecuta primero.
	//   - Uso típico: liberar recursos →  f, _ := os.Open(...); defer f.Close()
	defer fmt.Println("mundo")

	fmt.Println("hola")
	// Salida: "hola" y luego, al terminar main, "mundo".
}
