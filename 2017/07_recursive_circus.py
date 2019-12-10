#!/usr/bin/env python3
import re

input = []
with open('./inputs/07.txt') as file:
    discs = {}
    tops = []
    carriers = []

    for line in file:
        line = str(line)
        parts = line.replace(',', '').split()
        top = str(parts.pop(0))
        disc = {
            'weight': int(parts.pop(0).replace('(','').replace(')','')),
            'carriers': []
        }

        if re.search('->', line):
            parts.pop(0)
            disc['carriers'] = parts

        discs[top] = disc
        tops.append(top) # temp
        carriers.extend(parts) # temp

for top in tops:
    if top not in carriers:
        main = top

def tree(name, discs):
    disc = discs[name]
    # weight = (disc['weight'],)
    weight = disc['weight']
    if disc['carriers']:
        for carrier in disc['carriers']:
            weight += tree(carrier, discs)

    return weight


for carrier in discs[main]['carriers']:
    print(tree(carrier, discs))

