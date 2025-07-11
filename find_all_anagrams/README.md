🧩 find_anagrams.py

This script finds all starting indices of anagrams of a pattern string p in a larger string s. It uses a sliding window approach and character frequency comparison to efficiently match all anagram substrings.

🔧 Usage

python3 find_anagrams.py

🧠 How It Works

Keeps a fixed-size sliding window over s equal to the length of p

Compares character counts using dictionaries (hashmaps)

Records indices where the window's character distribution matches the pattern's

📦 Example

Input:

s = "cbaebabacd"
p = "abc"

Output:

[0, 6]

🚀 Why It Matters

This is a common interview question testing:

Hashmap usage

Sliding window technique

Understanding of string manipulation and pattern detection

