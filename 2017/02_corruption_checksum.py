"""https://adventofcode.com/2017/day/2"""

matrix = []
with open('./inputs/02.txt') as file:
    for line in file:
        matrix.append(list(map(int, list(map(str.strip, line.split())))))


def solve(matrix):
    checksum = 0
    for line in matrix:

        for x in line:
            for y in line:
                if x%y == 0 and x!=y:
                    checksum += int(x/y)

        # line.sort(reverse=True,key=int)
        # checksum += reduce(lambda x,y: x + y, reduce(lambda x, y: [int(z/y) if z%y == 0 and y!=z else 0 for z in line], line))

    return checksum


print(solve(matrix=matrix))
