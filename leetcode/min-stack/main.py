from collections import deque

class MinStack:

    def __init__(self):
        self.elements = []
        self.min = deque()

    def push(self, val: int) -> None:
        self.elements.append(val)

        if len(self.min) == 0:
            self.min.append(val)
            return
        
        # edge case, element being inserted is smaller than all elements
        # and is not captured in for loop
        if val < self.min[0]:
            self.min.appendleft(val)
            return

        for i, el in enumerate(self.min):
            if val < el:
                self.min.insert(i, val)
                return

        # didn't get inserted in the for loop, insert it here
        self.min.append(val)

    def pop(self) -> None:
        last = self.elements.pop()

        # it's sufficient to remove the first occurence
        self.min.remove(last)


    def top(self) -> int:
        return self.elements[-1]        

    def getMin(self) -> int:
        return self.min[0]



class MinStackWithMinElement:

    def __init__(self):
        self.elements = []
        self.min = 1000000000000

    def push(self, val: int) -> None:
        if val < self.min:
            self.min = val
        self.elements.append((val, self.min))
    

    def pop(self) -> None:
        last = self.elements.pop()
        if len(self.elements) == 0:
            self.min = 1000000000000
            return

        if self.elements[-1][1] > self.min:
            self.min = self.elements[-1][1]

    def top(self) -> int:
        return self.elements[-1][0]

    def getMin(self) -> int:
        return self.elements[-1][1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
        
minStack = MinStackWithMinElement()
minStack.push(-2)
minStack.push(0)
minStack.push(-3)
print(minStack.getMin()) # return -3
minStack.pop()
print(minStack.top())    # return 0
print(minStack.getMin()) # return -2