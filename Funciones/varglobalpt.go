package main
import "fmt"
import "os"

func main(){
    defer func ()  {
        if err:=recover(); err!=nil{
            fmt.Println("ups , al parece el promra no finalizo")
        }
        
    }()
    

    if file, err :=os.Open("example.txt"); err != nil{
        panic("no fue posible obetener el archivo")

    }else{
        defer func(){
            fmt.Println("cerramos el archivo")
            file.Close()
        }()
    
        contenido:= make([]byte,254)
    
        longitud,_ := file.Read(contenido)
    
        conenidoarch := string(contenido[0:longitud])
    
        fmt.Println(conenidoarch)
    
    }

    
}