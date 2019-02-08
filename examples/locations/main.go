package main

import (
	"fmt"

	"github.com/lunjon/tree"
)

type location string

func main() {
	earth := tree.Create("Earth")

	africa := earth.Root.Add(location("Africa"))
	africa.Add(location("South Africa"))
	europe := earth.Root.Add(location("Europe"))
	sweden := europe.Add(location("Sweden"))
	europe.Add(location("Norway"))
	sweden.Add(location("Skåne"))

	fmt.Println("Tree size:", earth.Size())
	fmt.Println(earth)
}
