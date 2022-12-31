# Advent of Code 2022

Welcome! 🖖

This year, I decided to attempt the [2022 Advent of Code](https://adventofcode.com/2022) contest. I've heard about it from a friend, who basically told me that one problem gets published every day for 25 days and that they're not too difficult. I've never really been into competitive programming, but I was always attracted to problem solving, which I didn't really have that much time for. The idea of the contest was really intriguing to be, and after reading the first two problems I decided to challenge myself to solve all problems during the 25 days, every day after work and university.

This has turned out to be quite a ride. At first, the puzzles were fairly easy, and I just had to get used to writing algorithms. I went with Go as the programming language, since I work as a Go programmer. Once I got in shape, the early days went fast, but the difficulty kept increasing and I even ended up spending hours (even spread across different days) on some problems. There were times around the 16th day when I felt a bit overwhelmed and felt that I was spending too much time on the problems, and was wondering if I was going to be able to complete all of them. But, I was set on completing all of them and kept solving.

In the end, I believe this was very beneficial to me, and even though it costed me some hangouts with friends and Friday nights (I think I spent somewhere around 30 hours to work out all these puzzles), I'm glad I did it. I feel that I have much more control over the fundamental algorithms, and my code writing speed has increased significantly. I was also able to improve my Go skills via this challenge, so every hour spend past 12 a.m. to grind these puzzles was worth it. It brought back the joy of getting a problem right back from the college days when I was doing math contests, and when I finally got the only problem I had remaining right, problem 19, I was so delighted 🥳! I would gladly recommend this contest to anybody, really (although some prior basic algorithms and data structures knowledge would help a ton).

On a related note, I never really played it just for speed. Increasing my coding speed was one of my main goals, but at all times I tried to keep the code clean, readable and have a good separation of concerns. I might have slacked out sometimes when I spent a lot of time on the puzzle because I was happy to just get it done, but I'll likely refactor the code I don't fancy.

Below you can find a _very simplified_ summary of all the contest's problems. The folders contain my own solutions to these problems, and each readme has some notes I found interesting. I did every puzzle myself, without looking for any hints or inspirations (aside looking up some basic techniques like backtracking to refresh the concepts), so the solutions may not be optimal, but I'm happy all of them ran in a good amount of time. I've also scored them by difficulty, because who doesn't love numbers? 😁 

### Enjoy!

> Read the updated version of this article on my [medium story](https://medium.com/@ozoniuss/advent-of-code-2022-contest-overview-and-puzzles-rating-2eb68266a80f).

---

1. [🍩🍫🍪](./01-calorie-counting/): intro

[This problem](https://adventofcode.com/2022/day/1) was pretty easy, the main idea bringing down to computing the first, and first 3 elements of an array. Maybe doing the second part at the same time of reading the input could have a touch of trickiness to it, but there was nothing really outstanding about this one.

I would give this one a difficulty score of one Christmas Tree: 🎄

2. [🧱📄✂️](./02-rock-paper-scissors/): rock, paper... scissors!

The [second day](https://adventofcode.com/2022/day/2) you had to simulate a rock-paper-scissors game. The first part was giving you the exact moves for several games, and you had to compute your score, while in the second part the same input actually represented the outcome of the games. There, you had to find out what you needed to play for each game and then compute the score for all games based on that. 

There wasn't really anything difficult about this one either, perhaps a nice thought exercise would be figuring out the simplest way of modelling the rules of the game, like what beats what. The difficulty score is still one Christmas Tree: 🎄

3. [💼📛📁](./03-rucksack-reogranization/): the common item

The first part of [this puzzle](https://adventofcode.com/2022/day/3) is pretty much taking a string with even length, splitting it exactly in half, and finding the guaranteed unique common letter. There are straightforward ways to do this, by just checking which letter from one half is in the letters of the other half, but this becomes more interesting once you try to "optimize" this (like that's needed for 50-character strings 😉)

Of course, me not having a life I did optimize it, by considering each letter being a bit (a _bit_, yes, not a _byte_; the entire alphabet has 52 big and small characters and that fits in a 7 byte number) and finding the intersection via bitwise operations, `and` to be specific. You can view my madness 😵 [here.](./03-rucksack-reogranization/main.go) 

The only difference in the second part is that you're intersecting three sets instead of two. It's more or less the same idea, you can just check that the current item is in the other two sets as well. With my completely useless optimization, you just do 3 bitwise `and` operations instead of 2.

Perhaps a bit more work if you want to optimize the set intersection, but you can still just get the problem done pretty fast. For this reason, I consider the overall difficulty of this one to still be one Christmas Tree: 🎄

4. [➕♾️➰](./04-camp-cleanup/): sketchy intervals

Time for some more math-y bussiness. With [problem 4](https://adventofcode.com/2022/day/4), you had to do some work with intervals. The first part was figuring out from two intervals if one was part of the other, whereas the second part was figuring out if the intervals intersect. Since the intervals consisted only of consecutive integers, you could just go through all the numbers, but you can also answer both questions just by looking at the interval's bounds. There's slightly more math involved in the bounds approach, but it doesn't take long to get the conditions right.

This was one of the shortest puzzles and it also gets one Christmas Tree in difficulty: 🎄

5. [🚢📦🚧](./05-supply-stacks/): moving boxes

This is the [first](https://adventofcode.com/2022/day/5) out of many simulation problems, and what you simulate is a crane moving boxes from a pile to another out of many piles. It's also the first one where the input wasn't straightforward to parse; in fact, parsing out the boxes was slightly challenging:

```
[D]                     [N] [F]    
[H] [F]             [L] [J] [H]    
[R] [H]             [F] [V] [G] [H]
[Z] [Q]         [Z] [W] [L] [J] [B]
[S] [W] [H]     [B] [H] [D] [C] [M]
[P] [R] [S] [G] [J] [J] [W] [Z] [V]
[W] [B] [V] [F] [G] [T] [T] [T] [P]
[Q] [V] [C] [H] [P] [Q] [Z] [D] [W]
 1   2   3   4   5   6   7   8   9 
```

This is probably the point where I should have started using regex to read the input file (or rather `fmt.Sscanf()`, had I known about it), they made my life much easier down the line once I added them. It would have definitely made reading the moves nicer, too:

```
move 1 from 3 to 9
```

Moving on to the simulation, at every step you move a number of crates from one pile to the other. The parts are almost the same, part 1 just reversing the order of the crates when dropping them at every move. It's also guaranteed that every move is valid, that is, there's always enough crates in the pile you're taking them from.

The simulation itself was not challenging, but storing state does introduce an additional layer of complexity, especially when you're trying to write clean code. This, combined with parsing the input, makes it the first puzzle that receives from me two Christmas Trees in difficulty: 🎄🎄

6. [📱❎🈶](./06-tuning-trouble/): finding the right substring

Back to string work. The input [this day](https://adventofcode.com/2022/day/6) was actually only a long string. The main idea behind this problem was finding the first substring (of consecutive characters) with a given length from the long string such that all the characters in the substring are different. In part (a), that length was 4, and in part (b) that length was 14, probably to add some more algorithmic thinking to those who at part (a) just checked that each letter is different from the other 3 (yes, myself included 🙈).

Either way, this problem was really simple, slightly similar work to problem 3. It's also quite a bit shorter than the previous one, so it only gets one Christmas Tree in difficulty: 🎄

7. [📁💾🖥️](./07-no-space-left-on-device/): cleaning the filesystem

Somewhat of a simulation, [this problem](https://adventofcode.com/2022/day/7) is inspired from the Linux filesystem and in my opinion is really well designed. Your input consists of two Linux commands: `cd` (change directory) and `ls` (list). The problem simulates navigating through the filesystem and displaying the files from various directories, including their size. Both parts can be answered by computing the size of all directories, in particular, finding the ones with total size at most 100,000 at part 1, finding the smallest one in size to delete to get an extra 40,000,000 in space at part 2.

Once you know the size of all directories, finding the answers is not hard at all, but the tricky part here is to actually find the size of each directory! If you know your working directory, you can pretty much interpret the meaning of every line, so this puzzle can be solved by reading the input line by line and holding the current directory and its size. In the input it's also pretty much guaranteed that the exploration pattern is always the same.

Even if the idea itself doesn't seem complicated, the execution is not trivial. All parent directories matter for a directory name, since it wasn't specified that directory names are unique. Also, for every directory size you have to add the size of the child directories. Modelling "being" in a current directory, likely using a stack to allow going back to the parent, also takes a bit of thought, thus this "simulation" was definitely more challenging than the crane one for puzzle 5, leaving the problem with three Christmas Trees in difficulty: 🎄🎄🎄

8. [🌲🏡🚁](./08-treetop-tree-house/): the tallest tree

This is the [first puzzle](https://adventofcode.com/2022/day/8) involving a matrix: a matrix of integers. This is also the first puzzle where I started benefiting from the fact that in Go, arrays are passed by value by default. Whenever I had to store positions, I went with the datatype `[2]int`, and later down the line, I started doing 

```go
type Location [2]int
```

as syntactic sugar.

To break down the requirement, it's easier for me to just use the problem's statement this time. Each integer represents the height of a tree, and at part 1 you had to find all trees visible from a row or column. Part 2 was similar, for each tree you had to compute how many trees it can "see" in all 4 directions, multiply these values together, and find the greatest result of all multiplications.

There likely are better optimizations, but the most intuitive approach of going in all four directions using loops does the trick for this problem. It's a bit more work, but the idea is simple and reading the matrix input is also simple, so I will give this problem one Christmas Tree in difficulty: 🎄

9. [🐍🌉🧵](./09-rope-bridge/): snake

On with the simulations: [this problem](https://adventofcode.com/2022/day/9) was pretty much simulating the snake game (or at least a very similar game), just that the snake here was called "rope". For part 1, the snake had length 2 and length 10 for part 2. The statement describes how the snake "moves", and the input consisted of the snake directions. In both parts, you had to count the number of squares the snake's tail visited.

Modelling the snake as an array of fixed length (in golang) works well here, and defining a `move(snake, direction)` function does a nice separation of concerns. Once the moves are defined on the model, solving part 1 is pretty quick, but applying the same idea for part 2 doesn't work, because the new length can produce different motions that were not present in part 1. Those motions were also not showcased in the example, so you had to find them yourself, and I did spend quite some time wondering why the original approach kept failing. For those reasons, the puzzle receives three Christmas Trees in difficulty: 🎄🎄🎄

10. [🕹️📺🔍](./10-cathode-ray-tube/): hardware instructions with a CRT screen

The simulations keep coming, [now](https://adventofcode.com/2022/day/10) simulating drawing pixels. At part 1, you were simulating addition and subtraction instructions to a register at every cycle, and you had to determine the register's values at some provided cycle numbers. Part 1 is done quickly, but part 2 was definitely a ramp up in difficulty: the cycle number was representing a pixel position on a CRT screen, and based on the register's value (which indicated three consecutive pixel positions) you had to determine whether you were allowed to draw that pixel or not. The fact that after 40 pixels you started a new row also added to the difficulty.

This was also the first and only puzzle that required a human interpretation of the result: at the end you had to draw the "screen", and write the letters displyed on the "screen" as the answer. It definitely was one of the longer problems, but the checks themselves were not difficult, maybe a bit annoying because you had to keep track of multiple things at once. In terms of difficulty, I would consider it on the same page with the previous one, receiving three Christmas Trees: 🎄🎄🎄

11. [🐒⚾🏓](./11-monkey-in-the-middle/): _monke_ games

[Here](https://adventofcode.com/2022/day/11) is where the monkeys (or _monke_-s) got in the picture. In this new simulation problem, at each round, the monkeys have a list of numbers, and each monkey does some operations on every number. Based on the result of the operation, the monkey decides which monkey to send the new number to. The requirement was to compute how many numbers were thrown away by each monkey after a number of rounds.

The difference between part 1 and part 2 were the operations with the numbers. The results of the operations at part 1 were reasonably small numbers, but at part 2 you would very quickly start to get huge numbers that didn't even fit to 64 bits, so you had to come up with a trick to keep the numbers under control. Figuring out why the approach from part 1 is failing and the trick to avoid that is not straightforward and might take some time, thus awarding the [_monke_-s](https://static01.nyt.com/images/2022/11/16/multimedia/16xp-monkeys-01-1-1ea0/16xp-monkeys-01-1-1ea0-superJumbo.jpg) with three Christmas Trees in difficulty: 🎄🎄🎄

12. [🏁🌳⛵](./12-hill-climbing-algorithm/): shortest path

We knew this had to come eventually: [the first backtracking puzzle!](https://adventofcode.com/2022/day/12) This one is quite a classic though: you start from one place in a matrix, and you have to find the shortest path to a different place, with the condition that the neighbours you can go to from some point are based on the letter they have. Part 2 is not that much different, except that there are multiple starting points and you must additionally find the starting point which gives the shortest path.

The idea to solve this problem is pretty well-known, and that is, [bread-first search](https://www.geeksforgeeks.org/shortest-path-unweighted-graph/) (commonly referred to as `BFS`). The idea for part 2 is almost identical, known as [multi-source bfs](https://www.geeksforgeeks.org/multi-source-shortest-path-in-unweighted-graph/). The problem can be done quickly by defining a `getNeighbours(Location) -> Location` function and applying these well-known techniques, but because backtracking is a harder concept to graps in general, the problem receives two Christmas Trees in difficulty: 🎄🎄

13. [🔃↔️🔄](./13-distress-signal/): infinite recursion

[This puzzle](https://adventofcode.com/2022/day/13) introduced perhaps one of the most difficult concepts to define formally, and that is, recursive objects. In particular, lists whose items can be either integers, or lists of integers, even empty lists. The problem's input is a bunch of pairs of this type of lists, and part 1 requires you to "compare" the lists, based on a recursive method defined by the statement. Part 2 is much easier assuming you got the comparison function right; it only requires writing all input lists, together with two additional ones, in the "correct" order based on the comparisons.

There are several tricky parts to this puzzle, that is, defining the recursive object that holds the lists and numbers, and converting between that object and its string representation. Once that's done, implementing the comparison function on the recursive object is less of a challenge since it's clearly defined in the statement, but the string convertion is quite brainy and for this reason, this puzzle is the first one I decided to gift four Christmas Trees in difficulty: 🎄🎄🎄🎄

---

I believe this was the point in the contest where I knew I would be in for quite a ride. Last two puzzles introduced two of the harder concepts in computer science, backtracking and recursion, and problems started to take increasingly longer to solve. At times I was wondering whether it was the right moment for me to do this challenge, since I often ended up working on these puzzles after midnight, but this last problem was simply beautiful and I decided to keep going nonetheless.

---

14. [⏳🕰️⛰️](./14-regolith-reservoir/): falling sand

Another [simulation](https://adventofcode.com/2022/day/14), this time sand falling down a map which also contains rocks in various places. The sand falls from some location and when it encounters a rock (or other sand) it can either stop or fall to the left or right. The puzzle input represents the rock patterns on the map, and one unit for sand falls at each step. Part 1 requires computing the quantity of sand that falls until all new sand units will no longer come to rest and fall down the map (which is guaranteed to happen), where in part 2 a straight rock "border" is placed at the bottom such that sand never falls down, and you have to find when the sand piles up high enough to block the source of sand.

This problem required a fair amount of work: the input is not straightforward to parse, adding the border at part 2 is a bit cumbersome, simulating the sand motion did involve a bit of recursivity and due to the nature of the problem, errors are more difficult to debug (you can draw the map, but you have to compute the region of the map you want to draw first, then actually draw it, and the input map is large...). However, the core idea wasn't really that difficult, especially compared to the previous problem and therefore collecting 3 Christmas Trees in difficulty: 🎄🎄🎄

15. [💎🛑🚧](./15-beacon-exclusion-zone/): map coverage

The [next problem](https://adventofcode.com/2022/day/15) is placed in the 2-dimensional plane. It provides the integer coordinates of a few sensors as input, each one being able to scan all points up to a given distance (using manhattan distance). Part 1 gives a line and asks how many integers points located on that line are detected by these sensors. At part 2, you were given some delimited area where you knew that exactly one point is not scanned by the sensors, and asked to find that point.

With an intuitive approach, one could write a function `isVisible(location, sensor) -> bool` which will get part 1 done, but would be slow for part 2. The catch here is that coordinates are massive: you can compute the manhattan distance instantly but there are multiple sensors and the area at part 2 is 4 million by 4 million, so just checking each individual point is too slow, at least in Golang. Because it did require an optimization, the sensor's coordinates were of the order of millions and computing the bounds of the line at part 2 is slightly trickier (after all you can't check all points on an infinite line), this problem covers 3 Christmas Trees in difficulty: 🎄🎄🎄

Btw, if you're parsing the input with regex, watch out for negative coordinates. That did cost me one hour 😉

16. [🗻🎡🌋](./16-proboscidea-volcanium/): best-scoring road

The contest introduced again a [backtracking problem](https://adventofcode.com/2022/day/16) in day 16. This time there were a bunch of valves and tunnels between the valves (so basically a graph with edges of length 1) and each opened valve would release a certain amount of pressure each minute (possibly 0). At every minute, you could either do nothing, open a valve or move one edge to a different valve, the requirement being to maximize the released pressure after 30 minutes. Part 2 was pretty much the same, except that you had 26 minutes and a ["companion"](https://a-z-animals.com/media/2022/05/elephant.jpg) to open valves with you, that could do the exact same operations.

Unlike week 12, this was a less straight-forward approach of backtracking with DFS (or BFS). One had to construct the graph from an input that wasn't trivial to parse, there were some challenges to model the steps nicely because it's actually harder to store the state minute by minute, and coming with an approach for part 2 does require some thought. The implementation definitely takes some time, and combined with the difficulty made the problem worthy of 4 Christmas Trees: 🎄🎄🎄🎄

---

On a more personal note, this is the first problem that discouraged me because I wasn't able to finish it the day I started it, and neither the next problem which was even harder. I actually completed it on day 21 I think, when I solved 3 problems. I kept doing the really stupid mistake of adding the starting point twice to the input graph at part 2, which kept messing up my final answer. When I realised what the issue was, my reaction was best described by 😒

This is also the first problem I optimized with parallelism, which decreased the time to solve part 2 substantially. I will definitely be looking to add more parallel programming when remastering these problems.

---

17. [🧱🔷🟩](./17-pyroclastic-flow/): some kind of Tetris

Around the corner was sitting likely the [puzzle](https://adventofcode.com/2022/day/17) that I liked the most. In this next simulation, you had to simulate a game of Tetris where complete rows don't vanish, and you were given 5 shapes and a sequence of left or right keyboard inputs. Both the sequence of shapes and keyboard inputs was repeated once it ended, and after each keyboard input the piece dropped exactly one unit. Parts 1 and 2 were the exact same task: to compute the height of the tower after 2022, and 1 trillion rocks fell, respectively.

Yeah, 1 trillion -- you heard that right. Simulating the entire game was possible at part 1 and the approach I went with, but that just wasn't going to cut it for part 2. The only hope there was finding some patterns. And finding a pattern that is proven mathematically to work is very difficult: what is the condition for the pattern to repeat? How do you even find the pattern? There were also 5 different pieces, and 10091 keyboard inputs (in my puzzle input at least), suggesting a pattern doesn't occur quickly. Not only that, but even once you found the pattern, the math required to compute the height using the pattern is just nuts.

After solving part 1 fairly quickly, the difficulty increase brought by part 2 was baffling to me. Adding everything together: modelling the input and states, finding the pattern (in some games you can prove there is _no cyclic pattern_, although not the case here) and computing the height based on the pattern, in my opinion completely justifies the maximum number of Christmas Trees I'm giving this problem in difficulty: 🎄🎄🎄🎄🎄

18. [🧊🔥🌪️](./18-boiling-boulders/) structure of mini-cubes

Breaking the monotony of the 2-dimensional space is, in my opinion, [one](https://adventofcode.com/2022/day/18) of the most elegant puzzles of the contest. This one pretty much consisted of placing unit-cubes in the 3-dimensional space. The input was giving the positions of a bunch of `1x1x1` cubes, and at part 1 you had to compute how many faces of these cubes are not touching another cube. Part 2 did make matters more interesting, and the easiest way to visualize the requirement for me is to think that the generated cube structure is plunged into water, and you had to count the number of faces that get wet.

While I managed to get part 1 done really quickly, part 2 brought additional layers of complexity, and did require using your spatial orientation a lot when designing an algorithm to solve it. It the end, still an exploration problem solved by BFS, but understanding what you had to do, visualizing the "hidden" spots not reached by water and coming up with a proven working algorithm is definitely challenging. Generally it is much harder to work with points in space than in the plane, but I'd say the problem was overall a bit easier than the previous one, making 4 Christmas Trees a fair difficulty score: 🎄🎄🎄🎄

19. [🤖⛏️⛑️](./19-not-enough-minerals/) and yet another backtracking problem

I'll just start by saying that [this problem](https://adventofcode.com/2022/day/19) did make me want to rip the hair off my head. It is the last puzzle I submitted, which got me the 50th star. You had to make the best investments: you were given 4 different robots, each being able to gather one type of resource (ore, clay, obsidian, geode) every minute. You were also given multiple blueprints with the cost of each robot, and you had to find out the best strategy to buy robots when you had the resources in order to maximize the number of geodes after 24 minutes, for each blueprint. Part 2 is essentially the same as part 1, except you only had to compute the maximum number of geodes for part of the input, but after 32 minutes.

The technique that solves this one is backtracking with DFS (or BFS, but that takes a huge toll on memory) like we've seen before, but the catch is that is suffers terribly from the state-space explosion problem without some good pruning. And finding that good pruning is hard. Even after hashing each possible state and trying to estimate the maximum number of geodes from every state, it still took a good 10 seconds to run part 2, and it would have probably taken days without any pruning at all. I might also be a bit biased here due to solving it last, but because it was challenging to model the states (which for me included the number of robots, number of states and some action) and to find some good pruning then implement it, I consider this problem alongside the "Tetris" one from day 17 to be the only two ones worth 5 Christmas Trees in difficulty: 🎄🎄🎄🎄🎄

Just to strengthen why I think modelling this problem is so difficult, I only figured out on day 25 that my original approach was incorrect and allowed creating robots one round before they were actually created, so I had to come up with a way that signaled if robots are built during the current state. After fixing this, the state explosion moved from minute 14 to minute 20, which was a significant improvement. I did however also find the statement itself to not reveal that detail clearly.

20. [1️⃣♻️#️⃣](./20-grove-positioning-system/) cyclic list

The waters chill down a bit on [day 20](https://adventofcode.com/2022/day/20), after what I think were 3 of the most difficult problems of the contest, one after the other. The statement had much more simplicity, the input being just the elements of a list of integers. The list is considered to be cyclic, as if the elements were in a circle, and for part 1 you had to go through all the elements from the input, and for each one move it on the circle a number of positions equal to its value. So if the element was 3, you'd move it past the 3 next elements, if it was -2 you'd move it before the previous two elements. Part 2 was the same, except that the numbers were all multiplied by some large number like `811589153`, and you had to repeat the process by going through the original list of numbers 10 times. The task required to find the numbers on three positions.

Although this problem is nowhere near as difficult as the previous one, it still has its quirks. There's deciding on how to store the elements; if using a normal array, moving the elements has some edge cases, otherwise one could make use of a linked list. Also, one should implement an optimization to avoid unnecessary round trips around the circle of numbers. And keep in mind that there are also duplicates, meaning that you have to know which number to move. Figuring all of these out increased the time I spent on this problem way past what I thought initially it would take me, for which reason I grant the problem three Christmas Trees in difficulty: 🎄🎄🎄

---

Fun fact: I originally got the answer wrong by using an array implementation, so I switched completely to using a linked list. I still got the same answer, though. It turned out that I changed my implementation completely just to find out that I had copied the answer from Microsoft's calculator with a comma: `27,726` instead of `27726` 🙄. Nevertheless, I much prefer the linked list approach, but I would have rather done something else than spending one more hour on this at 2 a.m.

---

21. [〰️🔱🌿](./21-monkey-math/) math with recursion and binary search trees

Coding the solution to [this puzzle](https://adventofcode.com/2022/day/21) was just nice. You were given a bunch of math expressions: on the left side, some variable. On the right side, either a math expression (`+`,`-`,`/`,`*`) with two different variables or some value. So basically you either knew the value of the variable, or had to compute it by going through the sub-expressions of the variables in the original variable's expression, recursively. This leads nicely to a binary tree of expressions. For part 1, you had to find the expression of the `root` variable, which was the top of the tree.

While part 1 was a simple recursion, part 2 was much more challenging. There was a variable called `humn` that had a value initially, but this time you had to specify which value it needs to have such that the two parts of the `root` expression have equal values. This did require coming up with a clever way to parse the tree to get an expression for `humn`, which was a leaf in the tree. Seeing the approach is not straightforward, and since recursivity and binary trees are hard concepts, my expressions determined that four Christmas Trees is the value for the difficulty variable: 🎄🎄🎄🎄

22. [🦧🗺️🎪](./22-monkey-map/): expanded cube

Incoming simulation alert: [this time](https://adventofcode.com/2022/day/22), an oddly shaped map and a set of moves and direction changes as input. Basically, you knew how many places to go forward before changing the direction to either left or right. For part 1, when you fell off the map, you spawned in the opposite location on the same line or column. There were also rocks that could obstruct your movement, even at the other end, when you tried to respawn. Part 1 required to compute the final location after the moves were completed.

The contest did make us familiar with simulations, and thus part 1 was nothing out of the ordinary. At part 2 though, it turned out that the map is in fact the expanded version of a cube, and on the two-dimensional expansion you had to simulate moving on the cube's faces. This significantly ramped up the difficulty, because mapping the edges to one another was a pain, and the behaviour now depended on direction too on the map's corners. Additionally, if you interpreted the input as a list with the direction in which you move every step, that would get messed up at part 2 because wrapping around the map could change direction (which cost me one hour...). This was amongst the harder simulations, but mostly a lot of work and not necessarily a difficult idea. The amount of work was still substantial, paving the path to 4 Christmas Trees in difficulty for the puzzle: 🎄🎄🎄🎄

23. [🗺️🏔️🌱](./23-unstable-diffusion/) another map simulation??

What, tired of simulations? You should have heard today's severe simulation warning⚠️. However, this time an easier [one](https://adventofcode.com/2022/day/23). Your input is a bunch of elves' locations on an infinite map. At each step, you know a list of directions where each elf can move, and the elves can't move in a direction if they are "too close" to neighbouring elves, or completely isolated. At part 1, you had to enclose the elves in a square with smallest area after simulating 10 moves, and had to count the empty squares. And at part 2, you simply had to run the simulation until all elves were isolated from their neighbours, obstructing all movement, and answer with the number of rounds it took to reach there.

Overall, the movements and move conditions were pretty well explained and easy to model, but there were trickier parts, in particular, dealing with an infinite grid and addressing move conflicts. This type of problem is also not easy to debug, especially if the map is large, therefore receiving a total number of three Christmas Trees in difficulty: 🎄🎄🎄

24. [➡️🚸⛈️](./24-blizzard-basin/) path finding on changing map

Last backtracking [puzzle](https://adventofcode.com/2022/day/24). Well, at least it was a challenging one. The input is a bounded map in the shape of a rectangle, a starting point and a finish point. On the map, there are also a big number of winds called blizzards going in one of four directions, passing through each other and wrapping around the map when they reach the bounds. You just have to find the shortest path from start to finish, with the constraint that you must never be on the same tile with a blizzard. Part 2 is basically the same task, repeated three times: you need to find the shortest path from start to finish, then finish to start, and then back again start to finish (there will be different paths because blizzards will be in different positions).

Of course, the main technique is still backtracking with BFS, but there's the additional difficulty of keeping track of the wind positions at each step, actually moving the winds, and remembering to move if a wind takes your place (it is guaranteed that you have a valid place to move). Moreover, you will need some pruning techniques, otherwise your algorithm will run in hours (at least in Go). Thankfully, exploring (step, location) pairs reached via different paths only once is a good enough pruning which allows BFS to complete in a few minutes, but the difficulty is still there and there was quite a lot of work to simulate those winds, making this problem one step above the previous simulation in difficulty at four Christmas Trees: 🎄🎄🎄🎄

25. [❓➗5️⃣](./25-full-of-hot-air/) a different form of base 5 numbers

After a bumpy ride, the adventure finally comes to an end with, luckily for me, a [math puzzle](https://adventofcode.com/2022/day/25). The problem comes up with a different counting system to write numbers, called SNAFU, the possible digits being `2`,`1`,`0`,`-`,`=` with values `2`,`1`,`0`,`-1`,`-2` respectively. To compute the decimal value of a number written in SNAFU, simply multiply each number with 5 to the power of its position and add those together, for example, `2=1` is `2*5^2 - 2*5^1 + 1*5^0 = 50 - 10 + 1 = 41`. The problem had only one part: you were given a list of numbers using this representation, had to add them up and write the sum as SNAFU.

Luckily, there were no negative numbers, and all you really had to do was follow the same process of converting to and from base 5, in a slightly different way. Converting from SNAFU to decimal is identical, and once you replace the possible remainder values with the set `{-2,1,0,1,2}`, you can re-define the quotient and remainder modulo 5 as `q = (x+2)/5` and `q = (x+2)%5-2`. Then, the standard technique of converting to base 5 can be used for this new counting system. I likely found this really simple because of being a mathematician in the past, but I still think the conversions could be a bit tricky to get right, and one also needs to associate with the symbols of the new counting system, making the end of this contest marked by a nice two Christmas Tree difficult problem: 🎄🎄

---

Note: thinking about how to do the conversions for negative numbers gets interesting, because the "zero" values don't align: a negative number can start with 1 in the new counting system.

---

That was the end of the 2022 Advent of Code contest. I had a tremendous amount of fun, and for those of you who participated, I hope you also enjoyed the journey! Overall, the puzzles and story were really well-thought, and they definitely help you improve your coding skills.

In the end, here are all the problems ordered by difficulty, in no particular order for those with the same difficulty.

- 🎄🎄🎄🎄🎄: days 17, 19

- 🎄🎄🎄🎄: days 13, 16, 18, 21, 22, 24

- 🎄🎄🎄: days 7, 9, 10, 11, 14, 15, 20, 23

- 🎄🎄: days 5, 12, 25

- 🎄: days 1, 2, 3, 4, 6, 8