#!/usr/bin/env python3

input = 540561
input = 9

elf_a_recipe_score = 3
elf_b_recipe_score = 7
elf_a_position = 1
elf_b_position = 2
scoreboard = (3, 7, 0, 1)

print(scoreboard)
print('a',elf_a_recipe_score, elf_a_position)
print('b',elf_b_recipe_score, elf_b_position)

sum_of_recipe_scores = str(elf_a_recipe_score + elf_b_recipe_score)

elf_a_position = 1 + elf_a_recipe_score
print(elf_a_position)
if (elf_a_position) > (len(scoreboard)-1):
    elf_a_position = elf_a_position%len(scoreboard)
    elf_a_position = len(scoreboard) if elf_a_position == 0 else elf_a_position

print(elf_a_position)
exit(0)

elf_b_position += 1 + elf_b_recipe_score
if elf_b_position > (len(scoreboard)-1):
    elf_b_position = elf_b_position%len(scoreboard)
    print(elf_b_position,len(scoreboard))
    print(elf_b_position)
    elf_b_position = len(scoreboard) if elf_b_position == 0 else elf_b_position

print(elf_b_position)

print(scoreboard)
print('a',elf_a_recipe_score, elf_a_position)
print('b',elf_b_recipe_score, elf_b_position)

elf_a_recipe_score = scoreboard[elf_a_position-1]
elf_b_recipe_score = scoreboard[elf_b_position-1]


print(scoreboard)
print('a',elf_a_recipe_score, elf_a_position)
print('b',elf_b_recipe_score, elf_b_position)

for n in sum_of_recipe_scores:
    scoreboard += (int(n), )
