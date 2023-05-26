package main

import (
	"fmt"
)

func main() {
	var a, b int
	c := a + b

	key := "https://mainnet.infura.io/v3/5d07495d44b745c2b2340053d7d4d932"
	fmt.Println(key)

	key2 := "https://mainnet.infura.io/v3/5d07495d44b745c2b2340053d7d4d932"
	fmt.Println(key2)

	key3 := "Hello, my Twilio key is TwilioKey12345678901234567890123456789012"
	fmt.Println(key3)

	fmt.Println(c)
	fmt.Println("main print")
}