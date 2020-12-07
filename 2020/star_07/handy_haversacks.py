"""https://adventofcode.com/2020/day/7"""

import re
from collections import namedtuple

Luggage = namedtuple('Luggage', ['main', 'others'])
Bag = namedtuple('Bag', ['color', 'quantity'], defaults=[None, 1])

def extractContainedBags(bags):
    if re.match('no other bags', contains):
        return []

    bags = contains.split(',')
    others = []
    for b in bags:
        m = re.match('(\d+)\s([a-z ]+)\sbag', b.strip())
        others.append(Bag(m[2].strip(), int(m[1].strip())))

    return others


luggage = []
with open('./input.txt', 'r') as file:
    for line in file:
        mainBag, contains = str(line).strip().split(' bags contain ')
        luggage.append(Luggage(Bag(mainBag.strip('.')), extractContainedBags(contains)))

def getColoredBags(color, luggage):
    colors = [color]
    for l in luggage:
        for b in l.others:
            if b.color == color:
                colors = colors + getColoredBags(l.main.color, luggage)

    return colors

def getAmountOfBags(bag, luggage):
    count = bag.quantity
    for l in luggage:
        if l.main.color == bag.color:
            for b in l.others:
                count += (bag.quantity * getAmountOfBags(b, luggage))

    return count


def solve_first_part(luggage):
    colors = getColoredBags("shiny gold", luggage)
    return len(list(set(colors))) -1

def solve_second_part(luggage):
    return getAmountOfBags(Bag("shiny gold", 1), luggage) - 1

print(solve_first_part(luggage))
print(solve_second_part(luggage))
