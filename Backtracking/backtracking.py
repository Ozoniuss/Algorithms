from typing import List
from math import factorial
import time

# original solution
def cartezianProductItervative(sets: List[List[int]]):
    product = [[]]
    for set in sets:
        new_lists = []
        for element in set:
            for each_list in product:
                new_lists.append(each_list+[element])
        product = new_lists
        print(product)

    return product

# original solution
def cartezianProductRecursive(sets: List[List[int]]):

    if len(sets) == 0:
        return [[]]

    result = []
    first_set = sets.pop(0)
    other_sets = cartezianProductRecursive(sets)
    for element in first_set:
        for lists in other_sets:
            result.append([element] + lists)

    return result

# stack overflow solution
def cartezian_iterativ(array):
  results = [[]]
  for i in range(len(array)):
    temp = []
    for res in results:
      for element in array[i]:
        temp.append(res+[element])
    results = temp

  return results


# some_lists = [[1,2,3], [4,5,6], [7,8,9]]
# product = cartezianProductRecursive(some_lists)
# for set in product:
#     print(set)

def permutations_bfs(word):
    final_words = ['']
    d = {}
    for x in word:
        d[x] = False
    for i in range(len(word)):
        temp = []
        for fw in final_words:
            for l in fw:
                d[l] = True
            for letter in word:
                if len(fw) == 0:
                    temp.append(fw+letter)
                elif d[letter] == False:
                    temp.append(fw+letter)
            for x in word:
                d[x] = False
        final_words = temp
    return final_words

def permutationsOfAWordIterative(word:str):
    permutations = ['']
    for i in range(len(word)):
        new_permutations = []
        for permutation in permutations:
            for add_letter in word:
                if add_letter not in permutation: # this search could compromise efficiency
                    new_permutations.append(permutation + add_letter)

        permutations = new_permutations

    return permutations

def permutationsOfAWordIterativeWithHash(word:str):
    permutations = ['']
    for i in range(len(word)):
        new_permutations = []
        for permutation in permutations:
            for add_letter in word:
                if add_letter not in set(permutation): # this search could compromise efficiency
                    new_permutations.append(permutation + add_letter)

        permutations = new_permutations

    return permutations


t1 = time.time()
i = 1
for word in permutationsOfAWordIterativeWithHash('abcdefghi'):
    #print(str(i) + word)
    i += 1

t1_final = time.time() - t1

t2 = time.time()
i = 1
for word in permutationsOfAWordIterative('abcdefghi'):
    #print(str(i) + word)
    i += 1

t2_final = time.time() - t2

t3 = time.time()
i = 1
for word in permutations_bfs('abcdefghi'):
    #print(str(i) + word)
    i += 1

t3_final = time.time() - t3

print("time elapsed hash: {:.3f}s".format(t1_final))
print("time elapsed nohash: {:.3f}s".format(t2_final))
print("time elapsed dictbfs: {:.3f}s".format(t3_final))





#
# i = 1
# for word in permutations_bfs('abcd'):
#     print(str(i) + word)
#     i += 1


#Let n and a1, a2, ... given. Decompose n as sum of a1, .., an
def bfs_decomp(numbers, x):
    final_list = []
    queue = []
    for n in numbers:
        queue.append([n, n])
    while(len(queue) > 0):
        first = queue[0]
        queue = queue[1:]
        if(first[-1] == x):
            final_list.append(first)
        else:
            max_number = first[-2]
            for number in numbers:
                if number <= max_number:
                    new_el = first.copy()
                    sum = first[-1]
                    sum = sum + number
                    if(sum <= x):
                        new_el[-1] = number
                        new_el.append(sum)
                        queue.append(new_el)
    return final_list

# print(bfs_decomp([1,1,1,2,3], 5))

#returns all the subsets of the arrat [1,2,....,n]
def subsets_of_1_n_iter(n):
    subsets = []
    queue = []
    for a in range(1, n+1):
        queue.append([a])
    while(len(queue) > 0):
        first = queue[0]
        queue = queue[1:]
        subsets.append(first)
        for i in range(first[-1] + 1, n+1):
            copy = first.copy()
            copy.append(i)
            queue.append(copy)
    return subsets

# print(subsets_of_1_n_iter(3))

def binomial_coefficient(n, k):   #computes the binomial coefficient
    if (n == k):
        return 1
    if (k == 0):
        return 1
    numerator = 1
    denominator = 1
    for i in range(1, k+1):
        numerator = (n-k+i) * numerator
        denominator = i * denominator
    return numerator // denominator

distributions = []
for i in range(1,7):
    for j in range(1,7):
        for k in range(1,7):
            if  i+j+k == 7:
                distributions.append((i,j,k))
sum = 0
for partition in distributions:
    sum += (binomial_coefficient(7,partition[0]) * binomial_coefficient(7-partition[0],partition[1]))

for i in range(1,7):
    sum += binomial_coefficient(7,i)

# print(sum)


# a = [1,2]
# b = [3,[3,4],5]
# c=a+b
# a[1]=6
# b[1] = 9
# c[0] = 9
# print(a)
# print(b)
# print(c)

