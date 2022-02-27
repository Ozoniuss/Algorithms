x = [1,2,3,4,5]
def f(x):
    return x % 2 == 0
filter = filter(f, x)
print(list(filter))
