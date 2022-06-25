package main

import "fmt"

func main ()  {
	// // // arreglo := [3]int{1,2}
	// // // //fmt.Println(len (arreglo))//len cuantos elemntos cabe en un arreglo
	// // // arreglo[2]=20
	// // // for i:= 0; i < len(arreglo);i++{
	// // // 	fmt.Println(arreglo[i])

	// // // }

	// // var matriz [3] [2]int//tre arreglos con 2 elementos
    // // matriz[0][1]= 2//en el arreglo 0  se asigna el na posin 1 en valor dos
	// // fmt.Println(matriz)
	
	// // var numero[5]int
	// //numero :=[5] int {100,200,300,400,500}
	// numero :=[...] int {100,200,300,400,500}
	// // numero[0]=100
	// // numero[1]=200
	// // numero[2]=300
	// // numero[3]=400
	// // numero[4]=500

	
	// fmt.Println(numero)
	// //fmt.Println(numero[6])


	money:=[...]string {0:"dolar cadab",3:"peso",1:"dolar",5:"euro"}
	// fmt.Println(money[0])
	// fmt.Println(money[1])
	// fmt.Println(money[2])
	// fmt.Println(money[5])

	submoney := money[0:3]//slice

	fmt.Println(submoney)
}