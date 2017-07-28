// A program that provides fnv hash of a file

package main

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"os"
)

func main() {
	var name string

	fmt.Print("Enter name of file: ")
	fmt.Scanln(&name)

	f, err := os.Open(name)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	hash := fnv.New64()

	io.Copy(hash, f)
	fmt.Println("FNV hash of the file:", hash.Sum64())

	hash2 := md5.New()

	io.Copy(hash2, f)
	// fmt.Println(hash2.Sum(nil)), will print a slice of bytes

	fmt.Printf("MD5 hash of the file: %x", hash2.Sum(nil))

	fmt.Scanln(&name)
}
