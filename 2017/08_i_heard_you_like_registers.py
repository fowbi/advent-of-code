#!/usr/bin/env python3
from operator import itemgetter

instructions = []
with open('./inputs/08.txt') as file:
    for line in file:
        instructions.append(list(map(str, str(line).split())))

print(instructions)

memory = {}

inc = lambda a, b: a + b
dec = lambda a, b: a - b
gt = lambda a, b: a > b
gte = lambda a, b: a >= b
lt = lambda a, b: a < b
lte = lambda a, b: a <= b
eq = lambda a,b: a == b

for instruction in instructions:
    operator_a, operation, operator_b, iff, condition_operator_a, condition, condition_operator_b = instruction

    if operator_a not in memory:
        memory[operator_a] = 0
    if condition_operator_a not in memory:
        memory[condition_operator_a] = 0


    valid_condition = (condition == '>' and gt(condition_operator_a, condition_operator_b)) \
        or (condition == '>=' and gte(condition_operator_a, condition_operator_b)) \
        or (condition == '<' and lt(condition_operator_a, condition_operator_b)) \
        or (condition == '<=' and lte(condition_operator_a, condition_operator_b)) \
        or (condition == '==' and eq(condition_operator_a, condition_operator_b))

    print(valid_condition)

    # valid_condition = False

    if valid_condition:
        memory[operator_a] = inc(memory[operator_a], int(operator_b)) if operation == 'inc' else dec(memory[operator_a], int(operator_b))


print(memory)
print(sorted(memory.items(), key=itemgetter(1)).pop())

# print(memory)
# print(max(memory), memory[max(memory)])

#     if condition == '>':
#     elif condition == '>=':
#     elif condition == '<':
#     elif condition == '<=':
#     else:

