"""https://adventofcode.com/2020/day/1"""

input = []
with open('./input.txt') as file:
    for line in file:
        input.append(str(int(line)))

def solve_first_part(input):
    for num_1 in input:
        n1 = int(num_1)
        for num_2 in input:
            n2 = int(num_2)
            if (n1 + n2 == 2020):
                return n1 * n2

    return 0

def solve_second_part(input):
    for num_1 in input:
        n1 = int(num_1)
        for num_2 in input:
            n2 = int(num_2)
            for num_3 in input:
                n3 = int(num_3)
                if (n1 + n2 + n3 == 2020):
                    return n1 * n2 * n3

    return 0


print(solve_first_part(input))
print(solve_second_part(input))
