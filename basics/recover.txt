package main

import "fmt"

func main() {
	fmt.Println("server started")
	process()
	fmt.Println("returned from process")
}

func process() {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	fmt.Println("start processing...")
	panic("something went wrong")
}
