package main

import "fmt"

func main()  {
	
	// else if
	//10 -8/9-6/7 -5
	var callificacion int

	fmt.Print("ingresa una calificacion")
	fmt.Scanf("%d",&callificacion)
//________________condiciones anidadas
	// if callificacion==10{
	// 	fmt.Println("felicidades")
	// }else{
	// 	if callificacion == 8 || callificacion ==9{
	// 		fmt.Println("aprobaste la materia")
	// 	}else{
	// 		if callificacion==6 ||callificacion==7{
	// 		fmt.Println("aprobaste,pero dbes esforzarte mas")
	// 	}else{
	// 		if callificacion >=0 && callificacion<= 5{
	// 			fmt.Println("no aprobaste")
	// 		}else{
	// 			fmt.Println("calificacion no valida")
	// 		}
	// 	}
	// }
//} 
    if callificacion==10{
    	fmt.Println("felicidades")
    }else if callificacion == 8 || callificacion ==9{
		fmt.Println("aprobaste la materia")
	}else if callificacion==6 ||callificacion==7 {
		fmt.Println("aprobaste,pero dbes esforzarte mas")
	}else if callificacion >=0 && callificacion<= 5{
		fmt.Println("no aprobaste")
	}else{
		fmt.Println("calificacion no valida")
	}
}