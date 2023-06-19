import os
import random

random.seed(3)

with open("nums5000000.txt", "w") as f:
    for i in range(5000000):
        n = random.randint(0, 9)
        f.write(str(n))
    f.write('\n')
    for i in range(5000000):
        n = random.randint(0, 9)
        f.write(str(n))