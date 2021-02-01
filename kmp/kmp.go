package kmp

import "fmt"

func searchModelStr(str string) (arr []int) {
	arr = make([]int, len(str))
	fmt.Println(arr)
	tmpStr := ""
	for key, v := range str {
		tmpStr += string(v)
		if key > 0 {
			fmt.Println(tmpStr)
			arrV := 0
			for tmpKey, _ := range tmpStr {
				if tmpKey == len(tmpStr) - 1 {
					break
				}
				fmt.Println(tmpStr[:tmpKey + 1])
				fmt.Println("--------")
				fmt.Println(tmpStr[len(tmpStr) - (tmpKey + 1):])
				fmt.Println("----end-----")
				if tmpStr[:tmpKey + 1] == tmpStr[len(tmpStr) - (tmpKey + 1):] {
					arrV = len(tmpStr[:tmpKey + 1])
				}
			}
			arr[key] = arrV


			if key > 0 && arrV != 0 {
				for i := key; i >= 0; {
					if v != int32(str[key + 1]) {
						break
					}
					if v == int32(tmpStr[len(tmpStr) - (len(tmpStr) - arr[i]) - 1]) {
						arr[key] = arr[len(tmpStr) - (len(tmpStr) - arr[i]) - 1]
						fmt.Println(arr[i])
						if arr[key] == 0 {
							break
						} else {
							i = (len(tmpStr) - (len(tmpStr) - arr[i])) - 1;
						}
					} else {
						break
					}
				}
			}
		}
	}
	return
}

func strHasMark(mark string, str string) (index int) {
	next := searchModelStr(mark)
	//for _, v := range str{
	//	for key, m := range mark{
	//		if v != m && key == 0 {
	//			break
	//		} else if v != m {
	//			key = key - (key - next[key - 1])
	//		} else if key == len(mark) - 1 {
	//			index =
	//		}
	//	}
	//}
	//j := 0
	//for key, v := range str{
	//	if v != int32(mark[j]) && j != 0 {
	//		j = j - (j - next[j - 1])
	//	}
	//	if v == int32(mark[j]) && j == len(mark) - 1 {
	//		fmt.Println(key, "------key");
	//		index = (key + 1) - len(mark)
	//		break
	//	}
	//	j ++
	//}
	//return

	index = -1
	i := 0
	j := 0
	if len(str) < len(mark) {
		return
	}
	for i < len(str) {
		if str[i] != mark[j] && j != 0 {
			j = j - (j - next[j - 1])
		} else if str[i] == mark[j] && j == len(mark) - 1 {
			index = (i + 1) - len(mark)
			break
		} else if str[i] == mark[j] {
			j ++
			i ++
		} else {
			i ++
		}
	}
	return
}

func checkStrHasMark(mark string, startKey int, str string) string {
	markLen := len(mark)
	return str[startKey:(startKey + markLen)]
}
