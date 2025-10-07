package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	str := filepath.Join("/", "foo", "bar")

	fmt.Println(str)
}
