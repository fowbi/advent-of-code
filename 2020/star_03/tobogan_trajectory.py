"""https://adventofcode.com/2020/day/3"""

from collections import namedtuple
import re

def split(word):
    return [char for char in word]

field = []
with open('./input.txt') as file:
    for line in file:
        field.append(split(str(line).strip()))

def countTrees(field, right, down):
    col = row = 0
    counter = col = row = 0
    maxCol = len(field[0])
    maxRow = len(field) -1

    while row < maxRow:
        if (field[row][col] == '#'):
            counter += 1

        row += down
        col = (col + right) % maxCol

    return counter


def solve_first_part(field, right, down):
    return countTrees(field, right, down)

def solve_second_part(field, slopes):
    counter = 1

    for slope in slopes:
        counter *= countTrees(field, slope[0], slope[1])

    return counter


print(solve_first_part(field, 3, 1))
print(solve_second_part(field, [[1, 1], [3, 1], [5, 1] , [7, 1], [1, 2]]))
