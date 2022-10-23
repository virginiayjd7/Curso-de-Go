package main

import "fmt"

type cafe struct{
	nombre string
	precio float64
	azucar bool
	leche int
}
func main()  {
	var micafe cafe=cafe{nombre: "expreso",precio: 5.22,azucar: false,leche: 0}
	fmt.Print(micafe)
	
}