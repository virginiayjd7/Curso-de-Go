/*Ejercicio 3:
Instrucciones: pide al usuario que indique su nombre y su edad. Como mensaje de salida le indicarás qué edad tuvo el año pasado y cuantos años tendrá el siguiente año.
Ejemplo: [nombre] el año pasado tenías X años y el próximo año cumplirás Y años.*/
package main
import "../jj"
import (
	"fmt"
	"time"
)
func main()  {

	var nombre string
	fmt.Print("Ingrese su nombre:")
	fmt.Scanf("%s\n", &nombre)

	var edad  int
	fmt.Printf("Ingrese su edad:")
	fmt.Scanf("%d\n", &edad)

	edadpasada:=edad - 1
	edadproxima:= edad + 1

	fmt.Printf("Hola %s ,el año pasado tenías %d años y el próximo año cumplirás %d años\n",nombre,edadpasada,edadproxima)
	
	time.Now().Add(90)
	ch.send(
		send(ch,7)
		ch<7
		7->sh
	)
	
}

