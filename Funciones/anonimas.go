package main
import "fmt"

func main(){
miFuncion :=func (username string) string {
	//fmt.Println("hola",username,"desde una funcion in nombre")
	message:=fmt.Sprintf("hola %s, te allll",username)
	return message
}	
mensaje:=miFuncion("cody 1")
fmt.Println(mensaje)




}