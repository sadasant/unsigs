from src import group_props_by, group_by_color, unsig_range, write, time
from itertools import product
import functools
import json

props = json.load(open('../props/props.json'))


@functools.lru_cache(maxsize=31119)
def horizontal_match_group_a(n):
    n_props = props[n]
    rotations = group_props_by(n_props, "rotation")
    # Extracting the properties with 90 and 270 rotations.
    group_a = (rotations.get("90") or []) + (rotations.get("270") or [])
    return json.dumps(group_a)

@functools.lru_cache(maxsize=31119 * 2)
def horizontal_match_group_b(n_with_inverted):
    n = n_with_inverted[0]
    inverted = n_with_inverted[1]
    n_props = props[n]
    colors = group_by_color(n_props)
    # Extracting all the colors with only one property (Group B).
    one_property_colors = list(map(lambda x: x[0], list(
        filter(lambda x: len(x) == 1, colors.values()))))
    # Not including normals (Group B).
    distributions = group_props_by(one_property_colors, "distribution")
    group_b_no_normals = distributions.get("CDF") or []
    # Extracting all the properties have 0 and 180 rotations (Group B).
    rotations = group_props_by(group_b_no_normals, "rotation")
    group_b = []
    if inverted is not True:
        # Removing CDFs from the unsig on the left if they have a rotation of 180 (Group B).
        group_b = rotations.get("0") or []
    else:
        # Removing CDFs from the unsig on the right if they have a rotation of 0 (Group B).
        group_b = rotations.get("180") or []
    # As long as Group B mirrors one side with the other.
    copied_group_b = list(map(lambda x: x.copy(), group_b))
    list(map(lambda x: x.pop("rotation"), copied_group_b))
    # One side will mirror another side if the multiplier is 1 or above

    def remove_05s(element_props):
        if element_props["multiplier"] != "0.5":
            element_props.pop("multiplier")
    list(map(remove_05s, copied_group_b))
    return json.dumps(copied_group_b)


@functools.lru_cache(maxsize=31119 * 2)
def is_horizontal_pair(pair):
    left_n = pair[0]
    right_n = pair[1]
    # Group A.
    left_group_a = horizontal_match_group_a(left_n)
    right_group_a = horizontal_match_group_a(right_n)
    # Group B.
    left_group_b = horizontal_match_group_b((left_n, False))
    right_group_b = horizontal_match_group_b((right_n, True))
    # If groups A and B are equal between one unsig and the other one.
    # print("left_group_a", left_group_a)
    # print("right_group_a", right_group_a)
    # print("left_group_b", left_group_b)
    # print("right_group_b", right_group_b)
    return left_group_a == right_group_a and left_group_b == right_group_b


print(is_horizontal_pair((26642, 26701)))
print(is_horizontal_pair((12142, 12397)))
print(is_horizontal_pair((15490, 16315)))
print(is_horizontal_pair((10796, 10798)))
print(is_horizontal_pair((1507, 144)))
print(is_horizontal_pair((148, 6)))
print(is_horizontal_pair((22849, 12736)))  # !!! Correct!
print(is_horizontal_pair((28418, 28353)))
print(is_horizontal_pair((23538, 23311)))
print(is_horizontal_pair((10, 231)))  # Should be true!


# /show numbers:6121,19880,1189,9110 columns:2
# Should be false!
# test2x2 = [(6121, 1189), (19880, 9110)]
# print(is_2x2(test2x2))


def generate():
    all_horizontal_pairs = []
    print(int(31119 / 11))
    for indices in product(range(1, int(31119 / 11)), repeat=2):
        if indices[0] is not indices[1]:
            if is_horizontal_pair(indices):
                all_horizontal_pairs.append(indices)
    return all_horizontal_pairs

time(lambda: write(generate(), "horizontal_pairs.json"))

print(is_horizontal_pair((6121, 19880)))  # Should be false!
print(is_horizontal_pair((1189, 9110)))  # Should be false!
