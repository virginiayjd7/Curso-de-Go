package main

import (
	"bufio" //direcionmo de memoria read
	"fmt"
	"os"
)

func main() {
	//edad := 22.20
	//fmt.Println("mi edad es %f\n ", edad)
	//fmt.Printf("mi edad es %f\n ", edad)//%f para decimales  ay v e  v s
	// var nombre string//para leer
	// fmt.Println("coloca tu nombre:")
	// fmt.Scanf("%s\n", &nombre)
	// fmt.Printf("hola %s\n", nombre)

	reader := bufio.NewReader(os.Stdin) //metodo new rader
	fmt.Println("ingresa tu nombre:")
	nombre, err := reader.ReadString('\n') //aqui retoram dos un eero y un bueno
	if err != nil {//la condicional diferente si erro es ingual a nil cuando sale bien sal nulo naamalo paso 
		fmt.Println(err)
	} else {
		fmt.Println("Hola" + nombre)
	}

}
