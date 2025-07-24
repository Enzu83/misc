# Performance Measurement

Go functions can have benchmarks to measure their performances.

We can benchmark them with `go test -bench=.`. 

## Optional arguments

To refine the benchmark, we can precise some behaviors:

- `count`: Set the number of times the benchmark test will run
- `benchtime`: Set the number of benchmark with iterations suffix `x` or set the duration of the benchmark with `s`.
- `benchmem`: Display memory consumption by the benchmark function.