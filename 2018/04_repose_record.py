#!/usr/bin/env python3
import operator
import re
from datetime import date, datetime

logs = []
sleep_log = []
with open('./inputs/04.txt') as file:
    for line in file:
        logs.append(str(line))

logs = sorted(logs)

id = None
minutes_asleep = dict()
time_fell_asleep = 0
total_sleep = dict()
for log in logs:
    dt = re.search('(\d{4})-(\d{2})-(\d{2})\s(\d{2}):(\d{2})\]', log)
    if "begins shift" in log:
        id = int(re.findall('#(\d+)', log)[0])
        if id not in minutes_asleep:
            minutes_asleep[id] = dict(list(enumerate(0 for i in range(60))))
            total_sleep[id] = 0
    elif "falls asleep" in log:
        """"""
        time_fell_asleep = int(dt.group(5))
    elif "wakes up" in log:
        """"""
        total_sleep[id] += (int(dt.group(5)) - time_fell_asleep)
        for i in range(time_fell_asleep, int(dt.group(5))):
            minutes_asleep[id][i] += 1

guard_id, _ = max(total_sleep.items(), key=lambda x: x[1])
minute, _ = max(minutes_asleep[guard_id].items(), key=lambda x: x[1])

max_minute = 0
max_guard = 0
max_times = 0

for g, m in minutes_asleep.items():
    minute, times = max(m.items(), key=lambda x: x[1])
    print(g, minute, times)

    if times > max_times:
        max_times = times
        max_minute = minute
        max_guard = g


print(guard_id * minute)
print(max_guard, max_minute, max_guard * max_minute)
