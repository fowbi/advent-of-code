from math import sqrt, ceil

input = 312051

# while i <= input:

def pos(input):
    r = ceil(sqrt(input)) // 2
    # print(r)
    d = (2*r)-1
    # print(d)

    i = input - d*d - 1
    # print(i)

    return (r, i-r+1) if i < d else (r-i+d, r) if i < 2*d + 2 else (-r, r-i-1+2*d+2) if i < 3*d + 2 else (i-r-3*d-2, -r)

print(sum(abs(c) for c in pos(input)))

# a = (int(sqrt(input))+1)/2
# print(a)


m, s, i = {(0, 0): 1}, 1, 2
while s <= input:
    (x, y), i = pos(i), i + 1
    m[x, y] = s = sum(m.get((x + j % 3 - 1, y + j//3), 0) for j in range(-3, 6))
print(s)
