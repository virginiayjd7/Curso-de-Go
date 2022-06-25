package main

import "fmt"

func main()  {
	
	numero:=100

	contador :=0
	// ciclo while mientras sea verdadero
	//<variablea> <condicion> <incremento>

	for numero>0{
		numero=numero / 10
		contador ++
	}
	fmt.Println("la cantidad de dijitos es:",contador)

}