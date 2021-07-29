package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil { return false }
	fastIndex, slowIndex := head.Next, head
	for slowIndex != nil && fastIndex != nil && fastIndex.Next != nil {
		if fastIndex == slowIndex {
			return true
		}
		fastIndex = fastIndex.Next.Next
		slowIndex = slowIndex.Next
	}
	return false
}


