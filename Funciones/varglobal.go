package main
import "fmt"

var username string 
func fun1(){
	username="user1"
	fmt.Println("funcion1",username)
}
func fun2(){
	username="user1"
	fmt.Println("funcion2",username)
}
func main(){
    fmt.Println("hola no escontramo main")
    username="yaneth"
    defer fun1()//se ejecuta cando el bloque finaliza para validar eerorres
    fun2()
    fmt.Println(username)

}