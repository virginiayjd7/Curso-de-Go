package main
import "fmt"

type operacion func (balance, cantidad int) int 

func crearOperacion(tipo string) operacion  {

	if tipo =="suma"{
		return func (balance,cantidad int) int {return balance + cantidad}  
	}else if tipo =="resta"{
		return func (balance , cantidad int) int  {return balance - cantidad}
	}else{
		return func (balance, cantidad int)int  {return balance * cantidad}
	}
	}
func main(){
	

suma:=crearOperacion("suma")

resul:= suma(40,50)
fmt.Println("el resultado es:",resul)
}