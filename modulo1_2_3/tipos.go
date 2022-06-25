package main

import (
    "fmt"
    "strconv"
)
func main() {

	edad := 22 //variable string 
    edad_int := strconv.Itoa(edad)//convertir de string  entero(atoi)

	fmt.Println("Mi edad es " +edad_int)
}
