package AVL

import (
	"math"
	"sort"
)

type avlTree struct {
	tree *node
}

type node struct {
	val int
	left *node
	right *node
	treeHeight int
}

func (node *node) getLeftChild() *node {
	return node.left
}

func (node *node) getRightChild() *node {
	return node.right
}

//func (avlTree *avlTree) createNode(val int, left *node, right *node) {
//	avlTree.tree = &node{right: right, left: left, val: val}
//}

func (node *node) appendToLeft(newNode *node) {
	node.left = newNode
}

func (node *node) appendToRight(newNode *node) {
	node.right = newNode
}

func createNewNode(val int) *node {
	return &node{val: val, left: nil, right: nil}
}

func createAvlTree() avlTree {
	return avlTree{tree: nil}
}

func (avlTree *avlTree) appendDataToTree(sourceData []int) {
	mid, data := findMid(sourceData)
	avlTree.tree = createNewNode(mid)
	for _, val := range data{
		sourceDataToCreateTree(avlTree.tree, val)
	}
}

func sourceDataToCreateTree(nowNode *node, val int) {
	if val > nowNode.val {
		if nowNode.getRightChild() != nil {
			sourceDataToCreateTree(nowNode.getRightChild(), val)
		} else {
			nowNode.appendToRight(createNewNode(val))
			return
		}
	} else {
		if nowNode.getLeftChild() != nil {
			sourceDataToCreateTree(nowNode.getLeftChild(), val)
		} else {
			nowNode.appendToLeft(createNewNode(val))
			return
		}
	}
}

func findMid(sourceData []int) (int, []int) {
	sort.Ints(sourceData)
	mid := int(math.Floor(float64(len(sourceData) / 2)))
	midData := sourceData[mid]
	sourceData = append(sourceData[:mid], sourceData[mid + 1:]...)
	return midData, sourceData
}