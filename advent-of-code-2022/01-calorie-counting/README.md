## [Day 1: Calorie Counting](https://adventofcode.com/2022/day/1)

This was the first problem of the 2022 advent of code contest, so I didn't really know what to expect. For a first problem, this was great at setting the tone of the contest and making me understand how the problems' format was going to be like. The problem itself was really easy, but it put a start to setting some habits which helped in developing speed, like knowing how to parse my input, which data structures to use, how to organize my code in functions and so on. Hell, I only realized now that it was probably one of my first times reading input from a file in go 😅 (at my job, I never really read a file using go). I settled on an approach that I went with for all problems:

```go
f, err := os.Open("input.txt")
if err != nil {
    panic(err)
}
defer f.Close()

scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line = scanner.Text()
}
```

You normally don't panic in production code or pass the file name directly, but whatever 😒

Now, about the problem. For part (a), the approach was quite straightforward, you had a bunch of lists of numbers separated by space, and you had to find the list with the biggest sum and compute that sum. And in part (b), you needed the three biggest sums. You can do this at the same time as reading the input file by keeping track of the maximum value(s) you read up until that point, so the problem was done fairly quickly (although the logic at part (b) did twist my mind a bit first).

Computing the maximum values at the same time as reading the input was my first solution, but is probably not the best separation of concerns since you're doing several things in the function that reads the input, and the memory optimization is not really worth it. What I would instead do now (I might have done it already if I refactored) would be

- When reading the input, compute the sum for each elf and add that sum to a vector;
- Find the top and top 3 elements from the vector of sums.

The main idea required to solve this problem is then isolated from parsing the input to just finding the top elements of an array.

Find the original solution [here.](https://github.com/Ozoniuss/Algorithms/commit/6cbdd80b44704428e2b01903a932e033295eb2a3)