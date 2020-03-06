// 引数をランダムに選択する
package main

import "fmt"
import "math/rand"
import "os"
import "time"

func main(){
	rand.Seed(time.Now().UnixNano())
	c := len(os.Args) - 1
	fmt.Printf("len %d \n", c)
	if c < 1 {
		fmt.Fprintf(os.Stderr,"[usage] %s choice1 choice2...\n",os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("choice %s \n", os.Args[rand.Intn(c) + 1])
}
