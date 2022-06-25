package main
import "fmt"

type Animal interface {
	Dormir()
	
}
type Mascota interface {
	Comer()
}
type Gato struct {
	Nombre string
}
func (self Gato)Dormir()  {
	fmt.Println("EL GATO DUERME")
	
}
func (self Gato)Comer()  {
	fmt.Println("EL GATO come")
	
}
func funcionparaunanimal(animal Animal)  {
	fmt.Println("el objeto es un animal")
}
func funcionparaunmascota(mascota Mascota)  {
	fmt.Println("el objeto es un mascota")
}
func main(){
	Gato := Gato{Nombre: "mi gato"}
	Gato.Dormir()
	Gato.Comer()
	funcionparaunanimal(Gato)
	funcionparaunmascota(Gato)




}