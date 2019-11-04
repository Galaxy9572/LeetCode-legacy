package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/solution/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
