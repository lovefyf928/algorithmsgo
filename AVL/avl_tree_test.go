package AVL

import (
	"fmt"
	"testing"
)

func Test(t *testing.T)  {
	//node := node{val: 1, left: nil, right: nil}
	//fmt.Println(node.getLeftChild())
	arr := []int{23,4,6,3,2,1,333}
	tree := createAvlTree()
	tree.appendDataToTree(arr)
	fmt.Println(tree.tree.getRightChild().getRightChild())
}
