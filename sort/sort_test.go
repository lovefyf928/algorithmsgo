package sort

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestSort(t *testing.T) {
	//ss := sortStruct{arr: []int{5, 4, 9, 8, 3, 2, 5, 4, 9, 8, 3, 2, 24, 21, 23, 12, 123, 4544, 12, 32, 123, 32, 18}}
	//ss := sortStruct{arr: []int{3, 2, 24, 21, 23, 12, 21, 22, 98}}

	//ss.margeSortEntree()
	//ss.shellSortForMarge(11, 11)
	//fmt.Println(ss.arr)
	sortCompare(10000000, 100)
	//ss := sortStruct{arr: []int{32,12,3}}
	//fmt.Println(!ss.less(12, 3))
	//ss.insertingSort()
	//sortCompare(440000000, 100)
	//ss.shellSort2()
	//sort.Ints(ss.arr)
	//fmt.Println(ss.arr)
	//for i := 0; i < 5; i ++ {
	//	wg.Add(1)
	//	go sortCompare(10000000, 100, &wg)
	//}
	//wg.Wait()
}