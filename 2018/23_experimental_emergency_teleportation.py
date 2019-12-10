#!/usr/bin/env python3
import re

from mpl_toolkits import mplot3d
import matplotlib.pyplot as plt


nanobots = []
with open('./inputs/23.txt') as file:
    for line in file:
        nanobots.append(list(map(int, re.findall("-?\d+", line))))

# fig = plt.figure()
# ax = fig.add_subplot(111, projection='3d')


rx, ry, rz, rr = max(nanobots, key=lambda x: x[3])
in_range = 0
# print('x','y','z','|','r','|','d','|','d1','d2', sep="\t")
# print('_______________________________________________________________________________')
for x,y,z,r in nanobots:
    d = (abs(x-rx) + abs(y-ry) + abs(z-rz))

    d2 = abs(x)+abs(y)+abs(z)

    # ax.scatter(x, y, z)

    # print(x,y,z,'|',r,'|',d,'|',d1,d2, sep="\t")

    if d <= rr:
        in_range += 1

print(in_range)

# plt.show()
