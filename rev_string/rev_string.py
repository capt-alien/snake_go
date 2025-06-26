#!/usr/bin/env python3

def rev_string(s):
    output = ""
    slist = list(s)
    counter = len(s) -1
    left = 0
    right = len(s) -1

    while right > left:
        slist[left], slist[right] = slist[right], slist[left]
        right -= 1
        left += 1

    return "".join(slist)


if __name__ == '__main__':
    s = "hello"
    print(rev_string(s))
