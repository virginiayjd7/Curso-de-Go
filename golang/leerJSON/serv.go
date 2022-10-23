package main

import (
	"io"
	"net/http"
)
func saludo(w http.ResponseWriter,peticion *http.Request)  {
	io.WriteString(w,"hola mundo")
}
func main() {
	http.HandleFunc("/",saludo)
	http.ListenAndServe(":8080",nil)


}