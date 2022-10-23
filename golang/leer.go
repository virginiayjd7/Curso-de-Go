package main

import (
	"fmt"
	"io/ioutil"
)

func mostraerro(e error) {
	if e != nil { //no es igual a nil
		{
			panic(e)
		}
	}
//nil no tiene vaor
}

func main() {

	/*datos, errordelectura := ioutil.ReadFile("txt.txt")
	mostraerro(errordelectura)
	fmt.Println(string(datos))*/
	contenido :=[]byte("estyuujjjjjjjjjjj")
	datos:=ioutil.WriteFile("txt.txt",contenido,0745)
	mostraerro(datos)
	fmt.Print("list")

}
