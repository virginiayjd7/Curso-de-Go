package main

import "fmt"

func main()  {
	
	/*
	&& comparar todos ve
	||
	!
	
	*/

	//resul:=20 == 20 && 30==30 && 20>40
	//resul:=20 == 20 || 30 ==30 || 20 > 40 basta que uno sea verdadero
	//resul:=20 != 20 || 30 != 30 || 20>40
             //  v        v           v         f
			 //               v       v
			 //                    v
	//resul := 15 == 15 && 60 < 100 &&(90<100 ||100==90)
	resul:= !true
	fmt.Println(resul)
}