package main

import "fmt"

func main(){
	// if nombre, edad :="Cody" , 7;nombre=="cody"{
	// 	fmt.Println("hola",nombre,"te damos la bienvenida al curso de go")
	// }else{
	// 	fmt.Println("los valores son",nombre,edad)
	// }

	animales :=[...]string{"Perro","Gato","Pez","Vaca","Cerdo"}

	/*for indice := 0; indice < len(animales);indice++{
		elemento := animales[indice]
		fmt.Println(elemento)
	}*/
	for _,indice ,elemento:=range animales{
		fmt.Println("el indice es:",indice,"el valro:",elemento)
	}


}