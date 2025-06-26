# Group Anagrams

## Description

This directory contains solutions to the "Group Anagrams" coding challenge.
The goal is to group a list of strings into collections of anagrams.

## Problem Statement

Given an array of strings `strs`, group the anagrams together.
You can return the answer in any order.

### Example

```
Input: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
Output: [["bat"], ["nat", "tan"], ["tea", "ate", "eat"]]
```

## Constraints

- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters

## Strategy

The main idea is to use a dictionary (hash map) where:
- The key is a tuple representing the character count or a sorted version of the string
- The value is a list of strings that match that key

This allows for fast grouping of all anagrams.

## Status

- [x] Python solution
- [ ] Go solution (coming soon)
