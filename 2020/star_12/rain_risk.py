"""https://adventofcode.com/2020/day/12"""

import re
import math

"""

    Action N means to move north by the given value.
    Action S means to move south by the given value.
    Action E means to move east by the given value.
    Action W means to move west by the given value.
    Action L means to turn left the given number of degrees.
    Action R means to turn right the given number of degrees.
    Action F means to move forward by the given value in the direction the ship is currently facing.

"""

cardinalDirections = ['N', 'E', 'S', 'W']

data = []
with open('./input.txt', 'r') as file:
    for line in file:
        m = re.match('(\w)(\d+)', str(line).strip())
        data.append((m[1], int(m[2])))

def determineCardinalDirection(currentCardinalDirection, direction, value):
    if direction == 'L':
        i = cardinalDirections.index(currentCardinalDirection) - int(value/90)

    if direction == 'R':
        i = cardinalDirections.index(currentCardinalDirection) + int(value/90)

    if i > len(cardinalDirections)-1:
        i = i - len(cardinalDirections)


    return cardinalDirections[i]

def rotateWayPoint(waypoint, origin, direction, value):
    value = value if direction == 'R' else (value * -1)
    radians = math.radians(value)
    x, y = waypoint
    offset_x, offset_y = origin

    adjusted_x = (x - offset_x)
    adjusted_y = (y - offset_y)

    cos_rad = math.cos(radians)
    sin_rad = math.sin(radians)

    qx = offset_x + cos_rad * adjusted_x + sin_rad * adjusted_y
    qy = offset_y + -sin_rad * adjusted_x + cos_rad * adjusted_y

    return (round(qx), round(qy))

def solve_first_part(data):
    instructions = {
            'N': lambda position, value: (position[0], position[1] + value),
            'E': lambda position, value: (position[0] + value, position[1]),
            'S': lambda position, value: (position[0], position[1] - value),
            'W': lambda position, value: (position[0] - value, position[1]),
    }

    end = (0, 0)
    cardinal_direction = 'E'

    for direction, value in data:
        if (direction == 'F'):
            end = instructions[cardinal_direction](end, value)
        elif (direction == 'L' or direction == 'R'):
            cardinal_direction = determineCardinalDirection(cardinal_direction, direction, value)
        else:
            end = instructions[direction](end, value)

    return abs(0-end[0]) + abs(0-end[1])

def solve_second_part(data):
    instructions = {
            'N': lambda position, value: (position[0], position[1] + value),
            'E': lambda position, value: (position[0] + value, position[1]),
            'S': lambda position, value: (position[0], position[1] - value),
            'W': lambda position, value: (position[0] - value, position[1]),
    }

    waypoint = (10, 1)
    end = (0, 0)
    cardinal_direction = 'E'

    for direction, value in data:
        if (direction == 'F'):
            end = (end[0] + (waypoint[0] * value), end[1] + (waypoint[1] * value))
        elif (direction == 'L' or direction == 'R'):
            waypoint = rotateWayPoint(waypoint, (0, 0), direction, value)
        else:
            waypoint = instructions[direction](waypoint, value)

    return abs(0-end[0]) + abs(0-end[1])

print('part 1:', solve_first_part(data))
print('part 2:', solve_second_part(data))
