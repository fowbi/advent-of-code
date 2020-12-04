"""https://adventofcode.com/2020/day/2"""

from collections import namedtuple
import re

Password = namedtuple('Password', ['min', 'max', 'char', 'password'])

passwords = []
with open('./input.txt') as file:
    for line in file:
        m = re.search('(\d+)-(\d+)\s(\w):\s(\w+)', str(line))
        passwords.append(Password(int(m.group(1)), int(m.group(2)), m.group(3), m.group(4)))

def solve_first_part(passwords):
    counter = 0
    for p in passwords:
        charCounter = p.password.count(p.char)
        if (charCounter >= p.min and charCounter <= p.max):
            counter = counter + 1;
            
    return counter

def solve_second_part(passwords):
    counter = 0
    for p in passwords:
        # skip passwords where both min and max are correct
        if (p.password[p.min-1] == p.char and p.password[p.max-1] == p.char):
            continue

        if (p.password[p.min-1] == p.char or p.password[p.max-1] == p.char):
            counter = counter + 1;

    return counter


print(solve_first_part(passwords))
print(solve_second_part(passwords))
