package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ali Farhan", "Ali"))
	fmt.Println(strings.Split("Ali Farhan", " "))
	fmt.Println(strings.ToLower("Ali Farhan"))
	fmt.Println(strings.ToUpper("Ali Farhan"))
	fmt.Println(strings.Trim("           Ali Farhan           ", " "))
	fmt.Println(strings.ReplaceAll("Ali Farhan Ali Farhan", "Ali Farhan", "Alfa"))
}
