package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	var n int
	var f float64
	var flag bool
	var s string
	var b byte
	var r rune

	fmt.Printf("%[1]v:%[1]T\n%[2]v:%[2]T\n%[3]v:%[3]T\n%[4]q:%[4]T\n%[5]v:%[5]T\n%[6]v:%[6]T\n", n, f, flag, s, b, r)

	size := 3_500_000 // байт
	mb := float64(size) / MB
	fmt.Printf("%.2f MB\n", mb)

}
