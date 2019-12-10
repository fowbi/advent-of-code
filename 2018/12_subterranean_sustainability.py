#!/usr/bin/env python3
import re

initial_state= "#..#.#..##......###...###"
# initial_state= "#.#..#..###.###.#..###.#####...########.#...#####...##.#....#.####.#.#..#..#.#..###...#..#.#....##."


combinations = dict()
with open('./inputs/12_test.txt') as file:
# with open('./inputs/12.txt') as file:
    for line in file:
        m = re.match("([#\.]{5}) => ([#\.])", str(line))
        combinations[m.group(1)] = m.group(2)

lowest_point = 0

current_state = ("."*10) + initial_state
print(0, current_state, sep="\t")
for generation in range(1, 3):
    new_state = ''

    # for position in range(1, len(current_state)):
    for position, pot in enumerate(current_state):
        offset = position-2 if position > 2 else 0
        # if position <= 2:
            # set = current_state[offset:3+position].rjust(5, '.')
        # else:
        set = current_state[offset:3+position].ljust(5, '.')

        if set in combinations:
            # print(position, set)
            new_state += combinations[set]
            # print(set, new_state)
        else:
            new_state += '.'
        print(set, new_state)


    print(generation, new_state, sep="\t")
    lowest_point += new_state.find('#') - current_state.find('#')


    current_state = ('.'*10)+new_state[new_state.find('#'):new_state.rfind('#')]


    # current_state = new_state


print(lowest_point)

point = lowest_point
sum = 0
for position, pot in enumerate(current_state):
    if pot == '#':
        sum += point

    point += 1

print(sum)

# sum = 0
# for position, pot in enumerate(current_state):
#     if pot == '#':
#         sum += (position-(len(initial_state)*2))
#
#
# print(sum)
# print(current_state.find('#')-(len(initial_state)*2))
# print(current_state.rfind('#')-(len(initial_state)*2))

