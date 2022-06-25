package main

import "fmt"

func main(){
	// // matriz := []int{1,2,3,4,5}
	// // //matriz [0][1]=1 
	// // fmt.Println(matriz)

	// matriz := []int{1,2} /// var no inicialo el valor es nulo 0 para cada dato un 

	// if matriz == nil{
	// 	fmt.Println("esta vacio")
	// }else{
	// 	fmt.Println(len(matriz))//len //arreglos 
	// 	//puntero al arreglo
	// 	//longitud del areglo al que apunta
	// 	//capacidad
	// }
	
	arreglo := [3]int{1,2,3}//esto es un arreglo
	slice :=arreglo[0:3]//aslide partir un arreglo
	fmt.Println(slice)
}