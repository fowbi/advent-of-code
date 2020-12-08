"""https://adventofcode.com/2020/day/8"""

from dataclasses import dataclass
from enum import Enum
from copy import deepcopy

class Op(str, Enum):
    ACCUMULATOR = "acc"
    JUMP = "jmp"
    NO_OPERATION = "nop"

@dataclass
class Instruction:
    id: int
    operation: Op
    argument: int

class BootCode:
    def __init__(self, instructions):
        self.pointer = 0
        self.instructions = instructions
        self.current = self.instructions[self.pointer]
        self.accumulator = 0
        self.instructionsRun = []
        self.completed = False

    def __iter__(self):
        return self

    def __next__(self):
        if self.current.id in self.instructionsRun:
            raise StopIteration

        self.instructionsRun.append(self.current.id)

        if self.current.operation == Op.ACCUMULATOR:
            self.accumulator += self.current.argument
            self.pointer += 1

        if self.current.operation == Op.JUMP:
            self.pointer += self.current.argument

        if self.current.operation == Op.NO_OPERATION:
            self.pointer += 1

        if self.pointer < len(self.instructions):
            self.current = self.instructions[self.pointer]
            return self.current
        
        self.completed = True
        raise StopIteration

    def wasCompleted(self):
        return self.completed

instructions = []
with open('./input.txt', 'r') as file:
    id = 0
    for line in file:
        operation, argument = str(line).strip().split(' ')
        instruction = Instruction(id, operation, int(argument))
        instructions.append(instruction)
        id += 1

def solve_first_part(instructions):
    bootCode = BootCode(instructions)
    for instruction in bootCode:
        pass

    return bootCode.accumulator

def solve_second_part(instructions):
    for index, instruction in enumerate(instructions):
        if instruction.operation == Op.ACCUMULATOR:
            continue;

        modified_instructions = deepcopy(instructions)

        if instruction.operation == Op.JUMP:
            modified_instructions[index].operation = Op.NO_OPERATION

        if instruction.operation == Op.NO_OPERATION:
            modified_instructions[index].operation = Op.JUMP

        bootCode = BootCode(modified_instructions)
        for instruction in bootCode:
            pass

        if bootCode.wasCompleted():
            return bootCode.accumulator
    return 0

print(solve_first_part(instructions))
print(solve_second_part(instructions))
