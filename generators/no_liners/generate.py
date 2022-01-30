from src import group_by_color, unsig_range, write, time
import json

props = json.load(open('../props/props.json'))


# This algorithm is different than the stated on https://unsig.info/no-liners/
# Here I am retrieving the multipliers of an unsig, then adding them up by color,
# which allows me to determine if an unsig is a no-liner if none of these sums are above 1.
#
# This is valid because:
# - Sums by color above 1 means that they become another shape (they develop lines).
# - Unsigs can't have more than one property of a given color with a 0.5 multiplier.
#
def is_no_liner_props(props):
    props_by_color = group_by_color(props).values()
    def get_multiplier(x): return float(x["multiplier"])
    def sum_multipliers(x): return sum(list(map(get_multiplier, x)))
    sum_by_color = list(map(sum_multipliers, props_by_color))
    above_one_colors = list(filter(lambda x: x > 1, sum_by_color))
    return len(above_one_colors) == 0


def is_no_liner(n):
    return is_no_liner_props(props[n])


time(lambda: write(list(map(is_no_liner, unsig_range)), "no_liners.json"))
