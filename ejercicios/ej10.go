package ejercicios

import (
	bt "github.com/untref-ayp2/data-structures/binary-tree"
)

func LadronDeCasas(arbol *bt.BinaryTree[int]) int {
	return ladronDeCasas(arbol.GetRoot())
}

func ladronDeCasas(nodo *bt.BinaryNode[int]) int {
	return 0
}
