# cpuid
Package cpuid provides information about the CPU running the current program.

CPU features are detected on startup, and kept for fast access through the life of the application.
Currently x86 / x64 (AMD64/i386) and ARM (ARM64) is supported, and no external C (cgo) code is used, which should make the library very easy to use.

You can access the CPU information by accessing the shared CPU variable of the cpuid library.

Package home: https://github.com/klauspost/cpuid

[![PkgGoDev](https://pkg.go.dev/badge/github.com/klauspost/cpuid)](https://pkg.go.dev/github.com/klauspost/cpuid)
[![Build Status][3]][4]

[3]: https://travis-ci.org/klauspost/cpuid.svg?branch=master
[4]: https://travis-ci.org/klauspost/cpuid

## installing

`go get -u github.com/klauspost/cpuid/v2`

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

	// Test if we have these specific features:
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

Not all operating systems provide ARM features directly 
and there is no safe way to do so for the rest.

However, a `DetectARM()` can be used if you are able to control your deployment,
it will detect CPU features. 
Otherwise a flag for detecting unsafe ARM features can be added. See below.
 
Note that currently only features are detected on ARM, 
no additional information is currently available. 

## Adding flags

It is possible to add flags that affects cpu detection.

For this the `Flags()` command is provided.

This must be called *before* `flag.Parse()` AND after the flags have been parsed `Detect()` must be called.

This means that any detection used in `init()` functions will not contain these flags.

Example:

```Go
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/klauspost/cpuid/v2"
)

func main() {
	cpuid.Flags()
	flag.Parse()
	cpuid.Detect()

	// Test if we have these specific features:
	if cpuid.CPU.Supports(cpuid.SSE, cpuid.SSE2) {
		fmt.Println("We have Streaming SIMD 2 Extensions")
	}
}
```

# license

This code is published under an MIT license. See LICENSE file for more information.
