#!/usr/bin/env python3
import re
import matplotlib.pyplot as plt

import numpy as np

data = {
    'points': [],
    'velocity': []
}
with open('./inputs/10.txt') as file:
    star = 0
    for line in file:
        m = re.match("position=<([^\>]+)> velocity=<([^\>]+)>", str(line))
        data['points'].append(list(map(int, m.group(1).split(', '))))
        data['velocity'].append(list(map(int, m.group(2).split(', '))))

points = np.array(data['points'], dtype="int16")
velocity = np.array(data['velocity'], dtype="int16")
num_of_points = len(points)

def build_and_print(points, counter):
    min_x, min_y = np.min(points, axis=0)
    max_x, max_y = np.max(points, axis=0)

    print(counter, int(max_x), int(min_x), int(max_y), int(min_y))
    print(int(max_x) - int(min_x) + int(max_y) - int(min_y))

    # plt.scatter(*zip(*points))
    # plt.show()
    # return points


def move_points(points, velocity, inc=1):
    for k, point in enumerate(points):
        point[0] += velocity[k][0] * inc
        point[1] += velocity[k][1] * inc


# build_and_print(points=points, counter=0)
# for i in range(1, 20000):
#     move_points(points, velocity)
#     build_and_print(points=points, counter=i)

def build(points):
    # plt.scatter(*zip(*points))
    # plt.show()

    for y in range(min_y, max_y):


move_points(points, velocity, inc=10727)
build(points)

plt.scatter(*zip(*points))
plt.show()

for i in range(1, 20000):
    move_points(points, velocity)
    build_and_print(points=points, counter=i)
