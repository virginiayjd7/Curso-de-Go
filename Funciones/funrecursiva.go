package main
import "fmt"

func factorial(numero int)int{
	if numero ==1{
		return 1
	}
	return factorial(numero - 1)*numero
}


func main(){
	resul :=factorial(3)
	fmt.Println("el factorial es :",resul)


}