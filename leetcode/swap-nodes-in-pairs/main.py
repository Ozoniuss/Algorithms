from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        current = head
        beforeCurrent: Optional[ListNode] = None
        while current != None:
            if current.next == None:
                return head

            n = current.next
            current.next = n.next
            n.next = current

            # first already processed
            if beforeCurrent != None:
                beforeCurrent.next = n
                beforeCurrent = current
            else:
                beforeCurrent = current
                head = n

            current = current.next

            
        return head
