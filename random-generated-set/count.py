from collections import Counter

import matplotlib as mpl
import matplotlib.pyplot as plt
import numpy as np

x = [1, 2, 3, 2, 3, 4, 5, 5, 4, 3]

c = Counter(x)

plt.plot(c.keys(), c.values())

plt.show()
