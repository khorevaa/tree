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
	africa.Add(location("South Africa"))
	europe := earth.Add(location("Europe"))
	sweden := europe.Add(location("Sweden"))
	europe.Add(location("Norway"))
	sweden.Add(location("Skåne"))

	fmt.Println("Tree size:", tree.Size())
	fmt.Println(tree)
	recPrint(tree.Root)
}

func recPrint(node *tree.Node) {
	if !node.HasChildren() {
		fmt.Println(node)
	} else {
		for _, n := range node.Children {
			recPrint(n)
		}
		fmt.Println(node)
	}
}
