# read Agent memory usage summary

import ast
import sys

with open(sys.argv[1]) as f:
    data = ast.literal_eval(f.read())

print(f"{sum(obj["Size"] for obj in data) / (1024*1024):8.2f} MB")
