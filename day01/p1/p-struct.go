package main

import "fmt"

type TreeNode struct {
	Value int
	LeftNode * TreeNode
	RightNode * TreeNode
}

func (root *TreeNode) Traverse() {
	if root == nil{
		return
	}
	root.LeftNode.Traverse()
	fmt.Println(root.Value)
	root.RightNode.Traverse()
}

func main() {
	root := TreeNode{Value:1}
	n1 := TreeNode{Value:2}
	n2 := TreeNode{Value:3}
	n3 := TreeNode{Value:4}
	root.LeftNode = &n1
	root.RightNode = &n2
	n1.RightNode = &n3
	root.Traverse()
}
