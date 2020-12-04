"""https://adventofcode.com/2020/day/4"""

from collections import namedtuple
import re

Passport = namedtuple('Passport', ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl',
    'pid', 'cid', 'block'])

def isValid(passport):
    if not passport.byr or not passport.iyr or not passport.eyr or not passport.hgt or not passport.hcl or not passport.ecl or not passport.pid:
        return False

    return True

"""
    byr (Birth Year) - four digits; at least 1920 and at most 2002.
    iyr (Issue Year) - four digits; at least 2010 and at most 2020.
    eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
    hgt (Height) - a number followed by either cm or in:
        If cm, the number must be at least 150 and at most 193.
        If in, the number must be at least 59 and at most 76.
    hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
    ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
    pid (Passport ID) - a nine-digit number, including leading zeroes.
    cid (Country ID) - ignored, missing or not.
"""
def isStrictValid(passport):
    if not isValid(passport):
        return False

    if int(passport.byr) not in range(1920, 2003):
        return False

    if int(passport.iyr) not in range(2010, 2021):
        return False

    if int(passport.eyr) not in range(2020, 2031):
        return False

    m = re.fullmatch('(\d+)(in|cm)', passport.hgt)
    if not m:
        return False

    if m.group(2) == 'cm' and int(m.group(1)) not in range(150, 194):
        return False

    if m.group(2) == 'in' and int(m.group(1)) not in range(59, 77):
        return False

    if not re.fullmatch('#[0-9a-f]{6}', passport.hcl):
        return False

    if not passport.ecl in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
        return False

    if not re.fullmatch('[0-9]{9}', passport.pid):
        return False

    return True


passports = []
with open('./input.txt', 'r') as file:
    content = file.read().split("\n\n");
    for block in content:
        m = re.split('\s|\n', block.strip())
        args = {'byr': '', 'iyr': '', 'eyr': '', 'hgt': '', 'hcl': '', 'ecl': '',
                'pid': '', 'cid': '', 'block': block}

        for line in m:
            if (not line):
                continue

            key, value = line.split(':')
            args[key] = value

        passports.append(Passport(**args))


def solve_first_part(passports):
    counter = 0

    for passport in passports:
        if isValid(passport):
            counter += 1

    return counter

def solve_second_part(passports):
    counter = 0

    for passport in passports:
        if isStrictValid(passport):
            counter += 1

    return counter


print(solve_first_part(passports))
print(solve_second_part(passports))
