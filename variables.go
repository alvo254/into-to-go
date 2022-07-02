package main

import (
	"fmt"
	"net/http"
	// "net/http"
)

//Starting variale with capital letters means its going to be exported
//variable can be defined with var or const keywords
//this is global scope
var i float64 = 21.3
var j int = 43

func formHundle(w http.ResponseWriter, r *http.Request){

}

func main(){
	// fileserver := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileserver)
	http.HandleFunc("/form", formHundle)
	fmt.Printf("%v %v\n", i, j)
	fmt.Printf("Enter name: ")
	var sname string
	fmt.Scan(&sname)
	fmt.Println(sname)


	fmt.Println("Welcome to go ",sname)
}