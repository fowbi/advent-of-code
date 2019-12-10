#!/usr/bin/env python3
import re

m = [[() for x in range(1000)] for y in range(1000)]
#m = [[() for x in range(10)] for y in range(20)]
ids = []

claims = 0
with open('./inputs/03.txt') as file:
    for line in file:
        parts = re.search('\#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)', str(line))
        id = int(parts.group(1))
        left = int(parts.group(2))
        top = int(parts.group(3))
        width = int(parts.group(4))
        height = int(parts.group(5))

        ids.append(id)
        for j in range(height):
            for i in range(width):
                m[top+j][left+i] += (id,)


for i in m:
    for j in i:
        if len(j) > 1:
            claims+=1
            for id in j:
                if id in ids:
                    ids.remove(id)

print(ids)
print(claims)
# print(m)
