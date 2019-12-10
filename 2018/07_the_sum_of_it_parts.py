#!/usr/bin/env python3
import re
import networkx

graph = networkx.DiGraph()

with open('./inputs/07.txt') as file:
    for line in file:
        m = re.search('Step (\w) must be finished before step (\w) can begin.', str(line))
        graph.add_edge(m.group(1), m.group(2))

print(''.join(networkx.lexicographical_topological_sort(graph)))
