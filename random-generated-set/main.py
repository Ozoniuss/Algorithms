import random

import matplotlib.pyplot as plt

from collections import Counter


def select_random(nr_elems: int):
    i = random.randrange(1, nr_elems)
    return i


def get_random_numbers(nr_elems: list):

    steps = 0
    existing = []

    while True:
        i = select_random(nr_elems)
        steps += 1
        if i in existing:
            return steps

        existing.append(i)


def compute_once(number_elems: int, results: list):

    steps = get_random_numbers(number_elems)
    results.append(steps)


def compute(generations: int, number_elems: int):

    results = []
    for _ in range(generations):
        compute_once(number_elems, results)

    return results


def plot(results: list):

    c = Counter(results)

    p = plt.plot(
        c.keys(),
        c.values(),
    )
    plt.show()
    # Plot some data on the axes.


def main():

    number_elems = 1000
    generations = 100000

    ask = False
    if ask:
        number_elems = int(input("nr elems >>> "))

    res = compute(generations, number_elems)
    res.sort()

    plot(res)


main()
