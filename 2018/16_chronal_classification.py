#!/usr/bin/env python3
from functools import reduce
from itertools import zip_longest


def grouper(iterable, n, padvalue=None):
    args = [iter(iterable)] * n
    return zip_longest(*args, fillvalue=padvalue)


def sanitize(line:str) -> str:
    return line.replace('[','').replace(']','').replace(',','').split()

opcodes = {
    # addition
    'addr': lambda register, a, b: register[a] + register[b],
    'addi': lambda register, a, b: register[a] + b,
    # multiplication
    'mulr': lambda register, a, b: register[a] * register[b],
    'muli': lambda register, a, b: register[a] * b,
    # logical and
    'banr': lambda register, a, b: register[a] & register[b],
    'bani': lambda register, a, b: register[a] & b,
    # logical or
    'borr': lambda register, a, b: register[a] | register[b],
    'bori': lambda register, a, b: register[a] | b,
    # assignment
    'setr': lambda register, a, b: register[a],
    'seti': lambda register, a, b: a,
    # greater-than testing
    'gtir': lambda register, a, b: 1 if a > register[b] else 0,
    'gtri': lambda register, a, b: 1 if register[a] > b else 0,
    'gtrr': lambda register, a, b: 1 if register[a] > register[b] else 0,
    # equality testing
    'eqir': lambda register, a, b: 1 if a == register[b] else 0,
    'eqri': lambda register, a, b: 1 if register[a] == b else 0,
    'eqrr': lambda register, a, b: 1 if register[a] == register[b] else 0,
}


instructions = list()
with open('./inputs/16_1.txt', 'r') as file:
# with open('./inputs/16_test.txt', 'r') as file:
    for lines in grouper(file, 4, ''):
        instruction = {}
        instruction['before'] = list(map(int,sanitize(lines[0])[1:5]))
        instruction['instruction'] = list(map(int, lines[1].split()))
        instruction['after'] = list(map(int,sanitize(lines[2])[1:5]))
        instructions += (instruction,)


_opcodes = [x for x in range(16)]
cs = []
known_opcodes = set()


def run(instructions, opcodes, _opcodes, known_opcodes, opcode_keys):
    for full_instruction in instructions:
        c = 0
        # o = ""
        # i = 100
        _before, instruction, after = full_instruction.values()

        for opcode, f in opcodes.items():
            if opcode in known_opcodes: continue

            before = _before.copy()
            before[instruction[3]] = f(before, instruction[1], instruction[2])

            if before == after:
                c+=1
                o = opcode
                i = instruction[0]

        if c == 1:
            return i,o
        # cs.append(c)


# print(cs)

# print(len(cs))
# print(reduce(lambda x, y : x+1 if y >=3 else x, cs, 0))


opc = [x for x in range(16)]
known = []
while len(known) < 16:
    __opcodes = [[] for x in range(16)]
    for i in range(16):
        if i not in opc: continue

        for instruction in instructions:
            before, instruction, after = instruction.values()
            if instruction[0] == i:
                for opcode, f in opcodes.items():
                    if str(opcode) in known: continue

                    _before = before.copy()
                    _before[instruction[3]] = f(_before, instruction[1], instruction[2])
                    if _before == after:
                        __opcodes[i].append(opcode)

        temp = set(__opcodes[i])
        if len(temp) == 1:
            o = temp.pop()
            known.append(o)
            opc[i] = o

        # print(known)
        #print(opc)

for key, opcode in enumerate(opc):
    print(key,opcode)


register = [0] * 4
with open('./inputs/16_2.txt', 'r') as file:
    for line in file:
        instruction = list(map(int, line.split()))

        opcode = opc[instruction[0]]
        f = opcodes[opcode]
        register[instruction[3]] = f(register, instruction[1], instruction[2])

print(register)
