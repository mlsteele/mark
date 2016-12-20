#!/usr/bin/env python
import sys
import time

last_line = time.time()

try:
    while True:
        line = sys.stdin.readline()

        now = time.time()
        if now - last_line > .3:
            sys.stdout.write("-" * 0 + "\n" * 3)
        last_line = now

        if line == "":
            break

        sys.stdout.write(line)
except KeyboardInterrupt:
    pass
