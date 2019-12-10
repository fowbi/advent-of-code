#!/usr/bin/env python3
import re

nodes = []
with open('./inputs/08.txt') as file:
    for line in file:
        nodes = [int(c) for c in line.split(' ')]


def check(nodes):
    child_nodes = nodes[0]
    meta_nodes =  nodes[1]

    rest = nodes[2:]
    total_sum = 0

    for i in range(child_nodes):
        rest, sum = check(rest)
        total_sum += sum

    for j in range(meta_nodes):
        total_sum += rest[j]

    return rest[meta_nodes:], total_sum

print(check(nodes))

