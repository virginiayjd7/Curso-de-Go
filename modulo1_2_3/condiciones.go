package main

import "fmt"
/*
== igual a
!= diferente de
< menor que 
> mayor que 
>= mayor que o igual que
&& and evalua v  cuando las conciones son verdaderas
||  or  solo una verdare  para ser verdadero

*/

func main() {
	yaneth := 10
	virginia := 15
	if yaneth > virginia {
		fmt.Printf("%d es mayor que %d \n", yaneth, virginia)
	}else if yaneth < virginia{
		fmt.Printf("%d es mayor que %d \n", virginia,yaneth)
	}else{
		fmt.Printf("son iguales")
	}
}
