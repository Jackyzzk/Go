package main

import "fmt"

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
/*
给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即 617 + 295
输出：2 -> 1 -> 9，即 912
进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即 617 + 295
输出：9 -> 1 -> 2，即 912
链接：https://leetcode-cn.com/problems/sum-lists-lcci
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	length := func(ln *ListNode) int {
		i := 0
		for ln != nil {
			i += 1
			ln = ln.Next
		}
		return i
	}
	if length(l1) < length(l2) {
		l1, l2 = l2, l1
	}
	p1, p2 := l1, l2
	pre := &ListNode{}
	//var pre *ListNode
	for p1 != nil && p2 != nil {
		p1.Val += p2.Val + carry
		carry = p1.Val / 10
		p1.Val %= 10
		pre = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil && carry != 0 {
		p1.Val += carry
		carry = p1.Val / 10
		p1.Val %= 10
		pre = p1
		p1 = p1.Next
	}
	if carry != 0 {
		next := ListNode{carry, nil}
		pre.Next = &next
	}
	return l1
}


func main() {
	//nums1, nums2 := []int{7, 1, 6}, []int{5, 9, 2}
	nums1, nums2 := []int{5}, []int{5}
	ret := addTwoNumbers(create(nums1), create(nums2))
	fmt.Println(ret)
}
