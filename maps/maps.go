package main

import "fmt"

func main(){
	// intのキーとstringの値のmap
    m := make(map[int]string)
    fmt.Println("map", m)
    m[1]= "subaru"
	m[2]= "honda"
	m[100] = "Daimler"

    fmt.Println("map", m, "\n")
    fmt.Println("map[100]", m[100])
    fmt.Println("map[101]", m[101], "\n")

	// キーが無い場合は""とfalseが入る
	s1, b1 := m[100]
	s2, b2 := m[101]
    fmt.Printf("map[100] %s %v \n", s1, b1)
    fmt.Printf("map[101] %s %v \n\n", s2, b2)

	// rangeは順不同
	for ky, vl := range m {
        fmt.Printf("range %d  %s \n", ky, vl)
	}

	l1 := len(m)
	m[101] = "Ford"
	l2 := len(m)
    fmt.Println("map", m)
	fmt.Printf("len():%d len():%d \n", l1, l2)

	delete(m, 100)
    fmt.Println("delete map", m)
	delete(m, 90)
    fmt.Println("delete map", m)
}
