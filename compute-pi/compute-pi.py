"""
Consider a function that generates a random number between 0 and 1, uniformly distributed. Compute PI.
"""
import math
import time

import random
from math import sqrt
from typing import List

t = time.time()

# Method 1. Geometric Probability

"""
Consider a square with coordinates (0,0), (0,1), (1,0) and (1,1). Also consider a circle centered at the origin with radius 1. The first quadrant of the circle will pass through the vertices (0,1) and (1,0) of the square, and will be entirely located within the square.

The quadrant has area pi/4, and the square has area 1. 

The first and most obvious idea is to use the generator to generate a finite amount of coordinates, with values between 0 and 1. This will generate points inside the square defined previously, that can either be outside the quadrant or inside the quadrant. Using geometric probability, the probability of a point being inside the quadrant is equal to the ratio of the areas of the quadrant and the square, which is pi/4.

This means that the more coordinates we generate, the ratio between the number of points inside the quadrant and inside the square will get closer and closer to pi/4 (can this be proven rigorously?). So, we could use that ratio to "compute" pi (i.e. approximate it for a sufficiently large number of points).

It's easy to determine if a point is inside the quadrant. Given it has coordinates x and y, the sum x^2 + y^2 must be smaller than 1. Keep in mind that we always generate positive coordinates between 0 and 1.
"""

# some empty lists to store numbers for different tests
l1 = []
l2 = []
l3 = []
l4 = []
l5 = []

# range representing number of points we generate
l1 = [(random.uniform(0,1),random.uniform(0,1)) for _ in range(100)]
l2 = [(random.uniform(0,1),random.uniform(0,1)) for _ in range(1000)]
l3 = [(random.uniform(0,1),random.uniform(0,1)) for _ in range(10000)]
l4 = [(random.uniform(0,1),random.uniform(0,1)) for _ in range(100000)]
l5 = [(random.uniform(0,1),random.uniform(0,1)) for _ in range(1000000)]
l6 = [(random.uniform(0,1),random.uniform(0,1)) for _ in range(5000000)] # to check time elapsed

def estimate_pi(l : List):
    p = 0
    for coordinates in l:
        if coordinates[0] ** 2 + coordinates[1] ** 2 <= 1:
            p += 1
    return 4*(p/len(l))

t = time.time()

# estimations
e1 = estimate_pi(l1)
e2 = estimate_pi(l2)
e3 = estimate_pi(l3)
e4 = estimate_pi(l4)
e5 = estimate_pi(l5)
e6 = estimate_pi(l6)

print('Results for different inputs using geometric probability:')
print(f"100:     {e1:<{10}} error: {abs(math.pi - e1):.5f}")
print(f"1000:    {e2:<{10}} error: {abs(math.pi - e2):.5f}")
print(f"10000:   {e3:<{10}} error: {abs(math.pi - e3):.5f}")
print(f"100000:  {e4:<{10}} error: {abs(math.pi - e4):.5f}")
print(f"1000000: {e5:<{10}} error: {abs(math.pi - e5):.5f}")
print(f"5000000: {e6:<{10}} error: {abs(math.pi - e6):.5f}")

print("time elapsed: {:.3f}s".format(time.time() - t))
print("------------------------------------")
# Method 2. Riemann sum
"""
Like in the first method, consider the circle centered at the origin with radius 1, which is defined by the equation x^2 + y^2 = 1. Setting y^2 = 1-x^2, the function y = sqrt (1 - x^2) will describe the graph defined by the circumference of the cirlce of the first quadrant. Again, the quadrand has area pi/4.

We can use now a different concept to compute the area. Because the function f(x) = sqrt (1 - x^2) is positive and continuous on the interval [0,1], the area of the first quadrant will be given by the definite integral of f(x) bounded by 0 and 1. We can approximate this integral using a Riemann sum.

To construct the Riemann sum, we use the random generator to generate points between 0 and 1. We can then compute the area of the rectangle between any two consecutive points to create a Riemann sum associated with that partition of the interval, and because the points are uniformly distrubuted the sum will converge to the area.
"""



l1 = [random.uniform(0,1) for _ in range(100)]
l2 = [random.uniform(0,1) for _ in range(1000)]
l3 = [random.uniform(0,1) for _ in range(10000)]
l4 = [random.uniform(0,1) for _ in range(100000)]
l5 = [random.uniform(0,1) for _ in range(1000000)]
l6 = [random.uniform(0,1) for _ in range(5000000)]

l1.sort()
l2.sort()
l3.sort()
l4.sort()
l5.sort()
l6.sort()


def riemann(l):
    """
    Generate a Riemann sum for partition of points l
    Result is multiplied by 4 to approximate the circle area
    """
    length = len(l)
    sum = 0
    for i in range(1, length-1):
        sum += (l[i+1] - l[i])*(sqrt(1 - l[i]*l[i]))
    return 4*sum

t = time.time()

# estimations
e1 = riemann(l1)
e2 = riemann(l2)
e3 = riemann(l3)
e4 = riemann(l4)
e5 = riemann(l5)
e6 = riemann(l6)


print('Results for different inputs using Riemann sums:')
print(f"100:     {e1:.10f} error: {abs(math.pi - e1):.5f}")
print(f"1000:    {e2:.10f} error: {abs(math.pi - e2):.5f}")
print(f"10000:   {e3:.10f} error: {abs(math.pi - e3):.5f}")
print(f"100000:  {e4:.10f} error: {abs(math.pi - e4):.5f}")
print(f"1000000: {e5:.10f} error: {abs(math.pi - e5):.5f}")
print(f"5000000: {e6:.10f} error: {abs(math.pi - e6):.5f}")

print("time elapsed: {:.3f}s".format(time.time() - t))


"""
The Riemann sum approach takes around the same amount of time to approximate the area (excluding the time needed to sort the lists), but starts being an effective approximation much quicker. The error for the Riemann sum method is almost always smaller than 10e-3 for a partition of size 10e4, whereas the error for the geometric probability method is sometimes greater than 10e3 even for a set of points of size 10e6.
"""