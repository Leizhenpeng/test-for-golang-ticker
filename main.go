package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(3 * time.Second)

	go func() {
		fmt.Println("hello") // First execution
		for range ticker.C {
			fmt.Println("hello")
		}
	}()

	select {}
}
