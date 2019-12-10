#!/usr/bin/env python3
from collections import deque
import numpy as np

width = 14
# width = 150

def print_grid(grid: deque):
    for x in grid:
        for y in x:
            print(y, end="")
        print()


carts = dict()
directions = deque(['left', 'straight', 'right'])
original_grid = []
row = 0
with open('./inputs/13_test.txt') as file:
# with open('./inputs/13.txt') as file:
    for line in file:
        m = [' ' for x in range(width)]
        for col, c in enumerate(list(line.rstrip())):
            m[col] = c

            if c in ['>', '<', '^', 'v']:
                carts[(row, col)] = {
                    'cart': c,
                    'direction': directions.copy(),
                    'previous': '-' if c in ['>', '<'] else '|’',
                }
        original_grid.append(m)
        row += 1


def is_whitespace(position, grid):
    return grid[position[0]][position[1]] not in ('|', '-', '>', '<', '^', 'v', '+', '\\', '/')


def fetch_surrounding_positions(position, grid):
    top = None if position[0]-1 < 0 else None if is_whitespace((position[0]-1, position[1]), grid) else (position[0]-1, position[1])
    right = None if position[1]+1 >= width else None if is_whitespace((position[0], position[1]+1), grid) else (position[0], position[1]+1)
    bottom = None if position[0]+1 >= len(grid) else None if is_whitespace((position[0]+1, position[1]), grid) else (position[0]+1, position[1])
    left = None if position[1]-1 < 0 else None if is_whitespace((position[0], position[1]-1), grid) else (position[0], position[1]-1)

    return top, right, bottom, left


grid = original_grid.copy()
for position, meta in carts.items():
    top, right, bottom, left = fetch_surrounding_positions(position, grid)
    print(position, top, right, bottom, left)


tracks = deque()
rows = len(original_grid)
# for y, row in enumerate(original_grid):
#     for x, c in enumerate(row):
        # if c == '/' and x+1 < width and original_grid[y][x+1] == '-':
            # tracks.append(trace_track((x, y), original_grid))
