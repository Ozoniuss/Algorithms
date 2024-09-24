class MinStack:

    def __init__(self):
        self.minstack = []
        self.minlist = []

    def push(self, val: int) -> None:
        self.minstack.append(val)
        if len(self.minlist) == 0 or val < self.minlist[-1]:
            self.minlist.append(val)
        else:
            self.minlist.append(self.minlist[-1])

    def pop(self) -> None:
        self.minstack.pop()
        self.minlist.pop()

    def top(self) -> int:
        return self.minstack[-1]

    def getMin(self) -> int:
        return self.minlist[-1]
