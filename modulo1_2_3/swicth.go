package main

import "fmt"

func main() {

	
	//10 -8/9-6/7 -5
	var callificacion int

	fmt.Print("ingresa una calificacion")
	fmt.Scanf("%d",&callificacion)
    	
   /* switch {
	case callificacion==10:
    	fmt.Println("felicitates")
    case callificacion == 8 || callificacion ==9:
		fmt.Println("aprobaste la materia")
	case callificacion==6 ||callificacion==7:
		fmt.Println("aprobaste,pero dbes esforzarte mas")
	case callificacion >=0 && callificacion<= 5:
		fmt.Println("no aprobaste")
	default:
		fmt.Println("calificacion no valida")
	}*/
	switch  callificacion {
	case 10:
    	fmt.Println("felicitates")
    case 8 ,9:
		fmt.Println("aprobaste la materia")
	case 6,7:
		fmt.Println("aprobaste,pero dbes esforzarte mas")
	case 1,2,3,4,5:
		fmt.Println("no aprobaste")
	default:
		fmt.Println("calificacion no valida")
	}
}