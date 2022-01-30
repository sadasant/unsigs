import json

vertical_pairs = json.load(open('vertical_pairs.json'))

def matches_vertically(pair):
    return pair in vertical_pairs

print(matches_vertically([26642,26701]))
print(matches_vertically([12142,12397]))
print(matches_vertically([16588,23666]))
print(matches_vertically([1507,148]))
print(matches_vertically([144,6]))
print(matches_vertically([23258,10370]))
print(matches_vertically([22888,28060]))

print("Faster?")

import numpy as np

vertical_booleans = np.full((31119 * 31119), False)
for vPair in vertical_pairs:
    vertical_booleans[vPair[0]*31119 + vPair[1]] = True

def is_vertical(a, b):
    return vertical_booleans[a*31119 + b]

print(is_vertical(26642,26701))
print(is_vertical(12142,12397))
print(is_vertical(16588,23666))
print(is_vertical(1507,148))
print(is_vertical(144,6))
print(is_vertical(23258,10370))
print(is_vertical(22888,28060))
