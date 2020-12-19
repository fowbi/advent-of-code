"""https://adventofcode.com/2020/day/13"""

import math
from functools import reduce

timestamp = 0
busses = []

with open('./input.txt', 'r') as file:
    data = []
    for line in file:
        data.append(line)

    timestamp = int(data[0])
    for index, bus in enumerate(data[1].split(',')):
        bus = bus.strip()
        if not bus == 'x':
            busses.append((int(bus), index))

def solve_first_part(timestamp, busses):
    next_bus_timestamp = 0
    next_bus = 0
    for bus, offset in busses:
        temp = (math.floor(timestamp / bus) + 1) * bus
        if next_bus_timestamp == 0 or next_bus_timestamp > temp:
            next_bus = bus
            next_bus_timestamp = temp

    return next_bus * (next_bus_timestamp - timestamp)

"""
https://rosettacode.org/wiki/Chinese_remainder_theorem#Python
"""
def chinese_remainder(n, a):
    sum = 0
    prod = reduce(lambda a, b: a*b, n)
    for n_i, a_i in zip(n, a):
        p = prod // n_i
        sum += a_i * mul_inv(p, n_i) * p
    return sum % prod

def mul_inv(a, b):
    b0 = b
    x0, x1 = 0, 1
    if b == 1: return 1
    while a > 1:
        q = a // b
        a, b = b, a%b
        x0, x1 = x1 - q * x0, x0
    if x1 < 0: x1 += b0
    return x1


def solve_second_part(busses):
    correct = False
    timestamp = 0
    counter = busses[0][0]

    rests, divisors = [], []
    for bus, offset in busses:
        rests.append((-offset) % bus)
        divisors.append(bus)

    return chinese_remainder(divisors, rests)

print('part 1:', solve_first_part(timestamp, busses))
print('part 2:', solve_second_part(busses))
