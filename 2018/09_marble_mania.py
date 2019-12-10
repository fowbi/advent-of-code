#!/usr/bin/env python3
from collections import deque


# a = deque([0,1,2,3,4])


def marble_game(players, last_marble_worth):
    circles = deque([0])
    for marble in range(1, last_marble_worth+1):
        if marble %23 == 0:
            circles.rotate(7)
            circles.pop()
            circles.rotate(1)
        else :
            circles.rotate(-1)
            circles.append(marble)
        print(circles)

    # print(circles)

    return 1

marble_game(10, 100)
exit(0)
assert marble_game(10, 1618) == 8317
assert marble_game(13, 7999) == 146373
assert marble_game(17, 1104) == 2764
assert marble_game(21, 6111) == 54718
assert marble_game(30, 5807) == 37305

print(marble_game(424, 71114))
