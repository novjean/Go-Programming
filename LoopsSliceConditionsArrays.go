package main

import "fmt"

func main(){

	//For Loop
	
	i :=1
	for i<=10{
		fmt.Println(i)
		i++
	}
	
	for j:=0; j<=5; j++{
		fmt.Println(j)
	}
	
	//If-else statement
	
	yourAge := 18
	if yourAge >= 16{
		fmt.Println("You can drive")
	}else if yourAge>=18{
		fmt.Println("You can vote")
	}else{
		fmt.Println("You can have fun")
	}
	
	//Switch statement
	
	switch yourAge{
		case 16: fmt.Println("Go Drive")
		case 18: fmt.Println("Go Vote")
		default: fmt.Println("Go have fun")
	}
		
	//Arrays
	
	var favNums[5] float64					//declaration

	favNums[0] = 025
	favNums[1] = 123
	favNums[2] = 234
	favNums[3] = 369
	favNums[4] = 452
	fmt.Println(favNums[3])
	
	favNums2 := [5]float64 {2,5,3,5,1}		//alternative declaration
	
	for i, value:= range favNums2 {			//print with the indexes in the loop
		fmt.Println(value, i)
	}
	
	for _, value:=range favNums2 { 			//print without the indexes
		fmt.Println(value)
	}
	
	//Slice
	
	numSlice := []int {3, 5, 2, 3, 4}		//array of int as slice
	fmt.Println("numSlice[:2] =", numSlice[:2])	//display till 2nd index
	fmt.Println("numSlice[2:] =", numSlice[2:])	//display till 2nd index
	
	numSlice2 := numSlice[3:5]				//gets the values from numSlice into numSlice2
	fmt.Println("numSlice2[1] =", numSlice2[1])
	
	numSlice3 := make([]int, 5, 10)     			//default value of 0 for 5 indexes, slice has a max index length of 10
	fmt.Println("numSlice3[1] =", numSlice3[1])
	
	copy(numSlice3,numSlice)						//copy numSlice into numSlice3
	fmt.Println("numSlice3[1] =", numSlice3[1])
	
	numSlice3 = append(numSlice3, 9 ,-1)  			//append to numSlice3, adding 0 and 1 to the array
	fmt.Println("numSlice3[5] =", numSlice3[5])
}
