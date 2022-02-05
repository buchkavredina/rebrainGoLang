package main

import (
	"fmt"
	"strconv"
)

func main() {

	var stroka string = "104"
	var chislo int = 35

	st, _ := strconv.Atoi(stroka)
	fmt.Println(st)
	fmt.Println(strconv.Itoa(chislo))
}
