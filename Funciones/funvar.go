package main
import "fmt"

func factorial(numero int)int{
    if numero ==1{
        return 1

    }
    return factorial(numero-1)*numero
}

type customefunct func(n int)int

func main(){

    //var miFuncion=factorial
	var miFuncion customefunct
    if miFuncion ==nil{
		miFuncion=factorial
	}
	//miFuncion=factorial
	resultado:= miFuncion(3)
	fmt.Println(resultado)
}