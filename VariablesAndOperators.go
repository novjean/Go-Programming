package main

//import format package
import "fmt"

func main(){
	
	var age int = 40
	var favNum float64 = 1.618
	//randNum := 1 			//got to use all variables or else it throws an error
	fmt.Println(age,favNum) //dont need to specify space between variables
	
	var numOne = 1.000
	var num99 = .9999
	fmt.Println(numOne - num99)   //floats are not going to be 100% accurate
	
	//Arithmetic
	fmt.Println("6+5: ", 6+5)
	fmt.Println("6/5: ", 6/5.0)
	fmt.Println("6%5: ", 6%5)
	
	//More Go types
	const pi float64 = 3.14159625
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", pi) //%T will display the datatype of the variable
	
	/*
	var(
		varA = 2
		varB = 3
	)
	*/
	
	var myName string = "Novjean J"
	fmt.Println(len(myName))
	fmt.Println(myName + " is \na robot")
	
	var isOver40 bool = true
	fmt.Printf("%t \n", isOver40)
	
	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100) //binary
	fmt.Printf("%c \n", 44) //character value
	fmt.Printf("%x \n", 17) //hexa
	fmt.Printf("%e \n", pi)	//scientific
	
	
}
