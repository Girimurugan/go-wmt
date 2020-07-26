package main


import "fmt"

func main(){

	x := 1
	p := &x // p is of time *int (pointer int)

	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(p)
	*p = 2
	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(p)

}