"""https://adventofcode.com/2020/day/6"""

import re

groups = []
with open('./input.txt', 'r') as file:
    content = re.split(r"(?:\r?\n){2,}", file.read().strip())
    groups = []
    for block in content:
        g = block.split("\n");
        groups.append(g)

def solve_first_part(groups):
    total = 0
    for people in groups:
        answers = []
        for person in people:
            for answer in person:
                if (not answer in answers) and answer:
                    answers.append(answer)
        total = total + len(answers)

    return total

def solve_second_part(groups):
    total = 0
    for people in groups:
        numOfPeople = len(people)
        answers = {}
        for person in people:
            for answer in person:
                if answer in answers:
                    answers[answer] += 1
                else:
                    answers[answer] = 1

        for k in answers.keys():
            if answers[k] == numOfPeople:
                total += 1

    return total

print(solve_first_part(groups))
print(solve_second_part(groups))
