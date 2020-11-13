package main

import "fmt"

/*
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
如果指定节点没有对应的“下一个”节点，则返回null。
输入: root = [2,1,3], p = 1
  2
 / \
1   3
输出: 2
示例 2:
输入: root = [5,3,6,2,4,null,null,1], p = 6
      5
     / \
    3   6
   / \
  2   4
 /
1
输出: null
链接：https://leetcode-cn.com/problems/successor-lcci
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	var ret *TreeNode
	pre := false
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre {
			ret = node
			pre = false
		}
		if node == p {
			pre = true
		}
		dfs(node.Right)
	}
	dfs(root)
	return ret
}

func main() {
	none := 1381
	nums, p := []int{2, 1, 3}, 1
	nums, p = []int{2, none, 3}, 2
	//nums, p = []int {5, 3, 6, 2, 4, none, none, 1}, 6
	ret := inorderSuccessor(create(nums, p))
	fmt.Println(ret)
}

func create(nums []int, p int) (*TreeNode, *TreeNode) {
	none := 1381
	if len(nums) == 0 {
		return nil, nil
	}
	var ret *TreeNode
	root := &TreeNode{nums[0], nil, nil}
	nums = nums[1:]
	que := []*TreeNode{root}

	for len(que) != 0 {
		node := que[0]
		if node.Val == p {
			ret = node
		}
		que = que[1:]
		if len(nums) != 0 {
			if nums[0] != none {
				node.Left = &TreeNode{nums[0], nil, nil}
				que = append(que, node.Left)
			}
			nums = nums[1:]
		}
		if len(nums) != 0 {
			if nums[0] != none {
				node.Right = &TreeNode{nums[0], nil, nil}
				que = append(que, node.Right)
			}
			nums = nums[1:]
		}
	}
	return root, ret
}