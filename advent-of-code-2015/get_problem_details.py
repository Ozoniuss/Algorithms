import sys
import re

def title_to_dir(title: str):
    """ 
    Takes the problem's title and converts it to the format required to name the
    problem's directory.
    """

    # Using parentheses in the findall function only returns the capture group
    # of the regex. This is used in order to strip out the day number and 
    # actual title.
    day: str = re.findall(pattern=r"Day (\d\d?)", string=title)[0]
    name: str = re.findall(pattern=r"--- Day \d\d?: (.*) ---", string=title)[0]

    if len(day) == 1:
        day = "0"+day
    
    # Get the names as lowercase
    name_words = [n.lower() for n in name.split(" ")]
    
    return "-".join([day] + name_words)

# The html content of the webpage is passed as input to the script.
content = sys.argv[1]
title = re.findall(pattern=r"--- Day \d\d?:.*---", string=content)
print(title_to_dir(title[0]))