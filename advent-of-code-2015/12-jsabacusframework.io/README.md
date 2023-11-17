Yet another problem I finished at 2 a.m.

I just want to quickly highlight the difference between the two approaches. In part (a), it was possible to read the entire file byte by byte and basically figure out what the numbers are. That was easy enough.

But part (b) was significantly more complicated to complete with this approach. You'd have to know the number is within an object which has another key whose value is "red", and within deeply nested fuckaries that's just difficult. So I had to resort to using a json marshaler to unpack the data into a Go object modelling a restricted JSON format (map[string]any) which was good enough for this case luckily, and traversing the nodes recursively using a depth-first search approach. It was quite easy in this way to figure out what needs to be ignored when traversing maps.

Nevertheless, this second approach can obviously also be used to solve part (a), with a more simplified solution even. I wanted to benchmark the difference in performance between the two approaches. And obviously the first one was faster (that is, of course, if you know how to optimize file parsing, string operations and string conversions).

```sh
$ go test -run=xxx -bench=. -benchmem *.go
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkPart1-16           7610            145605 ns/op           16178 B/op       1511 allocs/op
BenchmarkPart2-16           1741            711207 ns/op          359136 B/op       8285 allocs/op
PASS
ok      command-line-arguments  2.840s
```
