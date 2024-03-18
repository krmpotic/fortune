package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <file>\n", os.Args[0])
	os.Exit(1)
}

func RandElement[K any](slice []K) K {
	return slice[rand.IntN(len(slice))]
}

func RandLine(b []byte) string {
	lines := strings.Split(string(b), "\n")
	return RandElement(lines)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(RandLine(b))
}
