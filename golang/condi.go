package main

import (
	"fmt"
//	"time"
)


func main() {
	//condionale
	/*puntos :=500
	if puntos < 10{
		fmt.Print("puntos incorrectos")
	}else if puntos==100{
		fmt.Print("puntos cooecto")
	}else{
		fmt.Print("tus puntos son",puntos)
	}*/
	//bucle for
//hasta que 1 sea mejor a dies le sumamos un valor a i
	/*for i:=0; i<10;i++{
		if i==1{
			fmt.Println(i,"lelelel")
		}else{
			fmt.Println(i+1,"ejjljlj")
		}

	}*/
	//-----switch
	/*tiempo:= time.Now()

	switch diahoy :=tiempo.Weekday();diahoy{
	case 0:
		fmt.Println("descanso domin")
	case 1:
		fmt.Println("es lunes")
	default:
		fmt.Println("no hay")
	}*/
	/*var y int = 9
var z int = 0

for x := 0; x <= 10; x++ {
    z = y * x
    fmt.Println(y, " x ", x, " = ", z)
}*/
var a = 2
var b = 3
var c = 1
var opcion int

opcion = c / b
opcion = opcion + a

switch opcion {
case 1:
    fmt.Println("Opción 1")
case 2:
    fmt.Println("Opción 2")
case 3:
    fmt.Println("Opción 3")
default:
    fmt.Println("Ninguna opción")
}

}

