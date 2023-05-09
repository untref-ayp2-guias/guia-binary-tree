package main

import (
	"fmt"
	"guia12/binarytree"
)

func main() {
	node1 := binarytree.NewBinaryTree("B", nil, nil)
	node2 := binarytree.NewBinaryTree("C", nil, nil)
	node3 := binarytree.NewBinaryTree("-", node1, node2)
	node4 := binarytree.NewBinaryTree("D", nil, nil)
	node5 := binarytree.NewBinaryTree("*", node3, node4)
	arbol := binarytree.NewBinaryTree("+", nil, node5)
	fmt.Println("-----------PrintPreOrder-----------")
	arbol.PrintPreOrder()
	fmt.Println("-----------PrintPostOrder----------")
	arbol.PrintPostOrder()
	fmt.Println("-----------PrintInOrder------------")
	arbol.PrintInOrder()
}
