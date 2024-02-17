package main

import (
	"fmt"
	"math/rand"
)
func main(){
	//strings
	var nameOne string = "Raymond"
	var nameTwo = "Teddy"
	nameThree := "Bosire" 
	
	fmt.Println(nameOne, nameTwo, nameThree)
	fmt.Println("My favourite no is",rand.Intn(10) )
	fmt.Println("Hello Bro")

	//int
	var num1 int =1
	num2 := 7
	var num3 = 4
	
	fmt.Println(num1, num2, num3)

	//bit and memory
	var x int8 = 127 // you have to explicitly state int 16
	fmt.Println(x)

	//unsigned int - Non negative numbers. Basically numbers without a sign. You can also declare no of bits
	var y uint8 = 255
	fmt.Println(y)

	//float
	var temp float64 = 37333333333555555.33
	fmt.Println(temp)
}

