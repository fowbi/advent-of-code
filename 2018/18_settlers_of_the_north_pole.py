#!/usr/bin/env python3

area = []
with open('./inputs/18.txt', 'r') as file:
    for line in file:
       area.append(list(line.replace("\n","")))

height = len(area)
width = len(area[0])


def determine_next_content(acre, i , j):
    surrounding = []
    for x in range(-1, 2):
        nx = i+x
        if nx < 0 or (nx >= height): continue
        for y in range(-1, 2):
            ny = j+y
            if ny < 0 or (nx == i and ny == j) or (ny >= width): continue
            surrounding.append(area[nx][ny])

    if acre == '.':
        new = '|' if surrounding.count('|') >= 3 else '.'
    elif acre == '|':
        new = '#' if surrounding.count('#') >= 3 else '|'
    else:
        new = '#' if (surrounding.count('#') >= 1 and surrounding.count('|') >= 1) else '.'

    return new

def print_area(area):
    for x in area:
        for y in x:
            print(y, end="")
        print()

print_area(area)
print()


previous_areas = []


for minute in range(1, 1000000001):
    new_area = [['.' for x in range(height)] for y in range(width)]
    wood = 0
    lumberyard = 0
    for i in range(height):
        for j in range(width):
            acre = determine_next_content(area[i][j], i, j)
            new_area[i][j] = acre

            wood += 1 if acre == '|' else 0
            lumberyard += 1 if acre == '#' else 0

    previous_areas.append(area)
    area = new_area

    # 28 is the period in which this examples repeats. Find it by comparing areas.
    if minute%28 == 20:
        print_area(area)
        print()
        print("wood", wood)
        print("lumberyard", lumberyard)
        print("total resources", wood * lumberyard)
