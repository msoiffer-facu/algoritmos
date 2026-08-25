package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp0/ejercicios"
)
func leerArchivo(ruta string) []int {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	var numeros []int
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea != "" {
			num, _ := strconv.Atoi(linea)
			numeros = append(numeros, num)
		}
	}
	return numeros
}

func main(){
	arr1 := leerArchivo("archivo1.in")
	arr2 := leerArchivo("archivo2.in")

	ejercicios.Seleccion(arr1)
	ejercicios.Seleccion(arr2)

	// Comparar los arreglos ordenados
	resultado := ejercicios.Comparar(arr1, arr2)
	if resultado == -1 {
		for i := 0; i < len(arr2); i++ {
			fmt.Println(arr2[i])
		}
	} else if resultado == 1 {
		for i := 0; i < len(arr1); i++ {
			fmt.Println(arr1[i])
		}
	} else {
		fmt.Println("Array 1 es igual a Array 2")
	}
}