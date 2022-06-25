package main
import "fmt"
type operacion func(balance,cantidad int)int

func suma(num1,num2 int)int{//operacion
    return num1+num2
}
func resta(num1,num2 int)int{//oprearcion
	return num1-num2
}

func main(){
	
ejecuope(resta,100,50)


}
