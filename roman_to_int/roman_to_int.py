#!/usr/bin/env python3



def roman_to_int(roman_num):
    total = 0
    value_map = {'I': 1,
                 'V': 5,
                 'X': 10,
                 'L': 50,
                 'C': 100,
                 'D': 500,
                 'M': 1000}

    for x in range(len(roman_num)):
        current_val = value_map[roman_num[x]]

        if x + 1 < len(roman_num) and value_map[roman_num[x]] < value_map[roman_num[x+1]]:
            total -= current_val
        else:
            total += current_val

    return total




if __name__ == '__main__':
    s1= "LVIII" #58
    s2 = "MCMXCIV" #1994
    print(roman_to_int(s2))
