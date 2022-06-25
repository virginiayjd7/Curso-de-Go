package main
import "fmt"

func suma(num1 ,num2 int)(int,string)  {
	
	return num2+num2,"un mensaje de suma"
}	

func main(){

resul,mensaje:=suma(10,20)

fmt.Println(resul,mensaje)

}