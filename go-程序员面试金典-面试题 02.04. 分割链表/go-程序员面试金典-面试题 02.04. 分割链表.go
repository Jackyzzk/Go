package main

import "fmt"
/*
编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。
如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。
分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8
链接：https://leetcode-cn.com/problems/partition-list-lcci
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func create(nums []int) *ListNode {
	aux := ListNode{-1, nil}
	p := &aux
	for _, x := range nums {
		next := ListNode{x, nil}
		p.Next = &next
		p = p.Next
	}
	return aux.Next
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	aux := ListNode{-1, nil}
	aux.Next = head
	pre, cur := &aux, head
	for cur.Next != nil {
		if cur.Next.Val < x {
			pre.Next, cur.Next.Next, cur.Next = cur.Next, pre.Next, cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return aux.Next
}

func main()  {
	//nums, x := []int{3, 5, 8, 5, 10, 2, 1}, 5
	nums, x := []int{}, 0
	ret := partition(create(nums), x)
	fmt.Println(ret)
}
