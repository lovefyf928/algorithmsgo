package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T)  {
	tree := createTrie()
	tree = tree.appendWordToTrie("cat", 5)
	tree = tree.appendWordToTrie("have", 55)
	tree = tree.appendWordToTrie("has", 87)
	tree = tree.appendWordToTrie("sad", 99)
	tree = tree.appendWordToTrie("carry", 50)
	tree = tree.appendWordToTrie("cay", 3)
	tree = tree.appendWordToTrie("cut", 10)
	tree = tree.appendWordToTrie("you", 20)
	tree = tree.appendWordToTrie("pan", 3)
	tree = tree.appendWordToTrie("panda", 10)
	//fmt.Println(tree.searchWordPrefix("panda"))
	fmt.Println(tree.simpleMatchPrefix(".an"))
	fmt.Println(tree.searchWordPrefix("c"))
}