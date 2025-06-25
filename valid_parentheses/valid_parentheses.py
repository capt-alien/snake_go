#!/usr/bin/env python3

def valid_parentheses(s):
    v_map = {'{':'}',
             '(':')',
             '[':']'}
    t_stack = []

    for x in s:
        if x in v_map: # if its an opening paren
            t_stack.append(x)
        else:
            if not t_stack:
                return False
            opener = t_stack.pop()
            if v_map[opener] != x:
                return False

    return True


if __name__ == '__main__':
    s1 = "()"
    s2 = "()[]{}"
    s3 = "(]"
    s4 = "([)]"
    s5 = "{[]}"

    print(valid_parentheses(s3))
