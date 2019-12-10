#!/usr/bin/env python3
from functools import reduce
from itertools import cycle

frequencies = []
with open('./inputs/01.txt') as file:
    for line in file:
        frequencies.append(int(line))

frequency = reduce(lambda x, y: x+y, frequencies, 0)
print(frequency)

sums = []
c = cycle(frequencies)
sum = 0
while sum not in sums:
    sums.append(sum)
    sum += next(c)

print(sum)
