package main
import "fmt"

func main(){ //bloque 1
	var curso="curso proefecional de go"
	var version=10
	{ //bloque2
		fmt.Println(curso)
		{
			//bloque 3
			var version=5
			fmt.Println("estmos en bloque 3",version)
			fmt.Println(curso)
		}
		fmt.Println("estamos en el bloque 2",version)
	}
	fmt.Println(curso)
//cada varible es de un bloque
}