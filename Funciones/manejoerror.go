package main

import "fmt"
import "errors"
func diviion(dividento,divisor int)(int,error)  {
	if divisor ==0{
		return 0,errors.New("no es posible dividir sobre 0")
	}else{
		return dividento / divisor, nil//nil ausencia de valor
	}
	
}

func main()  {
	if resul,err := diviion(10,0);err!=nil {
		//fmt.Println(err)
		panic(err)
	}else{
		fmt.Println("el resutado es:",resul)
	}
	
}