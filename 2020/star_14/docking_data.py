"""https://adventofcode.com/2020/day/14"""

import re
from collections import namedtuple

InstructionSet = namedtuple('InstructionSet', ['mask', 'instructions'])

all_instructions = []
with open('./input.txt', 'r') as file:
    instructions = []
    mask = ''
    for line in file:
        if "mask" in line:
            if len(instructions) > 0:
                all_instructions.append(InstructionSet(mask, instructions))

            m = re.match('mask = (\w+)', line)
            mask = m[1]
            instructions = []

            continue

        m = re.match('mem\[(\d+)\] = (\d+)', line)
        instructions.append((int(m[1]), int(m[2])))

    all_instructions.append(InstructionSet(mask, instructions))

def convert_to_bin(value):
    return bin(value).replace("0b", "").zfill(36)

def convert_to_dec(value):
    return int(value, 2)

def add(m, v):
    v = convert_to_bin(v)
    r = ''
    for i in range(36):
        r += v[i] if m[i] == 'X' else m[i]
    return convert_to_dec(r)

def calculate_sum(program):
    s = 0
    for p in program.items():
        s += int(p[1])
    return s

def sprintf(c, lst, value):
    for l in lst:
        i = value.find('X')
        value = value[:i] + l + value[i + 1:]

    return value;

def determine_addresses(mask, address):
    bin_address = convert_to_bin(address)
    all_addresses = []
    new_mask = ''
    for i, c in enumerate(mask):
        new_mask += bin_address[i] if c == '0' else c

    count = new_mask.count('X')
    iterations = 2**count

    for iteration in range(iterations):
        bin_iteration = bin(iteration).replace("0b", "").zfill(count)
        after = sprintf('X', list(bin_iteration), new_mask)
        all_addresses.append(convert_to_dec(after))

    return all_addresses

def solve_first_part(all_instructions):
    program = {}
    for instruction in all_instructions:
        for address, value in instruction.instructions:
            program[address] = add(instruction.mask, value)

    return calculate_sum(program)

def solve_second_part(all_instructions):
    program = {}
    for instruction in all_instructions:
        for address, value in instruction.instructions:
            addresses = determine_addresses(instruction.mask, address)
            for a in addresses:
                program[a] = value

    return calculate_sum(program)

print('part 1:', solve_first_part(all_instructions))
print('part 2:', solve_second_part(all_instructions))
