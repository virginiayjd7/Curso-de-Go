package main
import "fmt"

func mostrarvariable(objeto interface{}){
	fmt.Println("el valor de la variable es:%v\n",objeto)
} 
func suma(num1, num2 int)int{
	return num1+num2
}
type user struct{
	nom string
}
func main(){
	
usu:=user{nom: "yaneth"}
mostrarvariable(usu)


}