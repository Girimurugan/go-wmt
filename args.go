package main

import (
    "fmt"
	"os"
	"strconv"
)

func main() {

	var s int64

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	
	for _, arg := range os.Args[1:]{

		n, _ := strconv.ParseInt(arg, 10, 10)
		s = s + n
	}
	fmt.Printf("Total is %d", s)
	fmt.Println()
}