package main

import (
	"fmt"
	"os"
	"strings"
	"tdas/cola"
)

type Posicion struct {
	fila, col int
}

var dirs = []Posicion{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var nombres = []string{"ARRIBA", "ABAJO", "IZQUIERDA", "DERECHA"}

type AnteriorNodo struct {
	anterior Posicion
	mov      string
}

func puedePasar(posicion Posicion, grilla []string, filas, cols int) bool {
	if posicion.fila < 0 || posicion.fila >= filas {
		return false
	}
	if posicion.col < 0 || posicion.col >= cols {
		return false
	}
	if grilla[posicion.fila][posicion.col] == '#' {
		return false
	}
	return true
}

func reconstruirCamino(anterior [][]AnteriorNodo, inicio, fin Posicion) []string {
	var movs []string
	actual := fin
	for actual != inicio {
		nodo := anterior[actual.fila][actual.col]
		movs = append(movs, nodo.mov)
		actual = nodo.anterior
	}
	for i := 0; i < len(movs)/2; i++ {
		j := len(movs) - 1 - i
		movs[i], movs[j] = movs[j], movs[i]
	}
	return movs
}

func resolver(grilla []string, filas, cols int, inicio, fin Posicion) ([]string, bool) {
	anterior := make([][]AnteriorNodo, filas)
	visitado := make([][]bool, filas)
	for i := 0; i < filas; i++ {
		anterior[i] = make([]AnteriorNodo, cols)
		visitado[i] = make([]bool, cols)
	}

	visitado[inicio.fila][inicio.col] = true
	colaC := cola.CrearColaEnlazada[Posicion]()
	colaC.Encolar(inicio)

	for !colaC.EstaVacia() {
		actual := colaC.Desencolar()
		if actual == fin {
			break
		}
		for i, d := range dirs {
			nf, nc := actual.fila+d.fila, actual.col+d.col
			vecino := Posicion{nf, nc}
			if puedePasar(vecino, grilla, filas, cols) && !visitado[nf][nc] {
				visitado[nf][nc] = true
				anterior[nf][nc] = AnteriorNodo{anterior: actual, mov: nombres[i]}
				colaC.Encolar(vecino)
			}
		}
	}

	if !visitado[fin.fila][fin.col] {
		return nil, false
	}

	return reconstruirCamino(anterior, inicio, fin), true
}

func resolverLaberintos() {
	for {
		var filas, cols int
		if _, err := fmt.Fscan(os.Stdin, &filas, &cols); err != nil {
			return
		}

		grilla := make([]string, filas)
		var inicio, fin Posicion

		for i := 0; i < filas; i++ {
			fmt.Fscan(os.Stdin, &grilla[i])
			for j, c := range grilla[i] {
				if c == 'S' {
					inicio = Posicion{i, j}
				} else if c == 'E' {
					fin = Posicion{i, j}
				}
			}
		}

		movs, encontrado := resolver(grilla, filas, cols, inicio, fin)
		if encontrado {
			fmt.Println(len(movs))
			fmt.Println(strings.Join(movs, " "))
		} else {
			fmt.Println("ERROR")
		}
	}
}

func main() {
	resolverLaberintos()
}
