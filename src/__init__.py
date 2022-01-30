from itertools import groupby
from timeit import default_timer as timer
import json

unsig_range = range(0, 31119)


def write(result, file_name):
    result_json = json.dumps(result)
    f = open(file_name, "w")
    f.write(result_json)
    f.close()


def group_props_by(props, key):
    def get_key(x): return x[key]
    sorted_by_color = sorted(props, key=get_key)
    grouped_by_color = {k: list(g)
                        for k, g in groupby(sorted_by_color, get_key)}
    return grouped_by_color


def group_by_color(props):
    return group_props_by(props, "color")


def time(f):
    start = timer()
    f()
    end = timer()
    print("Duration:", end - start)
