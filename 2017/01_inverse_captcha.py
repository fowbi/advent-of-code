"""https://adventofcode.com/2017/day/2"""

input = []
with open('./inputs/01.txt') as file:
    for line in file:
        input.append(str(int(line)))


def solve_first_part(input):
    total = 0
    l = len(input)
    for index, char in enumerate(input):
        next_index = 0 if index == l-1 else index+1
        total += int(char) if input[next_index] == char else 0

    return total


def solve_second_part(input):
    total = 0
    l = len(input)
    half_way = int(l/2)
    for index, char in enumerate(input):
        next_index = index + half_way if (index + half_way) < l else abs(half_way - index)
        total += int(char) if input[next_index] == char else 0

    return total


for numbers in input:
    print(solve_first_part(input=numbers))
    print(solve_second_part(input=numbers))
