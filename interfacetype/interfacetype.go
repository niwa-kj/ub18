package main

import "fmt"

func main(){
	var x, y interface{}
	x = 1
	fmt.Printf("x: %d y:%d\n", x, y)
	x = 2
	y = 3
	fmt.Printf("x: %d y:%d\n", x, y)
	x = 4
	y = "string"
	fmt.Printf("x: %d y:%s\n", x, y)
	switch yt := y.(type){
		case int:
			fmt.Printf("y:%d type: %s \n", y, yt)
		case string:
			fmt.Printf("y:%s type: %s \n", y, yt)
		default:
			fmt.Printf("y:%#v type: %s \n", y, yt)
	}
}

