"""https://adventofcode.com/2020/day/15"""

import re

starting_numbers = []
with open('./input.txt', 'r') as file:
    for line in file:
        starting_numbers = list(map(lambda n: int(n.strip()), line.split(',')))

def calculate_number(starting_numbers, end_number):
    a_numbers = starting_numbers
    c = len(starting_numbers)
    numbers = {}
    previous = {}

    for i in range(end_number):
        if i < c:
            p = starting_numbers[i]
            numbers[p] = i
            continue;

        if p in previous:
            p = numbers[p] - previous[p]

            if p in numbers:
                previous[p] = numbers[p]
                numbers[p] = i
                continue;
        else:
            p = 0
            if p in numbers:
                previous[p] = numbers[p]
                numbers[p] = i
                continue;

        numbers[p] = i
    return p

def solve_first_part(starting_numbers):
    return calculate_number(starting_numbers, 2020)

def solve_second_part(starting_numbers):
    return calculate_number(starting_numbers, 30000000)

print('part 1:', solve_first_part(starting_numbers))
print('part 2:', solve_second_part(starting_numbers))
