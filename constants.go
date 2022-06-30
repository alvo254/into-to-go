package main

import "fmt"

//constant has to be assignable at comiletime

//Naming convention

//Typed constants

//Untyped constants

// iota - this is a counter that can be used to create enumerated constants
const a = iota
//iota is nomarly scoped to the constant block
//_ right only varable dont care what this variable is
const (
	//if you dont allocate the next iota ie c,d the compiler is going to infer the parten that comes next
	b = iota
	c = iota
	d = iota
)

func main(){
	const myConst int = 23
	//var b = "something"
	fmt.Println("welcome to constats")
	//The %T is used to show typeof
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}