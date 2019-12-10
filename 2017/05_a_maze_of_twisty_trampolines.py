"""https://adventofcode.com/2017/day/5"""

input = []
with open('./inputs/05.txt') as file:
    for line in file:
        input.append(int(line))


def solve_with_recursion(offsets, start=0, steps=0):
    """Breaks because the recursion level is to high."""
    if start >= len(offsets):
        return steps
    else:
        jump = offsets[start]
        offsets[start] += (-1 if jump > 2 else 1)
        steps += 1
        return solve_with_recursion(offsets, start+jump, steps)


def solve2(offsets):
    steps = 0
    start = 0

    while start < len(offsets):
        jump = offsets[start]
        offsets[start] += (-1 if jump > 2 else 1)
        steps += 1
        start += jump

    return steps


# print(solve(offsets=input.copy()))
print(solve2(offsets=input.copy()))
