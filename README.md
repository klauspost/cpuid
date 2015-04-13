# cpuid
Package cpuid provides information about the CPU running the current program.

CPU features are detected on startup, and kept for fast access through the life of the application.
Currently x86 / x64 (AMD64) is supported, and no external C (cgo) code is used, which should make the library very easy to use.

You can access the CPU information by accessing the shared CPU variable of the cpuid library.

Package home: https://github.com/klauspost/cpuid

[![GoDoc][1]][2] [![Build Status][3]][4]

[1]: https://godoc.org/github.com/klauspost/cpuid?status.svg
[2]: https://godoc.org/github.com/klauspost/cpuid
[3]: https://travis-ci.org/klauspost/cpuid.svg
[4]: https://travis-ci.org/klauspost/cpuid

# features
Currently these CPU features are detected:
*	CMOV         (i686 CMOV)
*	AMD 3DNOW!                (AMD 3DNOW)
*	AMD 3DNNOW! Extension             (AMD 3DNowExt)
*	MMX                     (standard MMX)
*	MMX EXT                  (SSE integer functions or AMD MMX ext)
*	SSE                     (SSE functions)
*	SSE2                    (P4 SSE functions)
*	SSE3                    (Prescott SSE3 functions)
*	SSSE3                   (Conroe SSSE3 functions)
*	SSE4.1                   (Penryn SSE4.1 functions)
*	SSE4.2                   (Nehalem SSE4.2 functions)
*	AVX                     (AVX functions)
*	AVX2                    (AVX2 functions)
*	FMA3                    (Intel FMA 3)
*	FMA4                    (Bulldozer FMA4 functions)
*	XOP                     (Bulldozer XOP functions)
*	BMI1                    (Bit Manipulation Instruction Set 1)
*	BMI2                    (Bit Manipulation Instruction Set 2)
*	TBM                     (AMD Trailing Bit Manipulation)
*	LZCNT                   (LZCNT instruction)
*	AE-SNI                   (Advanced Encryption Standard New Instructions)
*	CLMUL                   (Carry-less Multiplication)

# installing

```go get github.com/klauspost/cpuid```

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
