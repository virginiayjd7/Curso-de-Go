package main
import "fmt"

func modificarvariable(parametro string){
	fmt.Println("valor acual:",parametro)
	parametro="cambio de valor"
	fmt.Println("nueo valor:",parametro)
	fmt.Printf("%p\n",&parametro)

}


func main(){
	var curso ="curso de go"

	modificarvariable(curso)

	fmt.Println(">><",curso)
	fmt.Printf("%p\n",&curso)
	



}