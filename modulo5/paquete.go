package main

import "fmt"
import "./codigofacilito"

func main()  {
	curso := codigofacilito.Curso{Titulo:"fffff"}
	fmt.Println(curso.ObtenerTitulo())
}