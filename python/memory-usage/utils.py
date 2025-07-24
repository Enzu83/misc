from pympler import tracker

def get_mem_stats() -> dict:
    memory_tracker = tracker.SummaryTracker()
    summary = memory_tracker.create_summary()

    stats = {}
    for entry in summary:
        ref, n, sz = entry
        entry_type = ref.split(' ')[0]

        stat = stats.get(entry_type, {})
        stat['num'] = stat.get('num', 0) + n
        stat['sz'] = stat.get('sz', 0) + n * sz

        entries = stat.get('entries', [])
        entries.append([ref, n, sz])
        stat['entries'] = entries

        stats[entry_type] = stat

    return stats

def write_mem_stats(filename: str):
    with open(filename, "w") as f:
        f.write(str(get_mem_stats()))

