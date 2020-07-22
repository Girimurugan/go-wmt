package main

import "fmt"

func main() {

	fmt.Println(spm(3, 2))
	
	s, p, m := spm(3, 2)
	fmt.Println(s, p, m)

}

func spm(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}
