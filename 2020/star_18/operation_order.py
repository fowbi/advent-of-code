"""https://adventofcode.com/2020/day/18"""

import re

expressions = []
with open('./input.txt', 'r') as file:
    for line in file:
        expressions.append(line.strip())


def calcSingleExpression(expression):
    operator = '+'
    calc = 0

    for c in expression.split():
        if c in ['+', '*']:
            operator = c
            continue

        calc = calc + int(c) if operator == '+' else calc * int(c)

    return calc

def calcSingleExpressionWithPrecedence(expression):
    operator = '+'
    calc = 0

    while True:
        match = re.search('(\d+\s\+\s\d+)', expression)
        if match == None:
            break;

        match_index_start, match_index_end = match.span(1)
        sub_calc = calcSingleExpression(match.group(1))
        expression = expression[:match_index_start] + str(sub_calc) + expression[match_index_end:]

    for c in expression.split():
        if c in ['+', '*']:
            operator = c
            continue

        calc = calc + int(c) if operator == '+' else calc * int(c)

    return calc


def calculateExpression(expression, withPrecedence):
    calc = 0

    while True:
        match = re.search('\(([\d\+\*\s]+)\)', expression)
        if match == None:
            break;

        match_index_start, match_index_end = match.span(0)
        sub_calc = calcSingleExpressionWithPrecedence(match.group(1)) if withPrecedence else calcSingleExpression(match.group(1))
        expression = expression[:match_index_start] + str(sub_calc) + expression[match_index_end:]

    return calcSingleExpressionWithPrecedence(expression) if withPrecedence else calcSingleExpression(expression)

def solve_first_part(expressions):
    calc = 0

    for expression in expressions:
        sub_calc = calculateExpression(expression, False)
        calc += sub_calc
        print(expression, '=', sub_calc)

    return calc

def solve_second_part():
    calc = 0

    for expression in expressions:
        sub_calc = calculateExpression(expression, True)
        calc += sub_calc
        print(expression, '=', sub_calc)

    return calc

print('part 1:', solve_first_part(expressions))
print('part 2:', solve_second_part())
