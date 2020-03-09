package main

import "flag"
import "fmt"

func main(){
	flag.Parse()
	args := flag.Args()
	fmt.Println("args len():", len(args))
	for i := 0; i < len(args); i++{
		fmt.Printf("args[%d]:%s\n", i, args[i])
	}
	fmt.Println(args)
}

