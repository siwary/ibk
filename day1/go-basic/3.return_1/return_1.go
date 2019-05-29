package main

import "fmt"

func main(){

	fmt.Println(testReturn())
}

func testReturn() (string, string, int){

	return "1","2", 3
}