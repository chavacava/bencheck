# bencheck

Spots deltas in benchmarking stats produced by [benchstat](https://github.com/golang/perf/tree/master/cmd/benchstat)

## Usage

```
bencheck [-t float]
```
`bencheck` reads its input from from stdin and generates its output to stdout. 

## Example

```bash
$ benchstat old.txt new.txt | bencheck -t 2
GobEncode   56.4MB/s ± 1%  65.1MB/s ± 1%  +15.36%  (p=0.016 n=4+5) delta is greater than 2.5
$ echo $?
1
```
