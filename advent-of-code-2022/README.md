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

to be continued...
