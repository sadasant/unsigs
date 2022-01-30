import json
import pandas as pd
from src import time, group_by_color

color_oc_data = pd.read_csv("../../data/color_oc_data_int.csv")

def get_ocs(n):
    raw_props = color_oc_data.iloc[n-1, color_oc_data.columns.str.match("OC")]
    no_na = raw_props.dropna()
    return list(zip(no_na, no_na.index))

def generate():
    ocs = list(map(get_ocs, range(0, 31119)))
    print(len(ocs))
    ocs_json = json.dumps(ocs)
    f = open("ocs.json", "w")
    f.write(ocs_json)
    f.close()

time(generate)