import json

horizontal_pairs = json.load(open('horizontal_pairs.json'))

print(horizontal_pairs[0])
print([0, 1] in horizontal_pairs)

def is_horizontal_pair(pair):
    return pair in horizontal_pairs

print(is_horizontal_pair([26642, 26701]))
print(is_horizontal_pair([12142, 12397]))
print(is_horizontal_pair([15490, 16315]))
print(is_horizontal_pair([10796, 10798]))
print(is_horizontal_pair([1507, 144]))
print(is_horizontal_pair([148, 6]))
print(is_horizontal_pair([22849, 12736]))  # !!! Correct!
print(is_horizontal_pair([28418, 28353]))
print(is_horizontal_pair([23538, 23311]))
print(is_horizontal_pair([10, 231]))  # Should be true!

print(is_horizontal_pair([6121, 19880]))  # Should be false!
print(is_horizontal_pair([1189, 9110]))  # Should be false!
