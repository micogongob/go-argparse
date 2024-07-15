package internal

import (
	"fmt"
	"os"
)

func Fail(msg string) {
	fmt.Printf("ERROR: %v\n", msg)
	os.Exit(1)
}
