package main

import "fmt"

func alumnos(alumno ...string/*indefinida/*,edad,direccion*/)  {
	for _, valor :=range alumno{
		fmt.Print(valor)


	}
	fmt.Print()
	
}

func main() {
	alumnos("yanth","ddd","hddh")
}

func areaRectangulo(base, altura float32) {
    return base * altura
}
fun areaRectangulo(base, altura float32) float32 {
    return base * altura
}