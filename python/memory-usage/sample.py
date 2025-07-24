# sample script that allocates a bunch of memory and write its memory usage in a file

from utils import write_mem_stats

zeros = [0] * 10_000_000
ones = [1] * 2_000_000

write_mem_stats("mem.txt")