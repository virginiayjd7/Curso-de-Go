package main

import "fmt"


func main() {

    var nom string
    var edad int
    var altura float32
	
    fmt.Print("ingrese tu nombre:")
    fmt.Scanf("%s", &nom)//lee lo que el usuaro escribe por consola %s por que es string el valor

	fmt.Print("ingrese tu edad:\n")
    fmt.Scanf("%d\n", &edad)

	fmt.Print("ingrese tu atura:")
	fmt.Scanf("%f", &altura)

	fmt.Printf("hola %s con una edad %d y una altura %.2f\n",nom,edad,altura)


	
}
