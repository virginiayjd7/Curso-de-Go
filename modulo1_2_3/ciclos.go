//solo hay ciclo for
package main

import "fmt"

func main ()  {
	// for i:=1; i<10;i=i+3 {//interador i
	// 	fmt.Println(i)

	// }
	i:=0
	for{
		if i== 2{
			i++//para cambiar la variable o agregar ya no valdria 2 
			continue
		}
		fmt.Println(i)
		i++
		if i>10{
			break
		}
	}
}