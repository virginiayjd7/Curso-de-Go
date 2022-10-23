package main

import "fmt"

func main() {

	/*var arrelo [2]string
	arrelo[0]="manzana"
	arrelo[1]="kk"*/
	/*arrelo:=[2]string{"yanwth","jjjj"}
	arrelo[1]="hhh"

	fmt.Print(arrelo)*/
	//arary multidim

	/*var  frutas [2][2]string
	//fruta1
	frutas[0][0]="manzana"
	frutas[0][1]="goldon"
	//fruta 2
	frutas[1][0]="uvas"
	frutas[1][1]="moscatel"

	fmt.Print(frutas[1][1])*/

	var frutas=[]string{"man","jjj","lll"}//indefinida lideilimita[] [2] no solo dos //infomcamion queinevtri de frutas slide,almacenar datos usuaro arrays
	//var frutas2[]string
	frutas=append(frutas,"peras")//slide agreda el valor
	fmt.Println(frutas[0:4])
	fmt.Print(len(frutas))
}

