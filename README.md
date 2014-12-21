Description
===========
This package contains an efficient implementation
of a fixed size in-memory bitset.

The operations supported on a BitSet are:

- Get
- Set, SetAll
- Clear, ClearAll
- Invert, InvertAll

There is also a Parse method that allows construction of BitSets
from a string of zeroes and ones.

Example
=======

	package main
	
	import (
		"fmt"
		"github.com/lazybeaver/bitset"
	)
	
	func main() {
		bs := bitset.New(4)
		bs.Set(0)
		bs.Set(2)
		bs.InvertAll()
		if bs.Get(1) {
			bs.Clear(3)
		}
		fmt.Println(bs.String())
	}
	
Benchmarks
==========
The index bounds checking and the function call to get position and mask
are currently the most expensive parts of the implementation.

	$ go test -bench=.
	PASS
	BenchmarkGet	500000000	         6.34 ns/op
	BenchmarkSet	500000000	         6.41 ns/op
	BenchmarkInvert	500000000	         6.49 ns/op
	ok		github.com/lazybeaver/bitset	11.581s
