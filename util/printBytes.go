package util

import (
	"fmt"
)

// PrintBytes prints the bytes in a human-readable format.
// example:
// ```
// 00 | 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 0123456789abcdef
// 16 | 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f 0123456789abcdef
// ```

func PrintBytes(buf []byte) {
	for i, b := range buf {
		if i%16 == 0 {
			fmt.Printf("\n%02x | ", i)
		}
		fmt.Printf("%02x ", b)
	}

	fmt.Println()
}
