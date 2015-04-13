# cpuid
Package cpuid provides information about the CPU running the current program.

CPU features are detected on startup, and kept for fast access through the life of the application.
Currently x86 / x64 (AMD64) is supported.

 You can access the CPU information by accessing the shared CPU variable of the cpuid library.

Package home: https://github.com/klauspost/cpuid

[![GoDoc][1]][2] [![Build Status][3]][4]

[1]: https://godoc.org/github.com/klauspost/cpuid?status.svg
[2]: https://godoc.org/github.com/klauspost/cpuid
[3]: https://travis-ci.org/klauspost/cpuid.svg
[4]: https://travis-ci.org/klauspost/cpuid

# example

```Go
package main

import (
	"fmt"
	"github.com/klauspost/cpuid"
)

func main() {
	// Print basic CPU information:
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model)
	fmt.Println("Features:", cpuid.CPU.Features)
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)

	// Test if we have a specific feature:
	if cpuid.CPU.SSE() {
		fmt.Println("We have Streaming SIMD Extensions")
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
We have Streaming SIMD Extensions
```