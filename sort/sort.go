package sort

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type sortStruct struct {
	arr []int
	sg bool
	maxCount int
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func SortCompare(n int, arrNum int) {
	rand.Seed(time.Now().UnixNano())
	var finalArr [][]int
	for j := 0; j < arrNum; j ++ {
		var arr []int
		for i := 0; i < n; i ++ {
			arr = append(arr, rand.Intn(n * 10000))
		}
		finalArr = append(finalArr, arr)
	}
	fmt.Println("数据生成完毕")
	for i := 0; i < len(finalArr); i ++ {
		ss := sortStruct{arr: finalArr[i]}
		start := getTimestamp()
		ss.margeSortEntree()
		//ss.shellSort2()
		//sort.Ints(ss.arr)
		use := getTimestamp() - start
		fmt.Println("第" + strconv.Itoa(i + 1) + "数组排序完成，耗时" + strconv.FormatFloat(float64(use) / float64(1000), 'f', 6, 64) + "s")
	}
}

func (sortStruct *sortStruct) less(a int, b int) bool {
	return a <= b
}

func (sortStruct *sortStruct) exch(aIndex int, bIndex int)  {
	tmp := sortStruct.arr[aIndex]
	sortStruct.arr[aIndex] = sortStruct.arr[bIndex]
	sortStruct.arr[bIndex] = tmp
}


// 选择排序
func (sortStruct *sortStruct) selectSort() {
	for i := 0; i < len(sortStruct.arr); i ++ {
		min := sortStruct.arr[i]
		key := i
		for j := i + 1; j < len(sortStruct.arr); j ++ {
			if sortStruct.arr[j] < min {
				min = sortStruct.arr[j]
				key = j
			}
		}
		sortStruct.exch(i, key)
	}
}

// 冒泡排序
func (sortStruct *sortStruct)bubbleSort() {
	for i := 0; i < len(sortStruct.arr); i ++ {
		for j := i + 1; j < len(sortStruct.arr); j ++ {
			if !sortStruct.less(sortStruct.arr[i], sortStruct.arr[j]) {
				sortStruct.exch(i, j)
			}
		}
	}
}




// 希尔排序(f)
func (sortStruct *sortStruct)shellSort()  {
	h := 4
	count := len(sortStruct.arr) / h
	startKey := 0
	for h > 1 {
		for k := 0; k < h; k++ {
			for i := startKey + 1; i < startKey + count; i++ {
				for j := i; j > startKey && !sortStruct.less(sortStruct.arr[j-1], sortStruct.arr[j]); j-- {
					sortStruct.exch(j, j-1)
				}
			}
			startKey += count
		}
		count *= 2
		startKey = 0
		h = h / 2
	}
	sortStruct.insertingSort()
}

// 希尔排序2(t)
func (sortStruct *sortStruct)shellSort2() {
	n := len(sortStruct.arr)
	h := 1
	for h < n / 3 {
		h = 3 * h + 1
	}
	for h >= 1 {
		for i := h; i < n; i ++ {
			for j := i; j >= h && sortStruct.less(sortStruct.arr[j], sortStruct.arr[j - h]); j -= h {
				sortStruct.exch(j, j - h)
			}
		}
		h = h / 3
	}
}


// 归并排序
func (sortStruct *sortStruct) margeSortEntree() {
	lo := 0
	hi := len(sortStruct.arr) - 1
	var tmp = make([]int, hi - lo + 1)
	sortStruct.sg = false
	sortStruct.inSituMargeSort(lo, hi, tmp, 0)
}


// 插入排序
func (sortStruct *sortStruct)insertingSort()  {
	for i := 1; i < len(sortStruct.arr); i ++ {
		for j := i; j > 0 && !sortStruct.less(sortStruct.arr[j - 1], sortStruct.arr[j]) ; j -- {
			sortStruct.exch(j, j - 1)
		}
	}
}


// 原地归并 (改进)
func (sortStruct *sortStruct)marge(lo int, hi int, mid int, tmp []int, count int) {
	if sortStruct.arr[mid] <= sortStruct.arr[mid + 1] {
		return
	}
	length := hi - lo + 1
	if length < 15 {
		for i := 1; i < length; i ++ {
			for j := i; j > 0 && sortStruct.less(sortStruct.arr[j + lo], sortStruct.arr[j - 1 + lo]) ; j -- {
				sortStruct.exch(j + lo, j - 1 + lo)
			}
		}
	} else {
		j := lo
		n := mid + 1
		//for i := lo; i <= hi; i ++ {
		//	tmp[i] = sortStruct.arr[i]
		//}
		//		for i := lo; i <= hi; i ++ {
		//			if j > mid {
		//				sortStruct.arr[i] = tmp[n]
		//				n ++
		//			} else if n > hi {
		//				sortStruct.arr[i] = tmp[j]
		//				j ++
		//			} else if sortStruct.less(tmp[n], tmp[j]) {
		//				sortStruct.arr[i] = tmp[n]
		//				n ++
		//			} else {
		//				sortStruct.arr[i] = tmp[j]
		//				j ++
		//			}
		//		}
		//fmt.Println(sortStruct.maxCount)
		//fmt.Println(count)
		//fmt.Println(lo, hi)
		//fmt.Println(sortStruct.arr)


	if (sortStruct.maxCount - 1) % 2 == 0 {
		if count % 2 != 0 {
			for i := lo; i <= hi; i ++ {
				if j > mid {
					sortStruct.arr[i] = tmp[n]
					n ++
				} else if n > hi {
					sortStruct.arr[i] = tmp[j]
					j ++
				} else if sortStruct.less(tmp[n], tmp[j]) {
					sortStruct.arr[i] = tmp[n]
					n ++
				} else {
					sortStruct.arr[i] = tmp[j]
					j ++
				}
			}
		} else {
			fmt.Println("in tmp arr")
			for i := lo; i <= hi; i ++ {
				if j > mid {
					tmp[i] = sortStruct.arr[n]
					n ++
				} else if n > hi {
					tmp[i] = sortStruct.arr[j]
					j ++
				} else if sortStruct.less(sortStruct.arr[n], sortStruct.arr[j]) {
					tmp[i] = sortStruct.arr[n]
					n ++
				} else {
					tmp[i] = sortStruct.arr[j]
					j ++
				}
			}
		}
	} else {
		if count % 2 == 0 {
			for i := lo; i <= hi; i ++ {
				if j > mid {
					sortStruct.arr[i] = tmp[n]
					n ++
				} else if n > hi {
					sortStruct.arr[i] = tmp[j]
					j ++
				} else if sortStruct.less(tmp[n], tmp[j]) {
					sortStruct.arr[i] = tmp[n]
					n ++
				} else {
					sortStruct.arr[i] = tmp[j]
					j ++
				}
			}
		} else {
			for i := lo; i <= hi; i ++ {
				if j > mid {
					tmp[i] = sortStruct.arr[n]
					n ++
				} else if n > hi {
					tmp[i] = sortStruct.arr[j]
					j ++
				} else if sortStruct.less(sortStruct.arr[n], sortStruct.arr[j]) {
					tmp[i] = sortStruct.arr[n]
					n ++
				} else {
					tmp[i] = sortStruct.arr[j]
					j ++
				}
			}
		}
	}

	if (sortStruct.maxCount - 1) % 2 == 0 && count == 0 {
		sortStruct.arr = tmp
	}
	}
}



// 递归分割数组
func (sortStruct *sortStruct) inSituMargeSort(lo int, hi int, tmp []int, count int) {
	if lo == hi {
		if !sortStruct.sg {
			sortStruct.maxCount = count
			sortStruct.sg = true
		}
		if count < sortStruct.maxCount {
			tmp[lo] = sortStruct.arr[lo]
		}
		return
	}
	mid := lo + (hi - lo) / 2
	sortStruct.inSituMargeSort(lo, mid, tmp, count + 1)
	sortStruct.inSituMargeSort(mid + 1, hi, tmp, count + 1)
	sortStruct.marge(lo, hi, mid, tmp, count)
}



















