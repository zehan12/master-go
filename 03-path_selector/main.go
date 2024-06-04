package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css")

	var fileName string
	_, fileName = path.Split("javascript/index.js")

	fmt.Println("dir:", dir, "file:", file)
	fmt.Println("file:", fileName)
}
