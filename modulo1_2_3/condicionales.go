package main

import "fmt"

func main()  {
	var edad int
	fmt.Print("ingrese la edad:")
	fmt.Scanf("%d",&edad)

	if edad >=18{
		fmt.Println("el usurio es mayor de edad")
	}else{
		fmt.Println("el usuario es menor de edad!")
	}


	// if false{
	// 	fmt.Println("la condicion se cumple")
	// } else{
	// 	fmt.Println("la condicon se cumple")
	// }
	

}