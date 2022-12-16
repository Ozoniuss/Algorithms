# Matrix ZigZag

Given a `3 x 3` matrix, print the numbers in the following way:

- first, print all numbers on the top line
- secondly, print the numbers on the secondary diagonal
- lastly, print all numbers on the last row

You can imagine this like drawing a `Z` over the matrix.

You can use any data structure to store the matrix, and you can populate it in any way you want (variable, read from file, string, etc.). Ideally, the algorithm should generalize to a `n x n` matrix.

### Example

Input:

```
[1, 2, 3]
[4, 5, 6]
[7, 8, 9]
```

Output:

```
1,2,3,5,7,8,9
```