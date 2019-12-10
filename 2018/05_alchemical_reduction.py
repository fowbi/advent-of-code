#!/usr/bin/env python3
import re

data = []
with open('./inputs/05.txt') as file:
    for line in file:
        data.append(str(line))

polymer = data[0]

list = '|'.join([chr(i)+chr(i).upper() for i in range(ord('a'),ord('z')+1)] \
       + [chr(i).upper()+chr(i) for i in range(ord('a'),ord('z')+1)])


alphabet = [chr(i) for i in range(ord('a'),ord('z')+1)]

min_polymer = None

for char in alphabet:
    reduced_polymer = polymer.replace(char, '').replace(char.upper(), '')

    polymer_sub = len(reduced_polymer)
    reduced_polymer = re.sub(list, '', reduced_polymer)
    while len(reduced_polymer) != polymer_sub:
        polymer_sub = len(reduced_polymer)

        reduced_polymer = re.sub(list, '', reduced_polymer)

    if min_polymer is None:
        min_polymer = len(reduced_polymer)

    min_polymer = min(min_polymer, len(reduced_polymer))

print(min_polymer)
