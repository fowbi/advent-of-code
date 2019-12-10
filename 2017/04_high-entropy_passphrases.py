"""https://adventofcode.com/2017/day/4"""
from itertools import permutations, product

input = []
with open('./inputs/04.txt') as file:
    for line in file:
        input.append(str(line).split())


def solve(words):
    _words = []
    for word in words:
        if word not in _words:
            _words.append(word)
            perms = list(map(''.join, permutations(word, len(word))))
            for perm in perms:
                _words.append(perm)
        else:
            print(word, False)
            return False

    print(word, True)
    return True



valid = 0
for line in input:
    valid += int(solve(line))
print(valid)
