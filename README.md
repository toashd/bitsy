# Short URL Generator and Coder

Go package for securely generating (encoding/decoding) TinyURL- and bit.ly-like URLs.

[![Build Status](https://travis-ci.org/toashd/bitsy.svg)](https://travis-ci.org/toashd/bitsy)

## Usages

The intended use is that incrementing, consecutive integers will be used as keys to generate the short URLs. For example, when creating a new URL, the unique integer ID assigned by a database could be used to generate the URL - or a simple counter may be used. As long as the same integer is not used twice, the same short URL will not be generated twice.

## Algorithm description
Package bitsy provides a `Coder` that uses a bit-shuffling approach to avoid the generation of consecutive, predictable strings and URLs, respectively. The algorithm is deterministic and ensures that no collisions will occur.

The URL `alphabet` is fully customizable and may contain any number of characters. By default, digits and lower-case letters are used, with some removed to avoid confusion between characters like o, O and 0. The default alphabet is shuffled and has a prime number of characters to further improve the results of the algorithm.

The `blockSize` specifies how many bits will be shuffled. The lower blockSize bits are reversed. Any bits higher than blockSize will remain as is. blockSize of 0 will leave all bits unaffected and the algorithm will simply be converting your integer to a different base.

The package supports both, encoding and decoding of URLs. The `minLength` parameter enables to pad the URL if you want it to be of specific length.

## Getting Started

1: Download the package

```bash
$ go get github.com/toashd/bitsy
```

2: Import bitsy to your Go project

```go
import "github.com/toashd/bitsy"
```

## Example

```go
package main

import (
	"fmt"
	"github.com/toashd/bitsy"
)

func main() {
	// The id, e.g database id, of the url to decode.
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
```

This example can also be found in the example.go file.

## Contribution

Please feel free to suggest any kind of improvements and refactorings.

Get in touch, file an issue, fork and submit a pull request.

## Get in touch

Tobias Schmid, toashd@gmail.com, [@toashd](http://twitter.com/toashd), [toashd.com](http://toashd.com)

## License

bitsy is available under the MIT license. See the LICENSE file for more info.
