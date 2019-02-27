package main

import (
	"go-lazy/assets"
	"os"
)

//Package name
//Struct name
//attribute:type
func main() {
	var structGenerator assets.StructGenerator
	structGenerator.GenerateClass(os.Args)
}
