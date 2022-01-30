import json

is_no_liner = json.load(open('no_liners.json'))

print(is_no_liner[30223])  # False
print(is_no_liner[11])  # False
print(is_no_liner[9])  # True
print(is_no_liner[2036])  # True
