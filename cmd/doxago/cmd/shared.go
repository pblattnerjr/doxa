// Shared functions for the doxago commands
package cmd

import (
	"fmt"
	"time"
)

func Elapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
	Doxa()
}
func Doxa() {
	fmt.Println(GLORY)
}
