"""https://adventofcode.com/2020/day/11"""

import sys

FLOOR = '.'
EMPTY_SEAT = 'L'
OCCUPIED = '#'

"""

    If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
    If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
    Otherwise, the seat's state does not change.

"""

data = []
with open('./input.txt', 'r') as file:
    for line in file:
        data.append(str(line).strip())

def hasOccupiedSeatsAdjacent(floorplan, y, x):
    counter = 0
    adjacent = [
        (y-1, x-1), (y-1, x), (y-1, x+1),
        (y, x-1), (y, x+1),
        (y+1, x-1), (y+1, x), (y+1, x+1),
    ]

    for position in adjacent:
        if position[0] < 0 or position[1] < 0:
            continue

        if position[0] >= len(floorplan) or position[1] >= len(floorplan[0]):
            continue

        if floorplan[position[0]][position[1]] == OCCUPIED:
            counter += 1

    return counter >= 4

def hasEmptySeatsAdjacent(floorplan, y, x):
    count = 0
    adjacent = [
        (y-1, x-1), (y-1, x), (y-1, x+1),
        (y, x-1), (y, x+1),
        (y+1, x-1), (y+1, x), (y+1, x+1),
    ]

    for position in adjacent:
        if position[0] < 0 or position[1] < 0:
            continue

        if position[0] >= len(floorplan) or position[1] >= len(floorplan[0]):
            continue

        if floorplan[position[0]][position[1]] == OCCUPIED:
            return False

    return True


def getVisibleOccupiedSeats(floorplan, x, y):
    count = 0
    positions = [
        (-1, -1), (-1, 0), (-1, 1),
        (0,  -1), (0,  1),
        (1,  -1), (1,  0), (1, 1)
    ]
    for (px, py) in positions:
        nx, ny = (x+px, y+py)
        while ((0 <= nx <= len(floorplan)-1) and (0 <= ny <= len(floorplan[0])-1)):
            if floorplan[nx][ny] == EMPTY_SEAT:
                break

            if floorplan[nx][ny] == OCCUPIED:
                count += 1
                break
            nx, ny = (nx+px, ny+py)

    return count

def printFloorplan(floorplan):
    p = ''
    for row in floorplan:
        p += (row  + '\n')
    print(p, end='\r')

def solve_first_part(floorplan):
    printFloorplan(floorplan)
    changed = True

    while changed == True:
        newFloorplan = []
        changed = False
        for y, row in enumerate(floorplan):
            newRow = ""
            for x, seat in enumerate(row):
                if seat == FLOOR:
                    newRow += FLOOR
                    continue

                if seat == OCCUPIED and hasOccupiedSeatsAdjacent(floorplan, y, x):
                    newRow += EMPTY_SEAT
                    changed = True
                    continue

                if seat == EMPTY_SEAT and hasEmptySeatsAdjacent(floorplan, y, x):
                    newRow += OCCUPIED
                    changed = True
                    continue
                
                newRow += seat
            newFloorplan.append(newRow)

        floorplan = newFloorplan
        printFloorplan(floorplan)

    occupied_seats = 0
    for row in floorplan:
        for col in row:
            if col == OCCUPIED:
                occupied_seats += 1

    return occupied_seats

def solve_second_part(floorplan):
    printFloorplan(floorplan)
    changed = True

    while changed == True:
        newFloorplan = []
        changed = False
        for x, row in enumerate(floorplan):
            newRow = ""
            for y, seat in enumerate(row):
                if seat == FLOOR:
                    newRow += FLOOR
                    continue

                visibleOccupiedSeats = getVisibleOccupiedSeats(floorplan, x, y)

                if seat == OCCUPIED and visibleOccupiedSeats >= 5:
                    newRow += EMPTY_SEAT
                    changed = True
                    continue

                if seat == EMPTY_SEAT and visibleOccupiedSeats == 0:
                    newRow += OCCUPIED
                    changed = True
                    continue
                
                newRow += seat
            newFloorplan.append(newRow)

        floorplan = newFloorplan
        printFloorplan(floorplan)

    occupied_seats = 0
    for row in floorplan:
        for col in row:
            if col == OCCUPIED:
                occupied_seats += 1

    return occupied_seats

print('part 1:', solve_first_part(data))
print('part 2:', solve_second_part(data))
