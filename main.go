package main

import "fmt"

func main() {
	fmt.Println("Hello, cometBFT")
	_ = NewKVStoreApplication(nil)
}
