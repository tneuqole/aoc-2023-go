from collections import deque
from datetime import datetime

from inputs import get_program, input9
from intcode import run

start = datetime.now()


def day9p1():
    p = get_program(input9)
    return run(p, inputs=deque([1]))


def day9p2():
    p = get_program(input9)
    return run(p, inputs=deque([2]))


if __name__ == "__main__":
    day9p1()
    day9p2()

    print(f"took={datetime.now() - start}")
