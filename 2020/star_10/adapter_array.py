"""https://adventofcode.com/2020/day/10"""

data = []
with open('./input.txt', 'r') as file:
    for line in file:
        data.append(int(line))

def solve_first_part(adapters):
    usedAdapters = []
    joltage = 0
    differences = { 1: 0, 2: 0, 3: 0 }
    
    adapters.append(max(adapters) + 3)
    adapters.sort()

    while len(usedAdapters) < len(adapters):
        for adapter in adapters:
            if adapter in [joltage+1, joltage+2, joltage+3]:
                differences[adapter-joltage] +=1
                usedAdapters.append(adapter)
                joltage = adapter
                break;

    return differences[1] * differences[3]

def findValidConnections(adapters, start, usedConnections):
    found = 0
    if start == adapters[-1]:
        return 1

    for i in range(1, 4):
        n = start + i
        if n in adapters:
            if n not in usedConnections:
                usedConnections[n] = findValidConnections(adapters, n, usedConnections)
            found += usedConnections[n]

    return found

def solve_second_part(adapters):
    return findValidConnections(adapters, 0, {})

print('part 1:', solve_first_part(data))
print('part 2:', solve_second_part(data))
