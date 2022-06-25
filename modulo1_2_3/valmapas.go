package main

import "fmt"

func main()  {
    user:= make(map[string]string)	

	user["yaneth"] ="ffffffffffff"
	//correo ,ok := user["yaneth"]
	if correo ,ok := user["ygh"];ok /*correo!=""*/ {
		fmt.Println(correo)
		
	}else{
		fmt.Println("no fue posible su valor")
	}

}