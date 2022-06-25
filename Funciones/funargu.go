package main
import "fmt"
type operacion func(balance,cantidad int)int

func suma(num1,num2 int)int{//operacion
    return num1+num2
}
func resta(num1,num2 int)int{//oprearcion
	return num1-num2
}
// func multiplicacion(numerol,numero2,numero3 int)int{
//         return numero1*numero2*numero3
//    } erro porqu envia 3
func ejecuope(funcion operacion,balance,cantidad int){
fmt.Println("antes de la operacion")

//resul:=funcion(10,20)
resul:=funcion(balance,cantidad) 

fmt.Println("el resultado es:",resul)
fmt.Println("despues de la opeacion")
}
func main(){
	
ejecuope(resta,100,50)
//ejecuope(multiplicacion,100,50) porque posese tre parametro

}
