package main

import (
	"fmt"

	"github.com/lunjon/tree"
)

type location string

func main() {
	tree := tree.Create()

	earth := tree.Root.Add(location("Earth"))

	africa := earth.Add(location("Africa"))
	africa.Add(location("South-Africa"))

	europe := earth.Add(location("Europe"))
	sweden := europe.Add(location("Sweden"))
	europe.Add(location("Norway"))
	skåne := sweden.Add(location("Skåne"))
	skåne.Add(location("Lund"))

	// fmt.Println("Tree size:", tree.Size())
	fmt.Println(tree)
}
