package node

type Node struct {
	Node int
	ConnectNode []int
}

func (node *Node) UpdateConnectNode(nowNode *Node, nodeMap []Node) []Node {
	arr := append(node.ConnectNode, nowNode.ConnectNode...)
	for i := 0; i < len(arr); i ++ {
		for j := i + 1; j < len(arr); j ++ {
			if arr[i] == arr[j] {
				arr = append(arr[:i], arr[i + 1:]...)
				i --
				break
			}
		}
	}
	for _, v := range arr {
		nodeMap[v].ConnectNode = arr
	}
	return nodeMap
}

func (node *Node) Fund(connectNode Node) bool {
	for _, v := range node.ConnectNode{
		if connectNode.Node == v {
			return true
		}
	}
	return false
}
