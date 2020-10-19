package tree

import (
	"fmt"
	"math"
	"sync"
)

var (
	mutex sync.Mutex
	wg sync.WaitGroup
)


type segmentTree struct {
	arr []int
	tree []int
	length int
}

type rage struct {
	form int
	to int
}

func CreateTree(arr []int) *segmentTree {
	return &segmentTree{ arr: arr, length: len(arr), tree: make([]int, len(arr) * 20)}
}

func (segmentTree *segmentTree) selectTree(i int, j int) int {
	return segmentTree.eachTree(0, 0, segmentTree.length - 1, i, j)
}

func (segmentTree *segmentTree) updateTree(index int, val int) {
	segmentTree.eachTreeUpdate(index, val, 0, 0, segmentTree.length - 1)
}

func (segmentTree *segmentTree) eachTreeUpdate(index int, val int, treeIndex int, left int, right int)  {
	leftChildFrom := left
	leftChildTo := int(math.Ceil(float64(left) + (float64(right) - float64(left)) / 2) - 1)
	if leftChildFrom > leftChildTo {
		leftChildTo = leftChildFrom
	}
	rightChildFrom := leftChildTo + 1
	rightChildTo := right
	if index == left && index == right {
		segmentTree.tree[treeIndex] = val
		return
	} else {
		originData := segmentTree.arr[index]
		segmentTree.tree[treeIndex] = segmentTree.tree[treeIndex] - originData + val
	}
	if index > leftChildTo {
		segmentTree.eachTreeUpdate(index, val, rightChild(treeIndex), rightChildFrom, rightChildTo)
	} else {
		segmentTree.eachTreeUpdate(index, val, leftChild(treeIndex), leftChildFrom, leftChildTo)
	}
}



func (segmentTree *segmentTree) eachTree(treeIndex int, left int, right int, i int, j int) int {
	leftChildFrom := left
	leftChildTo := int(math.Ceil(float64(left) + (float64(right) - float64(left)) / 2) - 1)
	if leftChildFrom > leftChildTo {
		leftChildTo = leftChildFrom
	}
	rightChildFrom := leftChildTo + 1
	rightChildTo := right
	if i == left && j == right {
		return segmentTree.tree[treeIndex]
	} else if (i <= leftChildTo && i >= leftChildFrom) && (j <= leftChildTo && j >= leftChildFrom) {
		return segmentTree.eachTree(leftChild(treeIndex), leftChildFrom, leftChildTo, i, j)
	} else if (i >= rightChildFrom && i <= rightChildTo) && (j >= rightChildFrom && j <= rightChildTo) {
		return segmentTree.eachTree(rightChild(treeIndex), rightChildFrom, rightChildTo, i, j)
	} else {
		return segmentTree.eachTree(leftChild(treeIndex), leftChildFrom, leftChildTo, i, leftChildTo) + segmentTree.eachTree(rightChild(treeIndex), rightChildFrom, rightChildTo, rightChildFrom, j)
	}
}

func leftChild(index int) int {
	return 2 * index + 1
}

func rightChild(index int) int {
	return 2 * index + 2
}

func (segmentTree *segmentTree) createNode(rootNode []int, treeIndex int) {
		rootSum := 0
		for _, v := range rootNode{
			rootSum += v
		}
		if len(rootNode) == segmentTree.length {
			segmentTree.tree[0] = rootSum
		}
		for key, v := range segmentTree.tree{
			if v == rootSum {
				//var left []int
				//var right []int
				left, right, leftKey, rightKey := segmentTree.createChildNode(rootNode, treeIndex, key,1)
				left, right, leftKey, rightKey = segmentTree.createChildNode(rootNode, treeIndex, key,2)
					if len(left) <= 1 {
						segmentTree.createChildNode(left, treeIndex + 1, leftKey,1)
					} else {
						segmentTree.createNode(left, treeIndex + 1)
					}
					if len(right) <= 1 {
						segmentTree.createChildNode(right, treeIndex + 1, rightKey,2)
					} else {
						segmentTree.createNode(right, treeIndex + 1)
					}
				//if len(left) >= 1 {
				//	segmentTree.createChildNode(left, treeIndex + 1, leftKey,1)
				//} else if len(right) >= 1 {
				//	segmentTree.createChildNode(right, treeIndex + 1, rightKey,2)
				//} else {
				//	segmentTree.createNode(left, treeIndex + 1)
				//	segmentTree.createNode(right, treeIndex + 1)
				//}
				break
			}
		}


	//if rootNode != nil {
	//	rootSum := 0
	//	for _, v := range rootNode{
	//		rootSum += v
	//	}
	//	segmentTree.tree = append(segmentTree.tree, rootSum)
	//	left, right := segmentTree.createChildNode(rootNode)
	//	fmt.Println(rootNode)
	//	fmt.Println(left, right)
	//	segmentTree.createNode(nil, left, right)
	//} else {
	//	left, right := segmentTree.createChildNode(leftNode)
	//	segmentTree.createNode(nil, left, right)
	//	left1, right1 := segmentTree.createChildNode(rightNode)
	//	segmentTree.createNode(nil, left1, right1)
	//}


	//wg.Add(1)
	//mutex.Lock()
	//if len(node) > 1 {
	//	if len(node) == segmentTree.length {
	//		rootSum := 0
	//		for _, v := range node{
	//			rootSum += v
	//		}
	//		segmentTree.tree = append(segmentTree.tree, rootSum)
	//	}
	//	mid := int(math.Floor(float64(len(node) / 2)))
	//	leftNode := node[:mid]
	//	rightNode := node[mid:]
	//	leftSum := 0
	//	rightSum := 0
	//	for _, v := range leftNode {
	//		leftSum += v
	//	}
	//	for _, v := range rightNode {
	//		rightSum += v
	//	}
	//	segmentTree.tree = append(segmentTree.tree, leftSum)
	//	segmentTree.tree = append(segmentTree.tree, rightSum)
	//	fmt.Println(leftNode, rightNode)
	//	fmt.Println(segmentTree.tree)
	//	mutex.Unlock()
	//	go segmentTree.createNode(leftNode)
	//	time.Sleep(time.Microsecond * 10)
	//	go segmentTree.createNode(rightNode)
	//} else {
	//	segmentTree.tree = append(segmentTree.tree, node...)
	//	mutex.Unlock()
	//	fmt.Println(segmentTree.tree)
	//}
}

func (segmentTree *segmentTree) createChildNode(node []int, treeIndex int, parentIndex int, nodePosition int) ([]int, []int, int, int) {
	mid := len(node) / 2
	leftNode := node[:mid]
	rightNode := node[mid:]
	leftSum := 0
	rightSum := 0
	for _, v := range leftNode {
		leftSum += v
	}
	for _, v := range rightNode {
		rightSum += v
	}
	if nodePosition == 1 {
		segmentTree.tree[2 * parentIndex + 1] = leftSum
	} else if nodePosition == 2 {
		segmentTree.tree[2 * parentIndex + 2] = rightSum
	}
	fmt.Println(leftNode, rightNode)
	fmt.Println(leftSum, rightSum)
	return leftNode, rightNode, treeIndex * parentIndex + 1, treeIndex * parentIndex + 2
}