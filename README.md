# GoPearls
Programming Pearls in Golang
<br />

## How to use 
<br />
Each problem is identified by its package import path and may be conveniently fetched, built, and installed using the go get command.
<br /><br />

### To run all the tests

```go
go test -tags all ./...
```

### To run all the benchmarks

```go 
go test -v -bench=. -tags benchmark  ./...
```
<br />

### Column 1

Problem | Description
------------ | -------------
[1.6.2](https://github.com/LuigiAndrea/GoPearls/tree/master/column1-oyster/bit-vectors)| Bit Vectors
[1.6.3; 1.6.7](https://github.com/LuigiAndrea/GoPearls/tree/master/column1-oyster/sort-file-with-bit-vectors) | Bitmap Sort
[1.6.4](https://github.com/LuigiAndrea/GoPearls/tree/master/column1-oyster/generate-k-random-integer) | Generate k unique random integers
[1.6.9](https://github.com/LuigiAndrea/GoPearls/tree/master/column1-oyster/sparse-vector) | Sparse Vector
<br />

### Column 2

Problem | Description
------------ | -------------
[2.6.1](https://github.com/LuigiAndrea/GoPearls/tree/master/column2-aha/anagram)| Anagram
[2.6.2](https://github.com/LuigiAndrea/GoPearls/tree/master/column2-aha/atleast-twice)| Find integer that appears at least twice in sequential file
[2.6.3](https://github.com/LuigiAndrea/GoPearls/tree/master/column2-aha/rotate)| Rotation algorithms
[2.6.8](https://github.com/LuigiAndrea/GoPearls/tree/master/column2-aha/k-element-subset)| k-element subset of the set that sums to at most t