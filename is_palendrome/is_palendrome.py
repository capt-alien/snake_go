#!/usr/bin/env python3
import re

def is_palendrome(s):
    sl = re.sub(r'[^a-zA-Z0-9]', '', s).lower()
    return sl[:: -1] == sl


if __name__ == '__main__':
    s = "A man, a plan, a canal: Panama"
    print(is_palendrome(s))
