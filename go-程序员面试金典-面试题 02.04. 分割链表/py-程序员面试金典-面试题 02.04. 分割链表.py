# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    """
编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。
如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。
分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8
链接：https://leetcode-cn.com/problems/partition-list-lcci
    """
    def partition(self, head, x):
        """
        :type head: ListNode
        :type x: int
        :rtype: ListNode
        """
        if not head:
            return head
        aux = pre = ListNode(-1)
        aux.next = cur = head
        while cur.next:
            if cur.next.val < x:
                pre.next, cur.next.next, cur.next = cur.next, pre.next, cur.next.next
            else:
                cur = cur.next
        return aux.next


def create(nums):
    aux = p = ListNode(-1)
    for x in nums:
        p.next = ListNode(x)
        p = p.next
    return aux.next


def main():
    nums, x = [3, 5, 8, 5, 10, 2, 1], 5
    nums, x = [], 0
    test = Solution()
    ret = test.partition(create(nums), x)
    print(ret)


if __name__ == '__main__':
    main()
