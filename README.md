# cpuid
Package cpuid provides information about the CPU running the current program.

CPU features are detected on startup, and kept for fast access through the life of the application.
Currently x86 / x64 (AMD64/i386) and ARM (ARM64) is supported, and no external C (cgo) code is used, which should make the library very easy to use.

You can access the CPU information by accessing the shared CPU variable of the cpuid library.

Package home: https://github.com/klauspost/cpuid

[![GoDoc][1]][2] [![Build Status][3]][4]

[1]: https://godoc.org/github.com/klauspost/cpuid?status.svg
[2]: https://godoc.org/github.com/klauspost/cpuid
[3]: https://travis-ci.org/klauspost/cpuid.svg?branch=master
[4]: https://travis-ci.org/klauspost/cpuid

## installing

```go get github.com/klauspost/cpuid/v2```

## example

```Go
package main

import (
	"fmt"
	"strings"

	"github.com/klauspost/cpuid/v2"
)

func main() {
	// Print basic CPU information:
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model)
	fmt.Println("Features:", fmt.Sprintf(strings.Join(cpuid.CPU.FeatureSet(), ",")))
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")

	// Test if we have a specific features:
	if cpuid.CPU.Supports(cpuid.SSE, cpuid.SSE2) {
		fmt.Println("We have Streaming SIMD 2 Extensions")
	}
}
```

Sample output:
```
>go run main.go
Name: Intel(R) Core(TM) i5-2540M CPU @ 2.60GHz
PhysicalCores: 2
ThreadsPerCore: 2
LogicalCores: 4
Family 6 Model: 42
Features: CMOV,MMX,MMXEXT,SSE,SSE2,SSE3,SSSE3,SSE4.1,SSE4.2,AVX,AESNI,CLMUL
Cacheline bytes: 64
We have Streaming SIMD 2 Extensions
```

## ARM64 feature detection

 -> TODO
 
## Adding flags

 -> TODO
 
# license

This code is published under an MIT license. See LICENSE file for more information.
