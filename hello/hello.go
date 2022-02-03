package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	formatted := fmt.Sprintf("%02d.%02d.%d",
		t.Day(), t.Month(), t.Year())
	fmt.Println("hello", formatted)
}
