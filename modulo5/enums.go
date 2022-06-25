package main
import "fmt"

type UserType int//tipo de dato
const(
	Teacher UserType=1//listado de contsntes
	Student UserType=1
)
type User struct{//limitando solo studen y theach
	Username string
	Type UserType
}

func main(){
	
	yaneth :=User{Username: "yaneth",Type: Student}
	vicky :=User{Username: "vicky",Type: Teacher}

	fmt.Println(yaneth)
	fmt.Println(vicky)
	if yaneth.Type==Student{
		fmt.Println("el usuri es estuante",yaneth.Username)
	}


}