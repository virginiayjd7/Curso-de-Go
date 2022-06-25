package main
import "fmt"

func modificarvalor(parametro *string){
	*parametro="cambio de valor"
}

func main(){
	
var curso = "curso profesional de go"
fmt.Println("antes de la funsion",curso)
modificarvalor(&curso) //referecnia espacio en memoria
fmt.Println("despues de la funcion",curso)


}