package main

import "flag"
import "fmt"

func main() {
    // オプション名、デフォルト値、helpテキスト
    flgs := flag.String("s", "", "String")
	flgi := flag.Int("i", 1, "Int default:1")
	flgb := flag.Bool("b", false, "Bool default:false")

    flag.Parse()

    fmt.Println("flgs: ", *flgs)
    fmt.Println("flgi: ", *flgi)
    fmt.Println("flgb: ", *flgb)
  }
