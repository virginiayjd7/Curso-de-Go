package main
import "fmt"

type Animal interface{
	Comer()
	Dormir()
	Cazar()

}
type Perro struct{
	Nombre string
}
func(self Perro) Comer(){
	fmt.Println("EL PERRO COME")
}
func(self Perro) Dormir(){
	fmt.Println("EL PERRO duerme")
}
func(self Perro) Cazar(){
	fmt.Println("EL PERRO caza")
}
func ejecutaraction(Animal Animal){
	Animal.Comer()
	Animal.Dormir()
	Animal.Cazar()
}
func main(){

	loki:=Perro{Nombre: "loki"}
	//fmt.Println(loki)
	ejecutaraction(loki)

}