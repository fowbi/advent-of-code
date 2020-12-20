"""https://adventofcode.com/2020/day/16"""

import re

fields = {}
number_ranges = []
tickets = []
my_ticket = []
with open('./input.txt', 'r') as file:
    n, y = False, False
    for line in file:
        m = re.match('(.+): (\d+)-(\d+) or (\d+)-(\d+)', line)
        if m:
            fields[m[1]] = [range(int(m[2]), int(m[3])+1), range(int(m[4]), int(m[5])+1)]

        m = re.findall('(\d+)-(\d+)', line)
        for mm in m:
            number_ranges.append(range(int(mm[0]), int(mm[1])+1))

        if "your" in line:
            y = True
            continue;

        if "nearby" in line:
            n = True
            continue;

        if y:
            my_ticket  = list(map(lambda x: int(x.strip()), line.split(',')))
            y = False

        if n:
            tickets.append(list(map(lambda x: int(x.strip()), line.split(','))))

def in_number_ranges(value, number_ranges):
    for nr in number_ranges:
        if value in nr:
            return True
    return False

def valid_ticket(ticket, number_ranges):
    for value in ticket:
        if not in_number_ranges(value, number_ranges):
            return False
    return True

def determine_field(fields, nums, excluded_fields):
    field_names = list(fields.keys())
    for ef in excluded_fields:
        field_names.remove(ef)
    
    for field in fields.items():
        if field[0] in excluded_fields:
            continue
        for num in nums:
            if (not num in field[1][0]) and (not num in field[1][1]):
                field_names.remove(field[0])
                break

    return field_names

def solve_first_part(number_ranges, tickets):
    calc = 0
    for ticket in tickets:
        for value in ticket:
            if not in_number_ranges(value, number_ranges):
                calc += value
    return calc

def solve_second_part(fields, number_ranges, tickets):
    valid_tickets = []
    for ticket in tickets:
        if valid_ticket(ticket, number_ranges):
            valid_tickets.append(ticket)

    found = {}
    calc = 1
    while len(found) < len(valid_tickets[0]):
        for i in range(len(valid_tickets[0])):
            if i in found.keys():
                continue
            nums = [ticket[i] for ticket in valid_tickets]
            names = determine_field(fields, nums, found.values())

            if len(names) == 1:
                found[i] = names[0]

                if 'departure' in names[0]:
                    calc *= my_ticket[i]

    return calc

print('part 1:', solve_first_part(number_ranges, tickets))
print('part 2:', solve_second_part(fields, number_ranges, tickets))
