package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaMayoresQue(bt *binarytree.BinaryTree[int], x int) int {
	if bt.IsEmpty() {
		return 0
	}

	return sumaMayoresQue(bt.GetRoot(), x)
}

func sumaMayoresQue(n *binarytree.BinaryNode[int], x int) int {
	return 0
}
