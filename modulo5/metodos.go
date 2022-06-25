package main

import "fmt"

type User struct {
	Name  string //ATRIBUTOS
	Email string
}

func (self *User) setName(name string) {
	self.Name = name
}
func (self *User) getName() string {
	return self.Name
}

func main() {
	cody := User{Name: "cody", Email: "yanethdd"}
	//meodo
	cody.setName("virgini")

	fmt.Println(cody.getName())

}
