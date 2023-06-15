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

	key3 := "Hello, my Twilio key is: TwilioKey-1234567890123456789012345678901234"
	fmt.Println(key3)

	key4 := "Hello, my Twilio key is: TwilioKey-1234567890123456789012345678901234 and https://mainnet.infura.io/v3/5d07495d44b745c2b2340053d7d4d932"
	fmt.Println(key4)

	fmt.Println(c)
	fmt.Println("main print")
}