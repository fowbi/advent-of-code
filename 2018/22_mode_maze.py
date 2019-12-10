#!/usr/bin/env python3

rocky, wet, narrow = 0, 1, 2


def generate_grid(depth, target):
    grid = {}

    for y in range(0, target[1] + 1):
        for x in range(0, target[0] + 1):
            if (x, y) in [(0, 0), target]:
                geo_index = 0
            elif x == 0:
                geo_index = y * 48271
            elif y == 0:
                geo_index = x * 16807
            else:
                geo_index = grid[(x-1, y)][0] * grid[(x, y-1)][0]

            erosion_level = (geo_index + depth) % 20183
            risk_level = erosion_level % 3
            grid[(x, y)] = (erosion_level, risk_level)

            # if risk_level == 0:
                # print('.', end="")
            # elif risk_level == 1:
                # print('=', end="")
            # elif risk_level == 2:
                # print('|', end="")
        # print()

    return grid


grid = generate_grid(510, (10, 10))
assert 114 == sum([value[1] for value in grid.values()])

grid = generate_grid(5355, (14, 796))
assert 11972 == sum([value[1] for value in grid.values()])
