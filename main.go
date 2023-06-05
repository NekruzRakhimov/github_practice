package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hey World!")

	time.Sleep(2 * time.Second)

	fmt.Println("Bye World!")
}