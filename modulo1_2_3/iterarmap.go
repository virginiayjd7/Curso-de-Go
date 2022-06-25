package main

import "fmt"

func main()  {
	
user := map[int] string{}
user[1]="user 1"
user[2]="user 2"
user[3]="user 3"
user[4]="user 4"

for id,username := range user{
	fmt.Println(id,username)
}
}