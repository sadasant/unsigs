import json
import pandas as pd
from timeit import default_timer as timer


def prop_to_dict(prop):
    parts = prop.split(" | ")
    return {
        "color": parts[0],
        "distribution": parts[1],
        "rotation": parts[2],
        "multiplier": parts[3]
    }

# Datasource used by Redegg
unsigs = pd.read_csv("../../data/unsig_master_22.01.11.csv")

# Retrieves the props of an unsig based on the unsig number, from the datasource used by Redegg.


def get_props(n):
    raw_props = unsigs.iloc[n, unsigs.columns.str.match("prop\d")]
    no_nan_props = filter(lambda v: v == v, raw_props.values)
    return list(map(prop_to_dict, no_nan_props))


def generate():
    props = list(map(get_props, range(0, 31119)))
    props_json = json.dumps(props)
    f = open("props.json", "w")
    f.write(props_json)
    f.close()

start = timer()
generate()
end = timer()
print(end - start)
