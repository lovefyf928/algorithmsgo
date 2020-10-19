package main

import (
	"algorithms/union_find/node"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	nodeMap []node.Node
	count int
	mapSize int
	ch = make(chan string, 1)
)

func initNodeMap() {
	fmt.Println("请输入map大小")
	mapSize, _ = strconv.Atoi(getInput())
	count = mapSize
	for i := 0; i < mapSize; i ++ {
		nodeMap = append(nodeMap, node.Node{Node: i, ConnectNode: []int{i}})
	}
}

func getInput() string {
	bf := bufio.NewReader(os.Stdin)
	str, _, err := bf.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func main() {
	initNodeMap()
	fmt.Println(nodeMap)
	for {
		select {
		case res := <- ch:
			fmt.Println(res)
			eventLoop()
		default:
			eventLoop()
		}
	}
}

func eventLoop() {
	fmt.Println("请输入q")
	q, _ := strconv.Atoi(getInput())
	fmt.Println("请输入p")
	p, _ := strconv.Atoi(getInput())
	if q > mapSize - 1 || p > mapSize - 1 {
		ch <- "输入的p｜q不在map中"
	} else {
		if !(nodeMap[q].Fund(nodeMap[p]) || nodeMap[p].Fund(nodeMap[q])) {
			nodeMap = nodeMap[q].UpdateConnectNode(&nodeMap[p], nodeMap)
			count --
			fmt.Println(q, p)
		} else {
			fmt.Println(q, p, "已连接")
		}
		fmt.Println(nodeMap[q].ConnectNode)
		fmt.Println(nodeMap[p].ConnectNode)
	}
}