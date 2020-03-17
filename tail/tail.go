package main

import "fmt"
import  "github.com/hpcloud/tail"

func main(){
	t, err := tail.TailFile("./test.log", tail.Config{Follow: true})
	if err != nil {
        fmt.Println("TailFile ERROR", err)
	}

    for line := range t.Lines {
	    if err != nil {
            fmt.Println("TailLines ERROR", line.Err)
        }
        fmt.Println("LOG " , line.Text)
        fmt.Println("Date " , line.Time)
    }
}

