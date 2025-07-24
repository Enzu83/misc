import ast
import sys

with open(sys.argv[1]) as f:
    data = ast.literal_eval(f.read())

top = sorted(data.items(), key=lambda x: x[1]['sz'], reverse=True)[:10]
for k, v in top:
    print(f"{k:30} {v['num']:8} objects, {v['sz'] / (1024*1024):8.2f} MB")