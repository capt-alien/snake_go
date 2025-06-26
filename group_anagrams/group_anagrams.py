#!/usr/bin/env python3
from collections import defaultdict

def group_anagrams(strs):
    tracker = defaultdict(list)
    for s in strs:
        uid = "".join(sorted(s))
        tracker[uid].append(s)

    return list(tracker.values())


if __name__ == '__main__':
    strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
    print(group_anagrams(strs))
