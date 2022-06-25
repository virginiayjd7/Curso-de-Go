package main

import "fmt"

func main()  {

	// for <condicion>
	// for <variable> ;<codificacion>; <expresion>
	// for <varieble>;<condicion>;<incremento>
	// for <variable>;<indice> := range <arreglo>
	// for{
	// 	fmt.Println("hola mundo")
	// }	
	
		var numero=1
		for{
			fmt.Println(numero)
			numero++

			if numero ==100 {
				//break
				panic("fin del cilo for")
				
			}
		}
	}
	

