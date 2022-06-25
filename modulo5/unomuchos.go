package main

import "fmt"

type Curso struct{

	titulo string
	videos []video
}
type video struct{//
	
	titulo string
	Curso Curso
}

func main(){
	video1:=video{titulo: "1.introduc"}
	video2:=video{titulo: "1.intalacion"}
    
    videos:= []video{video1,video2}

	Curso := Curso {titulo: "curso profesional de go",videos:videos}

	video1.Curso=Curso
	video2.Curso=Curso

	fmt.Println(video1.Curso.titulo)
	//fmt.Println(video2)
	//fmt.Println(Curso)

	for _,video := range Curso.videos{
		fmt.Println(video.titulo)
	}
}