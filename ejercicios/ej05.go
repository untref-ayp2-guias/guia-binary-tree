package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaMayoresQue(bt *binarytree.BinaryTree[int], k int) int {
	if bt.IsEmpty() {
		return 0
	}

	return sumaMayoresQue(bt.GetRoot(), k)
}

func sumaMayoresQue(n *binarytree.BinaryNode[int], k int) int {
	return 0
}
