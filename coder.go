// Package bitsy provides a coder for safely generating (encoding/decoding) TinyURL- and bit.ly-like URLs.
// The coder uses a bit-shuffling approach to avoid generating consecutive, predictable URLs.
// The algorithm is deterministic and ensures that no collisions will occur.
package bitsy

import (
	"fmt"
	"math"
	"strings"
)

// Default coder configuration parameters.
const (
	defaultAlphabet  = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	defaultBlockSize = 24
	defaultMinLength = 5
)

// Coder holds the coding characteristics.
type Coder struct {

	// alphabet is fully customizable and may contain any number of characters.
	// By default, digits and lower-case letters are used, with some removed to avoid
	// confusion between characters like o, O and 0. To further improve the results of
	// the algorithm, the alphabet's number of characters should be a prime number.
	alphabet string

	// blockSize specifies how many bits will be shuffled.
	// The lower blockSize bits are reversed. Any bits higher than blockSize
	// will remain as is. A blockSize of 0 will leave all bits unaffected and
	// the algorithm will simply convert the given integer to a different base.
	blockSize int

	// minLength pads the code to be of a specific length.
	minLength int

	// mask is calculated depending on blockSize.
	mask int

	// mapping is calculated depending on blockSize.
	mapping []int
}

// New initializes a new Coder with default values.
func New() *Coder {
	c := &Coder{
		alphabet:  defaultAlphabet,
		blockSize: defaultBlockSize,
		minLength: defaultMinLength,
		mask:      mask(defaultBlockSize),
		mapping:   mapping(defaultBlockSize),
	}
	return c
}

// Encode encodes the given integer with min length.
func (c *Coder) Encode(n int, minLength int) string {
	if minLength < c.minLength {
		minLength = c.minLength
	}
	return c.enbase(c.encode(n), minLength)
}

// Decode decodes the given string.
func (c *Coder) Decode(n string) int {
	return c.decode(c.debase(n))
}

func (c *Coder) encode(n int) int {
	return (n & ^c.mask) | c._encode(n&c.mask)
}

func (c *Coder) _encode(n int) int {
	var result = 0
	for i, m := range c.mapping {
		if n&(1<<uint(i)) != 0 {
			result |= (1 << uint(m))
		}
	}
	return result
}

func (c *Coder) decode(n int) int {
	return (n & ^c.mask) | c._decode(n&c.mask)
}

func (c *Coder) _decode(n int) int {
	var result = 0
	for i, m := range c.mapping {
		if n&(1<<uint(m)) != 0 {
			result |= (1 << uint(i))
		}
	}
	return result
}

// enabase converts the given integer to a different base.
func (c *Coder) enbase(x int, minLength int) string {
	var result = c._enbase(x)
	var padding = strings.Repeat(string(c.alphabet[0]), (minLength - len(result)))
	return fmt.Sprintf("%s%s", padding, result)
}

func (c *Coder) _enbase(x int) string {
	var n = len(c.alphabet)
	if x < n {
		return string(c.alphabet[x])
	}
	return c._enbase(x/n) + string(c.alphabet[x%n])
}

func (c *Coder) debase(x string) int {
	var n = len(c.alphabet)
	var result = 0
	for i, rune := range reverseString(x) {
		result += strings.Index(c.alphabet, string(rune)) * int(math.Pow(float64(n), float64(i)))
	}
	return result
}

// mask calculates mask depending on given block size.
func mask(blockSize int) int {
	return (1 << uint64(blockSize)) - 1
}

// mapping calculates mapping depending on given block size.
func mapping(blockSize int) []int {
	n := blockSize
	mapping := make([]int, n)
	for i := range mapping {
		n--
		mapping[n] = i
	}
	return mapping
}

// reverseString reverses the given string.
func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
