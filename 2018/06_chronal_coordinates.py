#!/usr/bin/env python3
import numpy as np

data = []
with open('./inputs/06_test.txt') as file:
    for line in file:
        data.append(str(line).split(', '))

points = np.array(data, dtype="int16")
min_x, min_y = np.min(points, axis=0)
max_x, max_y = np.max(points, axis=0)

finite_points = []

print(min_x,min_y,max_x,max_y)
for p in points:
    if p[0] == min_x or p[0] == max_x or p[1] == min_y or p[1] == max_y:
        """"""
    else:
        finite_points.append('.'.join([str(p[0]), str(p[1])]))

print(finite_points)


def find_closest(points):
    distances = set()
    for p in points:
        distances.add(('.'.join([str(p[0]), str(p[1])]), abs(p[0] - i) + abs(p[1] - j)))

    p, min_distance = min(distances, key=lambda x: x[1])

    similar = [(k,v) for k, v in distances if v == min_distance]

    return p, min_distance if len(similar) == 1 else None


# exit(0);
    #print(_p)

    # return (x,y), None/1


test = [['.' for x in range(11)] for y in range(11)]

area = dict()
for i in range(min_x, max_x+1):
    for j in range(min_y, max_y+1):

        # calc closest point |p_1 - q_1| + |p_2 - q_2|
        p, min_distance = find_closest(points)
        if min_distance is not None:
            area[p] = 0 if p not in area else area[p]+1

        upper = False
        if p == '.'.join([str(i),str(j)]):
            upper = True

        test[i][j] = p.replace('1.1','a').replace('1.6','b').replace('8.3','c').replace('3.4','d').replace('5.5','e').replace('8.9','f')
        if upper:
            test[i][j] = test[i][j].upper()


print(finite_points)


print(test)

print(area)
print(max(area[p] for p in area if p in finite_points))
