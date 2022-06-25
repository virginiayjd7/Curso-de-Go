package main
import "fmt"

type User struct{

	Name string
	Email string
	Active bool
}
type Student struct{//
	User User
	Id string
}

func main(){
	
yaneth:=User{Name: "yaneth",Email: "djdjd",Active: true}

vicky := User{Name: "vicky",Email: "kkckc" ,Active: true}
yanethstudent :=Student{User: yaneth,Id: "1f1"}
fmt.Println(vicky)
fmt.Println(yanethstudent.User.Active)

}