# Rectangle Around Points

Given a list of lattice points (represented by the x and y integer coordinates), find the coordinates of the rectangle with the smallest area that contains all of these points.

It doesn't matter how the points are stored (list, map) or how they are populated, e.g. a hardcoded list is also fine. The algorithm should generalize to a set of points of any length, including points that are repeated.

### Example

Input:

```
[0,2]; [1,3]; [2,0]; [2,3]; [2,5]; [3,2]; [3,4]; [6,2]; [6,4]; [8,5]; [9,3]; [9,4]
```

The rectangle that contains these points looks like this (the point at the lower
left cornet is at the origin (0,0)):

```
. . x . . . . . x .
. . . x . . x . . x
. x x . . . . . . x
x . . x . . x . . .
. . . . . . . . . .
. . x . . . . . . .
```

Output:

```
[0,0]; [0,5]; [9,0]; [9,5]
```