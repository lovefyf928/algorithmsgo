package trie

import (
	"errors"
	"fmt"
)

type trie []node

type node struct {
	isWord bool
	val int
	next []map[rune]node
}

func createTrie() trie {
	var trie trie
	return trie
}

func (node node) inNode(r rune) (bool, node) {
	for _, nowNode := range node.next{
		for key, nextNode := range nowNode{
			if r == key {
				return true, nextNode
			}
		}
	}
	return false, node
}

func (trie2 trie) appendWordToTrie(str string, val int) trie {
	if len(trie2) < 1 {
		trie2 = append(trie2, node{next: []map[rune]node{}, isWord: false})
	}
	trie2[0].next = trie2.create(str, 0, 1, trie2[0], val).next
	fmt.Println(trie2[0].next)
	return trie2
}

func (trie trie) create(originStr string, start int, end int, nowNode node, val int) node {
	str := originStr[start:end]
	var nextKey rune
	for _, v := range str{
		nextKey = v
	}
	if len(nowNode.next) > 0 {
		for i := 0; i < len(nowNode.next); i ++ {
			for key, v := range nowNode.next[i]{
				if key == nextKey {
					if len(originStr) != end {
						nowNode.next[i][key] = trie.create(originStr, start + 1, end + 1, v, val)
						return nowNode
					} else {
						v.val = val
						v.isWord = true
						return nowNode
					}
				} else if i == len(nowNode.next) - 1 {
					if end == len(originStr) {
						nowNode.next = append(nowNode.next, map[rune]node{nextKey: {next: []map[rune]node{}, isWord: true, val: val}})
						return nowNode
					} else {
						nowNode.next = append(nowNode.next, map[rune]node{nextKey: trie.create(originStr, start + 1, end + 1, createChildNode(nowNode, nextKey, false), val)})
						return nowNode
					}
				}
			}
		}
		return node{}
	} else {
		if end == len(originStr) {
			nowNode.next = append(nowNode.next, map[rune]node{nextKey: {next: []map[rune]node{}, isWord: true, val: val}})
			return nowNode
		} else {
			nowNode.next = append(nowNode.next, map[rune]node{nextKey: trie.create(originStr, start + 1, end + 1, createChildNode(nowNode, nextKey, false), val)})
			return nowNode
		}
	} 
}

func createChildNode(nowNode node, char rune, isWold bool) node {
	 nowNode.next = append(nowNode.next, map[rune]node{char: {next: []map[rune]node{}, isWord: isWold}})
	 return nowNode.next[len(nowNode.next) - 1][char]
}

func (trie trie) searchWordInTree(word string) string {
	return trie.searchWordInTree1(word, trie[0], 0, 1)
}

func (trie trie) simpleMatchPrefix(word string) bool {
	//inputPoint := ""
	points := []node{trie[0]}
	//strArr := []string{}
	if len(word) > 0 {
		for _, r := range word{
			if r == 46 {
				tmpArr := []node{}
				for i := 0; i < len(points); i ++ {
					for _, nextNodeMap := range points[i].next {
						for _, nextNode := range nextNodeMap {
							//if len(nextNode.next) > 0 {
								tmpArr = append(tmpArr, nextNode)
							//} else {
							//	appendNode := []map[rune]node{}
							//	appendNode = append(appendNode, map[rune]node{nodeKey: nextNode})
							//	tmpArr = append(tmpArr, node{next: appendNode, isWord: false})
							//}
						}
					}
					points = append(points[:i], points[i + 1:]...)
					i --
				}
				points = append(points, tmpArr...)
			} else {
				for i := 0; i < len(points); i ++ {
					if ok, nextNode := points[i].inNode(r); ok {
						//if len(nextNode.next) > 0 {
							points[i] = nextNode
						//} else {
						//	points[i] = node{next: []map[rune]node{{r: nextNode}}}
						//}
					} else {
						points = append(points[:i], points[i + 1:]...)
						i --
						if len(points) < 1 {
							return false
						}
					}
				}
			}
		}
		fmt.Println(len(points))
		fmt.Println(points)
		if len(points) > 0 {
			return true
		}
		//for _, point := range points{
		//	for _, str := range searchStrInNode(point) {
		//
		//	}
		//}

	}
	return false
}


func (trie trie) searchWordPrefix(word string) ([]string, int) {
	sg := false
	point := trie[0]
	strArr := []string{}
	if len(word) > 0 {
		for _, r := range word {
			if ok, nextNode := point.inNode(r); ok {
				point = nextNode
			} else {
				return strArr, 0
			}
		}
		if point.isWord {
			sg = true
			strArr = append(strArr, word)
		}
		sum := 0
		for _, strMap := range searchStrInNode(point){
			for val, str := range strMap{
				sum += val
				strArr = append(strArr, word + str)
			}
		}
		if sg {
			sum += point.val
		}
		return strArr, sum
	}
	return strArr, 0
}

func searchStrInNode(nextNode node) []map[int]string  {
	var nowString []map[int]string
	if len(nextNode.next) > 0 {
		for _, nowNode := range nextNode.next{
			for r, rNode := range nowNode {
				if rNode.isWord && len(rNode.next) > 0 {
					nowString = append(nowString, map[int]string{rNode.val: string(r)})
				}
				strArr := searchStrInNode(rNode)
				if len(strArr) > 0 {
					for _, strMap := range strArr {
						for val, str := range strMap{
							nowString = append(nowString, map[int]string{val: string(r) + str})
						}
					}
				} else {
					nowString = append(nowString, map[int]string{rNode.val: string(r)})
				}
			}
		}
		return nowString
	} else {

		return nowString
	}
}


func (trie *trie) searchWordInTree1(word string, nextNode node, wordStart int, wordEnd int) string {
	if len(word) == wordStart {
		if nextNode.isWord {
			return "find " + word
		} else {
			return "not find " + word
		}
	} else {
		strRune := stringToRune(word[wordStart:wordEnd])
		for i := 0; i < len(nextNode.next); i ++{
			for key, childNode1 := range nextNode.next[i]{
				if strRune == key {
					return trie.searchWordInTree1(word, childNode1, wordStart + 1, wordEnd + 1)
				} else if i == len(nextNode.next) - 1 {
					return "not find " + word
				}
			}
		}
		return "not find " + word
	}
}

func stringToRune(char string) rune {
	//fmt.Println(len(char))
	if len(char) != 1 { //  如输入字符为中文将判断放开
		panic(errors.New("输入的字符有问题"))
	} else {
		var r rune
		for _, v := range char{
			r = v
			break
		}
		return r
	}
}

func runeToString(r rune) string {
	return string(r)
}