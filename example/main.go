package main

import (
	"fmt"

	"github.com/toashd/bitsy"
)

func main() {
	// The id of the url to decode
	var url = 1337

	// Create new bitsy coder
	c := bitsy.New()

	// Encode url with length of 5
	var enc = c.Encode(url, 5)

	// Print encoded url
	fmt.Printf("Encoded url: %s\n", enc)

	// Print decoded url
	fmt.Printf("Decoded url: %v\n", c.Decode(enc))
}
