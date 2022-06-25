package main

import (
	"fmt"
	"reflect"
)

func main()  {
	
	//var curso string=" curso"
//var curso ="ssuppp"
    curso := "hhhhho"//lista de aracter comine de 0
	lon:=len(curso)//int longitu
	fmt.Println(lon)

	p := curso[0]//char es variable
    ultimo:=curso[len(curso)-1]
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))//para comprobar el char

	fmt.Printf("%c \n",p)//acceder  en caracter
	fmt.Printf("%c \n",ultimo)
}