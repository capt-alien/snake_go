# Valid Parentheses

## Problem

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['`, and `']'`, determine if the input string is **valid**.

A string is considered valid if:
1. Open brackets are closed by the same type of brackets.
2. Open brackets are closed in the correct order.
3. Every closing bracket has a corresponding open bracket of the same type.

## Examples

```python
Input: s = "()"
Output: True

Input: s = "()[]{}"
Output: True

Input: s = "(]"
Output: False

Input: s = "([)]"
Output: False

Input: s = "{[]}"
Output: True
```

Constraints
    •   Only (), {}, and [] characters will appear in the string.
    •   The input string will not be empty.
    •   Use a stack to solve this efficiently in O(n) time.
