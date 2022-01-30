from src import write, time
from itertools import product
import json

horizontal_pairs = json.load(open('../horizontal_pairs/horizontal_pairs.json'))
horizontal_pairs = list(map(tuple, horizontal_pairs))

vertical_pairs = json.load(open('../vertical_pairs/vertical_pairs.json'))
vertical_pairs = list(map(tuple, vertical_pairs))

vertical_booleans = [[False] * 31119] * 31119
for vPair in vertical_pairs:
    vertical_booleans[vPair[0]][vPair[1]] = True

def is_vertical(a, b):
    return vertical_booleans[a][b]

print(vertical_booleans[0])

# 0 1
# 2 3
def is_2x2(square):
    if (square[0] is square[2]) or (square[1] is square[3]):
        return False
    return (is_vertical(square[0], square[2])
            and is_vertical(square[1], square[3]))


print(is_2x2((10796, 10798, 10818, 10820)))
# Duration: 0.07869732499966631
time(lambda: is_2x2((10796, 10798, 10818, 10820)))

# /show numbers:10,231,218,2275 columns:2
# Should be true!
print(is_2x2((10, 231, 218, 2275)))

# /show numbers:6121,19880,1189,9110 columns:2
# Should be false!
print(is_2x2((6121, 19880, 1189, 9110)))
print(sum((6121, 19880)), sum((1189, 9110)))


def generate():
    # result = []
    count = 0
    print(len(horizontal_pairs))
    for i in product(horizontal_pairs, repeat=2):
        if i[0] is not i[1]:
            square = (i[0][0], i[0][1], i[1][0], i[1][1])
            if count % 10000 == 0:
                print(count, end="\r")
            if is_2x2(square):
                count += 1
                # result.append(square)
    return []


time(lambda: write(generate(), "2x2s.json"))
