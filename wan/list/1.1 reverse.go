package main

import (
	"fmt"

	"github.com/isdamir/gotype"
)

func main() {
	head := &gotype.LNode{}
	fmt.Println("就地逆序")
	gotype.CreateNode(head, 8)
	gotype.PrintNode("逆序前:", head)
	insertReverse(head)
	gotype.PrintNode("逆序后:", head)
}

// Reverse 带头结点的逆序
func Reverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *gotype.LNode
	var cur *gotype.LNode
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

// 插入法
func insertReverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *gotype.LNode
	var next *gotype.LNode

	cur = node.Next.Next
	node.Next.Next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}
