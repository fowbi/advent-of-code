#!/usr/bin/env python3
import difflib
from operator import itemgetter

ids = []
with open('./inputs/02.txt') as file:
    for line in file:
        ids.append(str(line).split()[0])

three = 0
two = 0

for id in ids:
    chars = list(id)

    counts = []
    for c in chars:
        a = chars.count(c)
        if a == 3 and 3 not in counts:
            three +=1
            counts.append(3)
        elif a == 2 and 2 not in counts:
            two +=1
            counts.append(2)


print(two * three)

ratios = {}
for id in ids:
    for _id in ids:
        if id == _id: continue
        ratio = difflib.SequenceMatcher(a=id, b=_id).ratio()

        if ratio > 0:
           match = ""
           for i in range(0, len(id)):
               if id[i] == _id[i]:
                   match += id[i]

           ratios[match] = ratio

print(sorted(ratios.items(), key=itemgetter(1)).pop())

