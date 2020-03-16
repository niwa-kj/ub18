package main

// https://godoc.org/net/http

import "fmt"
import "net/http"

    // ハンドラ関数
func indexHandler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hello, %q", req.URL.Path)
}

func main(){
	// ハンドラ定義用関数
	http.HandleFunc("/bar", indexHandler)
	http.ListenAndServe(":8080", nil)
}
