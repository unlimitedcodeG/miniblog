package main

import "fmt"

func mayPanic() {
	panic("a problem")
}



func SliceDemo(){
	var s1 [] = int 
	s2 := make([]int,9)
	fmt.Println(s2) 
	s1 = make([]int,9,9)
	print(s1)
	s1[0]  = 0 
	s1[1] = 1 
	fmt.Println(s1[0:])	
	s1 = append(s1, 1)
	println("---",s1[0])

}
func main() {
	SliceDemo()
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
