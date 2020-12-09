"""https://adventofcode.com/2020/day/9"""

data = []
with open('./input.txt', 'r') as file:
    for line in file:
        data.append(int(line))

def isValid(value, numbers):
    for n1 in numbers:
        for n2 in numbers:
            if (n1 != n2) and (n1 + n2 == value):
                return True

    return False

def findInvalidNumber(data, preamble):
    for index, num in enumerate(data):
        if index < preamble:
            continue

        f = index - preamble

        to = index

        if not isValid(num, data[f:to]):
            return num

    return 0

def solve_first_part(data, preamble):
    return findInvalidNumber(data, preamble)

def solve_second_part(data, preamble):
    invalidNumber = findInvalidNumber(data, preamble)

    check = data
    for index, num in enumerate(data):
        if (num >= invalidNumber):
            continue

        s = 0
        col = []
        for n2 in data[index:]:
            if s == invalidNumber:
                return min(col) + max(col)
                break

            if (n2 >= invalidNumber) or (s > invalidNumber):
                break
        
            s += n2
            col.append(n2)

    return 0

print(solve_first_part(data, 25))
print(solve_second_part(data, 25))
