import random

# I'm generating the output strings cause I'm lazy

prefix = "abcdefghijklmnopqrstuvwxyz"

with open("a.txt", "w") as f:
    
    for i in range(100):
        add = [str(random.randint(1, 100)) for i in range(30)]
        line = '"' + prefix + ''.join(add) + '"'
        
        f.write(line)
        
        # I'm gonna copy paste this motherfucker
        f.write(",")
        f.write("\n")
        
