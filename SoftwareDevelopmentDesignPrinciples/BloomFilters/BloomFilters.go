package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash"
)

type filter struct {
	bitfield [100]bool
}

var (
	hasher = sha1.New()
)

func main() {
	f := filter{}
	f.set("www.yahoo.com")
	f.set("hello")
	f.set("world")
	fmt.Printf("%v\n", f.bitfield)

	fmt.Printf("www.yahoo.com Get Result: %v\n", f.get("world"))
	fmt.Printf("hello Get Result: %v\n", f.get("hello"))
	fmt.Printf("world Get Result: %v\n", f.get("world"))

}

func (f *filter) set(s string) {
	f.bitfield[f.hashPosition(s)] = true
}

func (f *filter) get(s string) bool {
	return f.bitfield[f.hashPosition(s)]
}

func (f *filter) hashPosition(s string) int {
	hs := createHash(hasher, s)
	if hs < 0 {
		hs = -hs
	}
	return hs % len(f.bitfield)
}

func createHash(h hash.Hash, input string) int {
	bits := h.Sum([]byte(input))
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	return int(result) // cast the int64
}
