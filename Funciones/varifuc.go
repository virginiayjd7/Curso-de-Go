package main
import "fmt"
//variadic fuction
func promedio(calificaciones ...int)int  {

	var promedio,suma int

	for _, calificacion := range calificaciones{
		suma=suma+calificacion

	}
	promedio=suma / len(calificaciones)
	return promedio
}

func main(){
	
	resultado := promedio(10,9,5,2,7)
	fmt.Println("promedio",resultado)

}