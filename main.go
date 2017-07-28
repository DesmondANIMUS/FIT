// A program that provides fnv hash of a file

package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	hash := fnv.New64()

	io.Copy(hash, f)
	fmt.Println(hash.Sum64())
	fmt.Scan()
}
