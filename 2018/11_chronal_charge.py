#!/usr/bin/env python3
import math

import numpy as np

input = 7689 # real
input = 42 # test

def print_grid(grid):
    for y in grid:
        for x in y:
            print(x, end="\t")
        print()


def calculate_power_level(x, y, input):
    rack_id = x + 10
    power_level = (rack_id) * y
    power_level += input
    power_level *= rack_id
    power_level = math.floor((power_level/100)%10)
    power_level -= 5
    return power_level

grid = [[0 for x in range(300)] for y in range(300)]

for y, y_values in enumerate(grid):
    for x, x_values in enumerate(y_values):
        grid[y][x] = calculate_power_level(x, y, input)


# sums = dict()
# for y in range(8):
#     # if y+3 > 10: continue
#     for x in range(8):
#         # if x+3 > 10: continue
#         square = np_grid[x:x+3, y:y+3]
#         sums[(x,y)] = sum(map(sum, square))
#
# print(sums)


sums = []
np_grid = np.array(grid)
#for i in range(10, 299):
for i in range(11, 13):
# for i in range(20, 30):
    print(i)
    for y in range(300-i):
        for x in range(300-i):
            square = np_grid[x:x+i, y:y+i]
            sums.append({
                'position': (y,x,i),
                'sum': sum(map(sum, square))
            })
    print(max(sums, key=lambda x: x['sum']))

# print(sums)
# for sum in sums:
    # print(sum)
# print(max(sums, key=lambda x: x['sum']))
