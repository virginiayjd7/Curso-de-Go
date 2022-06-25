package main

import "fmt"

func main()  {
	// // numero := []int{1,2,3,4}//un slide no existe sin un arreglo
	
	// // numero=append(numero,5)//eagrega lementos al final de un lide
	// // numero=append(numero,7)
	// // numero=append(numero,8)
	// // numero=append(numero,10)

	// // newslice:=numero[0:5]
	// // secondslice :=numero[0:3]
	

	// // numero[0]=100

	// // fmt.Println(numero)
	// // fmt.Println(newslice)
	// // fmt.Println(secondslice)


    // meses:=[]string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio",
    // "Agosto","Septiembre"}

    // // Un puntero
    // // Una longitud
    // // Una capacidad
    // longitud:=len(meses)
    // capacidad:=cap(meses)

    // fmt.Printf("La longitud es:%v, la capacidad es:%v",longitud,capacidad)

	// meses= append(meses, "Octubre")//si la estructura se encuentra en su capacidad maxima se gera un unveo arreglo
    // meses= append(meses, "noviembre")
	// meses= append(meses, "dicmber")

	// longitud =len(meses)
    // capacidad =cap(meses)

	// fmt.Printf("La longitud es:%v, la capacidad es:%v %p\n",longitud,capacidad,meses)

	numero:= []int{1,2,3,4,5,6}


	inicio := numero[0:3]
	final :=numero[3:6]
	
	numero[0]=100
	numero[5]=600

	fmt.Println(numero)
	fmt.Println(inicio)
	fmt.Println(final)
}

