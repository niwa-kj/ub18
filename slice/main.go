package main

import "fmt"

func main(){

	var s0 []int
	fmt.Println("slice", s0)
	fmt.Println("len", len(s0))
	fmt.Println("cap", cap(s0), "\n")
	
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice", s1)
	fmt.Println("len", len(s1))
	fmt.Println("cap", cap(s1), "\n")

	s2 := make([] int , 10, 20)
	fmt.Println("slice", s2)
	fmt.Println("len", len(s2))
	fmt.Println("cap", cap(s2), "\n")

	s2 = append(s2, 4)
	fmt.Println("slice", s2)
	fmt.Println("len", len(s2))
	fmt.Println("cap", cap(s2), "\n")

	s2 = append(s2, s1...)
	fmt.Println("slice", s2)
	fmt.Println("len", len(s2))
	fmt.Println("cap", cap(s2), "\n")

	s3 := make([] int , 10, 20)
	s2 = append(s2, s3...)
	fmt.Println("slice", s2)
	fmt.Println("len", len(s2))
	fmt.Println("cap", cap(s2), "\n")

	s4 := []int{1, 2, 3, 4, 5}
	s5 := make([] int , 10, 20)
	fmt.Println("slice", s4)
	fmt.Println("slice", s5)
	copy(s5, s4)
	fmt.Println("slice", s5)
	fmt.Println("len", len(s5))
	fmt.Println("cap", cap(s5), "\n")

	for i, v := range s5 {
	    fmt.Printf("[%d]%d \n", i ,v)
	}
}

