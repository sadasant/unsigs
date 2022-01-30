from src import write, time
from itertools import product
import json

horizontal_pairs = json.load(open('../horizontal_pairs/horizontal_pairs.json'))
horizontal_pairs = list(map(tuple, horizontal_pairs))

vertical_pairs = json.load(open('../vertical_pairs/vertical_pairs.json'))
vertical_pairs = list(map(tuple, vertical_pairs))

# Maximum distance between elements in a vertical pair
max_vertical_distance = list(sorted(list(map(sum, vertical_pairs))))[-1] / 2
print(max_vertical_distance)


def is_horizontal(pair):
    return pair in horizontal_pairs


def is_vertical(pair):
    return pair in vertical_pairs


# 1 2
# 3 4
def is_2x2(square):
    unique = set(square)
    if len(unique) != 4:
        return False
    # 31108 is the maximum distance between two pairs
    if (square[0] + square[2]) > max_vertical_distance or (square[1] + square[3]) > max_vertical_distance:
        return False
    # print(is_horizontal(pair_a), is_horizontal(pair_b), is_vertical((pair_a[0], pair_b[0])), is_vertical((pair_a[1], pair_b[1])))
    return (is_horizontal((square[0], square[1]))
            and is_horizontal((square[2], square[3]))
            and is_vertical((square[0], square[2]))
            and is_vertical((square[1], square[3])))

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
    result = []
    print(len(horizontal_pairs))
    max_sum = 50
    for i in product(horizontal_pairs, repeat=2):
        if i[0] is not i[1] and sum(i[0]) < max_sum and sum(i[1]) < max_sum:
            square = (i[0][0], i[0][1], i[1][0], i[1][1])
            print(len(result), square, end="\r")
            if is_2x2(square):
                result.append(square)
    return result


time(lambda: write(generate(), "2x2s.json"))
