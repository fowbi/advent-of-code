#!/usr/bin/env python3
import re

stars = []
with open('./inputs/25.txt') as file:
    for line in file:
        stars.append(list(map(int, re.findall("-?\d+", line))))

constellations = [
    [stars.pop()]
]


def is_in_constellation(star, s):
    return (abs(star[0]-s[0]) + abs(star[1]-s[1]) + abs(star[2]-s[2]) + abs(star[3]-s[3])) <= 3


while len(stars) > 0:
    found = False
    for constellation in constellations:
        for s in constellation:
            i = 0
            while i < len(stars):
                if is_in_constellation(s, stars[i]):
                    constellation.append(stars.pop(i))
                    found = True
                else:
                    i += 1

        print(len(constellations))

    if not found:
        constellations.append([stars.pop(0)])


print(constellations)
print(len(constellations))

exit(0)



for star in stars:
    not_found = True
    for constellation in constellations:
        if is_in_constellation(star, constellation):
            not_found = False
            constellation.append(star)
            break;

    if not_found:
        constellations.append([star])


print(constellations)

print(len(constellations))






