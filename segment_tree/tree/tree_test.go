package tree

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T)  {
	segmentTree := CreateTree([]int{9, 40, 20, -9, 4})
	segmentTree.createNode(segmentTree.arr, 1)
	segmentTree.tree = segmentTree.tree[:segmentTree.length * 4]
	fmt.Println(segmentTree.tree)
	fmt.Println(segmentTree.selectTree(0, 3))
	segmentTree.updateTree(1, 10)
	segmentTree.updateTree(3, 9)
	fmt.Println(segmentTree.selectTree(0, 3))
}

func Test1(t *testing.T)  {
	//fmt.Println(Testf1(1, 2))
	fmt.Println(Testf2(1))
}


func Testf1(a int, b int) int {
	if b > 10 {
		return 1
	}
	c := Testf1(a + 2, b + 2)
	return a * b * c
}

func Testf2(a int64) int64 {
	if a < 10 {
		return Testf2(a + 1) * a
	} else {
		return  1
	}
}