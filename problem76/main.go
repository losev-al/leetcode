package main

import "fmt"

func main() {
	fmt.Println(minWindow("bba", "ab"))
	fmt.Println(minWindow("abc", "bc"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))

}
