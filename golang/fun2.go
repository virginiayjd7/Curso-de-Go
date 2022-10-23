package main

import "fmt"

func gando (num int)(int,string){
	vacas := func () int  {

		return num * 10
	}
	return vacas()+1,"vacas"
}

func main() {
	fmt.Print(gando(1))
}