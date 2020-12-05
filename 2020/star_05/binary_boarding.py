"""https://adventofcode.com/2020/day/5"""

from collections import namedtuple

BoardingPass = namedtuple('BoardingPass', ['boardingPass', 'row', 'seat', 'id'])

boardingPasses = []
with open('./input.txt', 'r') as file:
    for line in file:
        boardingPasses.append(str(line).strip())

def parseBoardingPasses(boardingPassList):
    boardingPasses = []
    for boardingPass in boardingPassList:
        rows = range(0, 128)
        seats = range(0, 8)
        for index, char in enumerate(boardingPass):
            if index < 7:
                half = int(len(rows)/2)
                if char == 'F':
                    rows = rows[:half]
                if char == 'B':
                    rows = rows[half:]

            if index >= 7:
                half = int(len(seats)/2)
                if char == 'L':
                    seats = seats[:half]
                if char == 'R':
                    seats = seats[half:]

        boardingPasses.append(BoardingPass(boardingPass, rows[-1], seats[-1], (rows[-1] * 8) + seats[-1]))

    return boardingPasses

def boardingPassExists(passes, row, seat):
    for p in passes:
        if (p.row == row and p.seat == seat):
            return True
    return False

def solve_first_part(boardingPassList):
    boardingPasses = parseBoardingPasses(boardingPassList)

    largestId = 0
    for bp in boardingPasses:
        if (bp.id > largestId):
            largestId = bp.id

    return largestId

def solve_second_part(boardingPassList):
    boardingPasses = parseBoardingPasses(boardingPassList)

    for row in range(0, 127):
        for seat in range(0, 7):
            if not boardingPassExists(boardingPasses, row, seat):
                if boardingPassExists(boardingPasses, row-1, seat) and boardingPassExists(boardingPasses, row+1, seat):
                    return (row * 8) + seat

    return 0

print(solve_first_part(boardingPasses))
print(solve_second_part(boardingPasses))
