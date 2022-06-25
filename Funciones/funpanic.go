package main

import "fmt"

func main()  {
	//panic
	var dividendo, divisor int
	fmt.Print("ingresa un  valor para el dividendo:")
	fmt.Scanf("%d", &dividendo)

	fmt.Print("ingresa un  valor para el divisor:")
	fmt.Scanf("%d", &divisor)
	if divisor ==0{
		panic("no es posible dividir")
	}

	resul := dividendo / divisor

	fmt.Println(resul)

}