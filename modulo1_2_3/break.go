package main

import "fmt"

func main()  {
	//break continue
	for i:=1;i<=10;i++{
		if i==5 {//
			continue//finaliza la altucal  o sea el 5 y procede el siguiente 6 saltar
		}
		if i==8{

			break //finaliza el actual finaliza y no visualiza despues del numero 8 finalizar
		}
		fmt.Println(i)
	}
}