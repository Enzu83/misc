import ast
import sys

with open(sys.argv[1]) as f:
    data: list[dict] = ast.literal_eval(f.read())

for element in data:
    element["Type"] = element["Type"][7:-1]
    del element["Entries"]

for element in data:
    print(f"{element["Type"]:30} {element['NObjects']:8} objects, {element['Size'] / (1024*1024):8.2f} MB")