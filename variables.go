package main

import ("fmt"
	
	)

//Starting variale with capital letters means its going to be exported
//variable can be defined with var or const keywords
//this is global scope
var i float64 = 21.3
var j int = 43

func main(){
	fmt.Printf("%v %v", i, j)


	fmt.Println("Welcome to ge")
}