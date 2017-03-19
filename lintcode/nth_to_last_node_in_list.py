"""
Definition of ListNode
class ListNode(object):

    def __init__(self, val, next=None):
        self.val = val
        self.next = next
"""
class Solution:
    """
    @param head: The first node of linked list.
    @param n: An integer.
    @return: Nth to last node of a singly linked list. 
    """
    def nthToLast(self, head, n):
        nth = head
        for _ in xrange(n):
            nth = nth.next
        if not nth:
            return head
        current = head
        while nth.next:
            nth = nth.next
            current = current.next
        return current.next
