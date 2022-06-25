package main
import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "os/exec"
)

var reader *bufio.Reader //global para utilia en culqui funio

type User struct{
	id int
	username string
	email string
	age int
}
var id int
//ALMACENAR EN MAPAS
var users map[int]User//

func crearUsuario()  {
	crearUsuario()
	fmt.Print("ingrese un valor user:")
	username := readline()

	fmt.Print("ingrese un valor email:")
	email :=readline()

	fmt.Print("ingrese un valor edad:")
	age,err := strconv.Atoi(readline())

	if err != nil{
		panic("no es posible converit de un sting aun enter")
	}

    id++
	user := User { id, username, email, age }
	users[id] = user
	fmt.Println(users)

	fmt.Println(">>>>usuariocreado exitoso!\n")
}
func listarUsuario()  {
	clearConsole()

	for id,user := range users{
		fmt.Println(id,"_",user.username)
	}
	fmt.Println("\n")

	//fmt.Println("usuario listado")
}
func updateUsuario()  {
	fmt.Println("usuario actualizado exitoso")
}
func deleteUsuario()  {
	clearConsole()
	fmt.Print("ingrese el id el usurioa elminr")
	id,err:= strconv.Atoi(readline())
	if err !=nil{
		panic("no es posibre converti de un strin a un etro")
	}
	if _, ok :=users[id];ok{
		delete(users,id)
	}
	fmt.Println("><><>usuario eliinado exitoso!\n")
}
func clearConsole()  {
	cmd := exec.Command("clear")
	cmd.Stdout =os.Stdout
	cmd.Run()
}
func readline() string {

	if option, err := reader.ReadString('\n');err != nil{
		panic("no es posible obtener el valr")
	}else{
        //fmt.Println(option)
        return strings.TrimSuffix(option,"\n")
	}
	
}
func main()  {

	
	var option string//se inico con u stricg vacio
	users = make(map[int]User)
	reader=bufio.NewReader(os.Stdin)//lo que el usuaro ingrese
	
	for{
		fmt.Println("A) Crear")
		fmt.Println("B) listar")
		fmt.Println("C) actualizar")
		fmt.Println("D) eliminar")

		fmt.Print("ingree un aopcion valida:")
		
		option= readline()
	
		if option =="quit"|| option =="q" {
			break
		}
		switch option{
		    case "a" ,"crear":
		    	crearUsuario()
		    case "b","listar":
		    	listarUsuario()
		    case "c","actualizar":
		    	updateUsuario()
		    case "d","eliminar":
		    	deleteUsuario()
		    default	:
		    fmt.Println("opcio no vaida")	
		}
	}

	fmt.Println("adios")
}