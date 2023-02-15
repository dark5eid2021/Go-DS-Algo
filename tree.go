package main

import (
	"fmt"
)

// This is the tree struct
type Tree struct {
	Leftnode  *Tree
	Value     int
	RightNode *Tree
}

// Tree insertion method. m = middle position
func (tree *Tree) insert(m int) {
	if tree != nil {

		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {

				if tree.Leftnode != nil {

					tree.Leftnode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

// print method for the tree
func print(tree *Tree) {
	if tree != nil {

		fmt.Println(" Value", tree.Value)
		fmt.Printlf("Tree Node Left")
		print(tree.Leftnode)
		fmt.Print("Tree Node Right")
		print(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

// main method
func main() {
	var tree *Tree = &Tree{nil, m, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.Leftnode.insert(7)
	print(tree)
}
