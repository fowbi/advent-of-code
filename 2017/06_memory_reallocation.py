#!/usr/bin/env python3

import math
from functools import reduce

input = []
with open('./inputs/06.txt') as file:
    for line in file:
        input.append(list(map(int, str(line).split())))


d = input[0]
# d = [int(i) for i in input().strip().split()]

for _ in [0]*2:
    s = set()
    while tuple(d) not in s:
        s.add(tuple(d))
        n = max(d)
        i = d.index(n)
        d[i] = 0
        for j in range(n):
            d[(i+j+1) % len(d)] += 1
    print(len(s))




def solve(original_state):
    cycles = 0
    states = []
    new_state = original_state.copy()

    while cycles == 1 or new_state not in states:
        states.append(new_state)
        new_state = new_state.copy()

        bank = max(new_state)
        dups = reduce(lambda x, y: x+1 if y == bank else x+0, new_state, 0) > 1

        num_of_banks = len(new_state) if not dups else len(new_state)-1

        split = math.ceil(bank/num_of_banks)
        rest = abs(bank - (split * num_of_banks))
        rest = rest if rest != 0 else split

        dup = False
        for k, i in enumerate(new_state):
            if i != bank:
                new_state[k] += split
            else:
                if dups and not dup:
                    new_state[k] = 0
                else:
                    new_state[k] = new_state[k]+rest if dup else rest
                dup = True

        cycles += 1

    return cycles

print(solve(input[0]))
