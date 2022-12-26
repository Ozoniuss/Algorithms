# Advent of Code 2022

Welcome! 🖖

This year, I decided to attempt the [2022 Advent of Code](https://adventofcode.com/2022) contest. I've heard about it from a friend, who basically told me that one problem gets published every day for 25 days and that they're not too difficult. I've never really been into competitive programming, but I was always attracted to problem solving, which I didn't really have that much time for. The idea of the contest was really intriguing to be, and after reading the first two problems I decided to challenge myself to solve all problems during the 25 days, every day after work and university.

This has turned out to be quite a ride. At first, the problems were fairly easy, and I just had to get used to writing the algorithms. I went with Go as the programming language, since I'm a Go programmer. Once I got in shape by solving the first two or three problems, it took me 20-30 minutes to solve the easier ones, but the difficulty kept increasing, and I even ended up spending hours (sometimes even in different days) on some problems. I must say, sometimes after the 16th day, not exactly sure which one, there were times when I felt it was a bit overwhelming and that I was spending too much time on the problems, and was wondering if I was going to be able to complete all of them. 

In the end, I believe this was very beneficial to me, and even though it cost me some hangouts with friends and Friday nights (I think somewhere around 30 hours of work to solve all these problems), I'm glad I did it. I feel that I have much more control over the fundamental algorithms, and my code writing speed has increased significantly. I was also able to improve my Go skills via this challenge, so every hour spend past 12 a.m. to grind these problems was worth it. It brought back the joy of getting a problem right back from the old college days when I was doing math contests, and I have to admit that the last day when I finally got the only problem I had remaining right, problem 19, I was so delighted 🥳! I would gladly recommend this contest to anybody, really (although some prior basic algorithms and data structures knowledge would be an enormous benefit).

On a related note, I never really played it just for speed. Increasing my coding speed was one of my main goals, but at all times I tried to keep the code clean, readable and have a good separation of concerns. I might have slacked out sometimes when I spent a lot of time on the problem because I was happy to just get it done, but I'll likely refactor the code I don't fancy. There will be a link to the original solutions anyway for each problem.
 
Below you can find a _very simplified_ summary of all the contest's problems. The folders contain my own solutions to these problems, and each readme has some notes I found interesting. I did every problem myself, without looking for any hints or inspirations (aside looking up some basic techniques like backtracking to refresh the concepts), so the solutions may not be optimal, but I'm happy all of them ran in a good amount of time.

### Enjoy!

---

1. [🍩🍫🍪](./01-calorie-counting/): intro

This problem was pretty easy, the main idea bringing down to computing the first, and first 3 elements of an array. Maybe doing the second part at the same time of reading the input could have a touch of trickiness to it, but there was nothing really outstanding about this one.

I would give this one a difficulty score of one Christmas Tree: 🎄

2. [🧱📄✂️](./02-rock-paper-scissors/): rock, paper... scissors!

Like the title says, here you had to simulate a rock-paper-scissors game. The first part was giving you the exact moves for several games and you had to compute your score, while in the second part the same input actually represented the outcome of the games; you had to find out what you needed to play for each game and then compute the score for all games based on that. There wasn't really anything difficult about this one either, perhaps a nice thought exercise would be figuring out the simplest way of modelling the rules of the game, like what beats what.

The difficulty score is still one Christmas Tree: 🎄

3. [💼📛📁](./03-rucksack-reogranization/): the common item

The first part of this problem is pretty much taking a string with even length, splitting it exactly in half, and finding the guaranteed unique common letter. There are straightforward ways to do this, by just checking which letter from one half is in the letters of the other half, but this becomes more interesting once you try to "optimize" this (as if you need to for 50-character strings 😉)

Of course, me being a huge nerd I did optimize it, by considering each letter being a bit (a _bit_, yes, not a _byte_; the entire alphabet has 52 characters and that fits in a 7 byte number) and finding the intersection via bitwise operations, `and` to be specific. You can view my madness 😵 [here.](./03-rucksack-reogranization/main.go) 

The only difference in the second part is that you're intersecting three sets instead of two. It's more or less the same idea, you just check that the current item is in the other two sets as well. With my completely useless optimization, you just do 3 bitwise `and` operations instead of 2.

Perhaps this was not as straightforward as the other two, especially if you wanted to optimize a bit the set intersection, but you can still just get the problem done pretty fast. For this reason, I consider the overall difficulty of this one to still be one Christmas Tree: 🎄

4. [➕♾️➰](./04-camp-cleanup/): sketchy intervals

Time for some more math-y bussiness. With problem 4, you had to do some work with intervals. The first part was figuring out from two intervals if one was part of the other, whereas the second part was figuring out if the intervals intersect. Since the intervals consisted only of consecutive integers, you could just go through all the numbers, but you can also answer both questions just by looking at the interval's bounds. Perhaps figuring out the conditions with the bounds instantly is a bit trickier, but it doesn't take long to get them right.

For this reason, this was one of the shortest problems and it also gets one Christmas Tree: 🎄

5. [🚢📦🚧](./05-supply-stacks/): moving boxes

This is the first (if you don't count rock, paper, scissors) out of many simulation problems, and it simulates a crane moving boxes from a pile to another out of many piles. It's also the first one where the input wasn't straightforward to parse; in fact, parsing out the boxes was slightly challenging:

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

This is probably the point where I should have started using regex to read the input file, they made my life much easier down the line once I added them. It would have definitely made reading the moves nicer too:

```
move 1 from 3 to 9
```

Moving on to the simulation, at every step you move a number of crates from one pile to the other. Part (b) is the exact same, except that it reverses the order of the crates at every move. It's also guaranteed that every move is valid, that is, there's always enough crates in the pile you're taking them from.

The simulation itself was not challenging, but storing state does introduce an additional layer of complexity, especially when you're trying to write clean code. This, combined with parsing the input makes this the first problem that receives from me two Christmas Trees in difficulty: 🎄🎄

6. [📱❎🈶](./06-tuning-trouble/): finding the right substring

Back to string work. The input was actually only a long string this time. The main idea behind this problem was finding the first substring (with consecutive characters) with a given length from the long string such that all the characters in the substring are different. In part (a), that lenght was 4, and in part (b) that length was 14, probably to add some more algorithmic thinking to those who at part (a) just checked that each letter is different from the other 3 (yes, myself included 🙈).

Either way, this problem was really simple, slightly similar underlying idea to problem 3 where you had to check for a common character. The most straightforward ideas would be converting the substring to a set, or checking char by char in two `for` loops. Doing that in a function while advancing with the substring one letter at a time through the main string also makes for a nice separation of concerns. Technically you could also optimize by keeping track of the characters in the previous check, but that's not really necessary.

```
dcbc(sbblhhgdgssmcm)qccdw  -->  dcbcs(bblhhgdgssmcmq)ccdw  --> ...
```

This problem was quite a bit shorter than the previous one, so it only gets one Christmas Tree in difficulty: 🎄

7. [📁💾🖥️](./07-no-space-left-on-device/): cleaning the filesystem

Somewhat of a simulation as well, this problem is inspired from the Linux filesystem and in my opinion is really well thought-out. Your input consists of two Linux commands: `cd` (change directory) and `ls` (list). The problem simulates navigating through the filesystem and displaying the files from various directories, including their size. In both parts, the requirement was pretty much computing the size of each directory, and identifying the one(s) that satisfies a specific constraint, in particular, having a size greater than some number at part (a), and being the smallest one with the size greater than some number at part (b).

Once you know the size of all directories, answering the questions is not hard at all, but the tricky part here is to actually find the size of each directory! A stack is your best friend to keep track of the current directory. If you know your working directory, you can pretty much interpret the meaning of every line, so this problem can be solved by reading the input line by line and holding the current directory and its size. In the input it's also pretty much guaranteed that the exploration pattern is always the same.

Even if the idea itself doesn't seem complicated, the execution is not trivial. One must take into account all parent directories for a directory name, since I didn't find the statement to specify whether directory names are unique, and must also keep in mind that to the size of a directory one must add the size of the child directories. With all these in mind, this "simulation" was definitely more challenging than the crane one for problem 5, giving the problem three Christmas Trees in difficulty: 🎄🎄🎄

8. [🌲🏡🚁](./08-treetop-tree-house/): the tallest tree

This is the first problem involving a matrix: a matrix of integers. This is also the first problem where I started benefiting from the fact that in Go, arrays are passed by value by default. Whenever I had to store positions, I went with the datatype `[2]int`, and later down the line, I started doing 

```go
type Location [2]int
```

as syntactic sugar.

To break down the requirement, at part (a) you had to find all numbers in the matrix that's either greater than all numbers in front or behind it, either on its row or column. Part (b) was similar, for all numbers in the matrix you had to compute a property based on how many numbers on its row or column are greater than itself, and find the position with the highest value for that property.

There likely are better optimizations, but the most intuitive approach of going in all four directions using loops does the trick for this problem. And since that and reading a matrix input is straightforward, I will give this problem one Christmas Tree in difficulty: 🎄

9. [🐍🌉🧵](./09-rope-bridge/): snake

On with the simulations: this problem was pretty much simulating the snake game (or at least a very similar game). For part (a), the snake had length 2 and length 10 at part (b). The statement describes how the snake "moves", and the input consisted of the snake directions. In both parts, you had to say which squares the snake visited.

Using the same trick as before for positions and modelling the snake as an array of fixed length (in golang) works well here, and defining a `move(snake, direction)` function does a nice separation of concerns. Once the modelling is done, part (a) is done pretty quickly, but applying the same idea for part (b) doesn't work, because the motion can place the snake in certain positions that don't come to mind initially, where the movement is different. These were also not showcased in the example, so I did spend quite some time wondering why the approach at part (a) kept failing. I'll describe that case in more detail on the problem page, but for that reason, this problem receives three Christmas Trees in difficulty: 🎄🎄🎄

10. [🕹️📺🔍](./10-cathode-ray-tube/): hardware instructions with a CRT screen

The simulations keep going, now simulating cycles of a circuit and drawing pixels. At part (a), the instructions were simulating adding and subtracting values to a register at every cycle, and you had to keep track of the register's values at some specific cycle numbers. Part (a) went quickly, but part (b) was definitely a ramp up in difficulty: the cycle number was simulating a pixel position on a CRT screen, and based on the register's value (which indicated three consecutive pixel positions) you had to determine whether you were allowed to draw a pixel or not. The fact that after 40 pixels you started a new row also added to the difficulty.

This was also the first and only problem that required a human interpretation of the result: at the end you had to draw the "screen", and write the letters displyed on the "screen" as the answer. It definitely was one of the longer problems, but the checks themselves were not difficult, maybe a bit annoying because you had to keep track of multiple things at once. Talking about difficulty, I would consider it on the same page with the previous one, receiving three Christmas Trees: 🎄🎄🎄

11. [🐒⚾🏓](./11-monkey-in-the-middle/): _monke_ games

This is where the monkeys (or _monke_-s) got in the picutre. In this new simulation problem, at each round, the monkeys have a list of numbers, and each monkey does some operations to every number. Based on the result of the operation, the monkey decides which monkey to send the new number to. The requirement was to compute how many numbers were thrown away by each monkey after a number of rounds.

The difference between part (a) and part (b) were the operations with the numbers. The results of the operations at part (a) were reasonably small numbers, but at part (b) you would very easily start to get huge numbers that didn't even fit to 64 bits fairly quickly, so you had to come up with a trick to keep the numbers under control. Figuring out why the approach from part (a) is failing and the trick to avoid that is not straightforward and might take some time, thus awarding the _monke_-s with three Christmas Trees in difficulty: 🎄🎄🎄

12. [🏁🌳⛵](./12-hill-climbing-algorithm/): shortest path

We knew this had to come eventually: the first backtracking problem! This one is quite a classic though: you start from one place in a matrix and you have to find the shortest path to a different place, with the condition that the neighbours you can go to from some point are based on the letter they have. Part (b) is not that much different, except that there are multiple starting points and you must additionaly find the starting point which gives the shortest path.

The idea to solve this problem is pretty well-known, and that is, [bread-first search](https://www.geeksforgeeks.org/shortest-path-unweighted-gr) (commonly referred to as `BFS`). The idea for part (b) is almost identical, known as [multi-source bfs](https://www.geeksforgeeks.org/multi-source-shortest-path-in-unweighted-graph/). This problem can be done quickly by defining a `getNeighbours(Location) -> Location` and applying these well-known techniques, but because backtracking is a harder concept to graps in general, the problem receives two Christmas Trees in difficulty: 🎄🎄

13. [🔃↔️🔄](./13-distress-signal/): infinite recursion

This problem introduced perhaps one of the most difficult concepts to define formally, and that is, recursive objects. In particular, lists whose items can be either integers, or lists of integers, even emtpy lists. The problem's input is a bunch of pairs of these types of lists, and part (a) requires you to "compare" the lists, based on a recursive method defined by the statement. Part (b) is much easier assuming you got part (a) right; it only requires writing all input lists, together with two additional ones, in the "correct" order based on the comparison method.

There are several tricky parts to this method, that is, defining the recursive object that holds the lists and numbers, and converting between that object and its string representation. Once that's done, implementing the comparison function on the recursive object is less of a challenge since it's clearly defined in the statement, but the string convertion is quite brainy and for this reason, this problem is the first one I decided to give four Christmas Trees in difficulty: 🎄🎄🎄🎄

---

I believe this was the point in the contest where I knew I would be in for quite a ride. Last two problems introduced two of the harder concepts in computer science, backtracking and recursion, and problems started to take increasingly longer to solve. At times I was wondering whether it was the right moment for me to do this challenge, since I often ended up working on these problems after midnight, but this last problem was simply beautiful and I decided to keep going nonetheless.

---

14. [⏳🕰️⛰️](./14-regolith-reservoir/): falling sand

The simulations keep coming, this time sand falling down a map which also contains rocks in various places. The sand falls from some location and when it encounters a rock (or other sand) it can either stop or fall to the left or right. The puzzle input represents the rock patterns on the map, and one unit for sand falls at each step. Part (a) requires to compute the quantity of sand that falls until all new sand units will no longer come to rest and fall down the map (which is guaranteed to happen), where in part (b) a straight rock "border" is placed at the bottom such that sand never falls down, and you have to find when the sand source gets blocked by the pile of sand.

This problem required a fair amount of work: the input is not straightforward to parse, adding the border at part (b) is a bit cumbersome, simulating the sand motion did require a bit of recursivity and due to the nature of the problem, erros are more difficult to debug (you can draw the map, but you have to compute the region of the map you want to draw first, and then actually draw it...). However, the core idea wasn't really that difficult, especially compared to the previous problem and therefore collecting 3 Christmas Trees in difficulty: 🎄🎄🎄