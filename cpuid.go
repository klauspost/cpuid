// Copyright (c) 2015 Klaus Post, released under MIT License. See LICENSE file.

// Package cpuid provides information about the CPU running the current program.
//
// CPU features are detected on startup, and kept for fast access through the life of the application.
// Currently x86 / x64 (AMD64) is supported.
//
// You can access the CPU information by accessing the shared CPU variable of the cpuid library.
//
// Package home: https://github.com/klauspost/cpuid
package cpuid

import "strings"

// Vendor is a representation of a CPU vendor.
type Vendor int

const (
	Other Vendor = iota
	Intel
	AMD
)

const (
	CMOV        = 1 << iota // i686 CMOV
	AMD3DNOW                // AMD 3DNOW
	AMD3DNOWEXT             // AMD 3DNowExt
	MMX                     // standard MMX
	MMXEXT                  // SSE integer functions or AMD MMX ext
	SSE                     // SSE functions
	SSE2                    // P4 SSE functions
	SSE3                    // Prescott SSE3 functions
	SSSE3                   // Conroe SSSE3 functions
	SSE4                    // Penryn SSE4.1 functions
	SSE42                   // Nehalem SSE4.2 functions
	AVX                     // AVX functions
	AVX2                    // AVX2 functions
	FMA3                    // Intel FMA 3
	FMA4                    // Bulldozer FMA4 functions
	XOP                     // Bulldozer XOP functions
	BMI1                    // Bit Manipulation Instruction Set 1
	BMI2                    // Bit Manipulation Instruction Set 2
	TBM                     // AMD Trailing Bit Manipulation
	LZCNT                   // LZCNT instruction
	AESNI                   // Advanced Encryption Standard New Instructions
	CLMUL                   // Carry-less Multiplication
	HTT                     // Hyperthreading (enabled)

	// Performance indicators
	SSE2SLOW // SSE2 is supported, but usually not faster
	SSE3SLOW // SSE3 is supported, but usually not faster
	ATOM     // Atom processor, some SSSE3 instructions are slower
)

var flagNames = map[Flags]string{
	CMOV:        "CMOV",        // i686 CMOV
	AMD3DNOW:    "AMD3DNOW",    // AMD 3DNOW
	AMD3DNOWEXT: "AMD3DNOWEXT", // AMD 3DNowExt
	MMX:         "MMX",         // Standard MMX
	MMXEXT:      "MMXEXT",      // SSE integer functions or AMD MMX ext
	SSE:         "SSE",         // SSE functions
	SSE2:        "SSE2",        // P4 SSE2 functions
	SSE3:        "SSE3",        // Prescott SSE3 functions
	SSSE3:       "SSSE3",       // Conroe SSSE3 functions
	SSE4:        "SSE4.1",      // Penryn SSE4.1 functions
	SSE42:       "SSE4.2",      // Nehalem SSE4.2 functions
	AVX:         "AVX",         // AVX functions
	AVX2:        "AVX2",        // AVX functions
	FMA3:        "FMA3",        // Intel FMA 3
	FMA4:        "FMA4",        // Bulldozer FMA4 functions
	XOP:         "XOP",         // Bulldozer XOP functions
	BMI1:        "BMI1",        // Bit Manipulation Instruction Set 1
	BMI2:        "BMI2",        // Bit Manipulation Instruction Set 2
	TBM:         "TBM",         // AMD Trailing Bit Manipulation
	LZCNT:       "LZCNT",       // LZCNT instruction
	AESNI:       "AESNI",       // Advanced Encryption Standard New Instructions
	CLMUL:       "CLMUL",       // Carry-less Multiplication
	HTT:         "HTT",         // Hyperthreading (enabled)

	// Performance indicators
	SSE2SLOW: "SSE2SLOW", // SSE2 supported, but usually not faster
	SSE3SLOW: "SSE3SLOW", // SSE3 supported, but usually not faster
	ATOM:     "ATOM",     // Atom processor, some SSSE3 instructions are slower

}

// CPUInfo contains information about the detected system CPU.
type CPUInfo struct {
	BrandName      string // Brand name reported by the CPU
	VendorID       Vendor // Comparable CPU vendor ID
	Features       Flags  // Features of the CPU
	PhysicalCores  int    // Number of physical processor cores in your CPU. Will be 0 if undetectable.
	ThreadsPerCore int    // Number of threads per physical core. Will be 1 if undetectable.
	LogicalCores   int    // Number of physical cores times threads that can run on each core through the use of hyperthreading. Will be 0 if undetectable.
	Family         int    // CPU family number
	Model          int    // CPU model number
	CacheLine      int    // Cache line size in bytes. Will be 0 if undetectable.
}

// CPU contains information about the CPU as detected on startup,
// or when Detect last was called.
//
// Use this as the primary entry point to you data,
// this way queries are
var CPU CPUInfo

func init() {
	Detect()
}

// Detect will re-detect current CPU info.
// This will replace the content of the exported CPU variable.
//
// Unless you expect the CPU to change while you are running your program
// you should not need to call this function.
// If you call this, you must ensure that no other goroutine is accessing the
// exported CPU variable.
func Detect() {
	CPU.BrandName = brandName()
	CPU.CacheLine = cacheLine()
	CPU.Family, CPU.Model = familyModel()
	CPU.Features = support()
	CPU.ThreadsPerCore = threadsPerCore()
	CPU.LogicalCores = logicalCores()
	CPU.PhysicalCores = physicalCores()
	CPU.VendorID = vendorID()
}

// Generated here: http://play.golang.org/p/TMb5NyBvq1

// Cmov indicates support of CMOV instructions
func (c CPUInfo) Cmov() bool {
	return c.Features&CMOV != 0
}

// Amd3dnow indicates support of AMD 3DNOW! instructions
func (c CPUInfo) Amd3dnow() bool {
	return c.Features&AMD3DNOW != 0
}

// Amd3dnowExt indicates support of AMD 3DNOW! Extended instructions
func (c CPUInfo) Amd3dnowExt() bool {
	return c.Features&AMD3DNOWEXT != 0
}

// MMX indicates support of MMX instructions
func (c CPUInfo) MMX() bool {
	return c.Features&MMX != 0
}

// MMXExt indicates support of MMXEXT instructions
// (SSE integer functions or AMD MMX ext)
func (c CPUInfo) MMXExt() bool {
	return c.Features&MMXEXT != 0
}

// SSE indicates support of SSE instructions
func (c CPUInfo) SSE() bool {
	return c.Features&SSE != 0
}

// SSE2 indicates support of SSE 2 instructions
func (c CPUInfo) SSE2() bool {
	return c.Features&SSE2 != 0
}

// SSE3 indicates support of SSE 3 instructions
func (c CPUInfo) SSE3() bool {
	return c.Features&SSE3 != 0
}

// SSSE3 indicates support of SSSE 3 instructions
func (c CPUInfo) SSSE3() bool {
	return c.Features&SSSE3 != 0
}

// SSE4 indicates support of SSE 4 (also called SSE 4.1) instructions
func (c CPUInfo) SSE4() bool {
	return c.Features&SSE4 != 0
}

// SSE42 indicates support of SSE4.2 instructions
func (c CPUInfo) SSE42() bool {
	return c.Features&SSE42 != 0
}

// AVX indicates support of AVX instructions
// and operating system support of AVX instructions
func (c CPUInfo) AVX() bool {
	return c.Features&AVX != 0
}

// AVX2 indicates support of AVX2 instructions
func (c CPUInfo) AVX2() bool {
	return c.Features&AVX2 != 0
}

// FMA3 indicates support of FMA3 instructions
func (c CPUInfo) FMA3() bool {
	return c.Features&FMA3 != 0
}

// FMA4 indicates support of FMA4 instructions
func (c CPUInfo) FMA4() bool {
	return c.Features&FMA4 != 0
}

// XOP indicates support of XOP instructions
func (c CPUInfo) XOP() bool {
	return c.Features&XOP != 0
}

// BMI1 indicates support of BMI1 instructions
func (c CPUInfo) BMI1() bool {
	return c.Features&BMI1 != 0
}

// BMI2 indicates support of BMI2 instructions
func (c CPUInfo) BMI2() bool {
	return c.Features&BMI2 != 0
}

// TBM indicates support of TBM instructions
// (AMD Trailing Bit Manipulation)
func (c CPUInfo) TBM() bool {
	return c.Features&TBM != 0
}

// Lzcnt indicates support of LZCNT instruction
func (c CPUInfo) Lzcnt() bool {
	return c.Features&LZCNT != 0
}

// HTT indicates the processor has Hyperthreading enabled
func (c CPUInfo) HTT() bool {
	return c.Features&HTT != 0
}

// SSE2Slow indicates that SSE2 may be slow on this processor
func (c CPUInfo) SSE2Slow() bool {
	return c.Features&SSE2SLOW != 0
}

// SSE3Slow indicates that SSE3 may be slow on this processor
func (c CPUInfo) SSE3Slow() bool {
	return c.Features&SSE3SLOW != 0
}

// AesNi indicates support of AES-NI instructions
// (Advanced Encryption Standard New Instructions)
func (c CPUInfo) AesNi() bool {
	return c.Features&AESNI != 0
}

// Clmul indicates support of CLMUL instructions
// (Carry-less Multiplication)
func (c CPUInfo) Clmul() bool {
	return c.Features&CLMUL != 0
}

// Atom indicates an Atom processor
func (c CPUInfo) Atom() bool {
	return c.Features&ATOM != 0
}

// Intel returns true if vendor is recognized as Intel
func (c CPUInfo) Intel() bool {
	return c.VendorID == Intel
}

// AMD returns true if vendor is recognized as AMD
func (c CPUInfo) AMD() bool {
	return c.VendorID == AMD
}

// Flags contains detected cpu features and caracteristics
type Flags uint64

// String returns a string representation of the detected
// CPU features.
func (f Flags) String() string {
	return strings.Join(f.Strings(), ",")
}

// Strings returns and array of the detected features.
func (f Flags) Strings() []string {
	s := support()
	r := make([]string, 0, 20)
	for i := uint(0); i < 64; i++ {
		key := Flags(1 << i)
		val := flagNames[key]
		if s&key != 0 {
			r = append(r, val)
		}
	}
	return r
}

func maxExtendedFunction() uint32 {
	eax, _, _, _ := cpuid(0x80000000)
	return eax
}

func maxFunctionID() uint32 {
	a, _, _, _ := cpuid(0)
	return a
}

func brandName() string {
	if maxExtendedFunction() >= 0x80000004 {
		name := make([]byte, 0, 64)
		for i := uint32(0); i < 3; i++ {
			a, b, c, d := cpuid(0x80000002 + i)
			name = append(name, valAsString(a, b, c, d)...)
		}
		return strings.Trim(string(name), " ")
	}
	return "unknown"
}

func threadsPerCore() int {
	if maxFunctionID() < 0xb {
		return 1
	}

	_, b, _, _ := cpuidex(0xb, 0)
	return int(b & 0xffff)
}

func logicalCores() int {
	switch vendorID() {
	case Intel:
		if maxFunctionID() < 0xb {
			return 0
		}
		_, b, _, _ := cpuidex(0xb, 1)
		return int(b & 0xffff)
	case AMD:
		_, b, _, _ := cpuid(1)
		return int((b >> 16) & 0xff)
	default:
		return 0
	}
}

func familyModel() (int, int) {
	if maxFunctionID() < 0x1 {
		return 0, 0
	}
	eax, _, _, _ := cpuid(1)
	family := ((eax >> 8) & 0xf) + ((eax >> 20) & 0xff)
	model := ((eax >> 4) & 0xf) + ((eax >> 12) & 0xf0)
	return int(family), int(model)
}

func physicalCores() int {
	switch vendorID() {
	case Intel:
		return logicalCores() / threadsPerCore()
	case AMD:
		_, _, c, _ := cpuid(0x80000008)
		return int(c&0xff) + 1
	default:
		return 0
	}
}

func vendorID() Vendor {
	_, b, c, d := cpuid(0)
	if b == 0x756e6547 && d == 0x49656e69 && c == 0x6c65746e {
		return Intel
	} else if b == 0x68747541 && d == 0x69746e65 && c == 0x444d4163 {
		return AMD
	} else {
		return Other
	}
}

func cacheLine() int {
	if maxFunctionID() < 0x1 {
		return 0
	}

	_, ebx, _, _ := cpuid(1)
	cache := (ebx & 0xff00) >> 5 // cflush size
	if cache == 0 && maxExtendedFunction() >= 0x80000006 {
		_, _, ecx, _ := cpuid(0x80000006)
		cache = ecx & 0xff // cacheline size
	}
	// TODO: Read from Cache and TLB Information
	return int(cache)
}

func support() Flags {
	mfi := maxFunctionID()
	if mfi < 0x1 {
		return 0
	}
	rval := uint64(0)
	_, b, c, d := cpuid(1)
	if (d & (1 << 15)) != 0 {
		rval |= CMOV
	}
	if (d & (1 << 23)) != 0 {
		rval |= MMX
	}
	if (d & (1 << 25)) != 0 {
		rval |= MMXEXT
	}
	if (d & (1 << 25)) != 0 {
		rval |= SSE
	}
	if (d & (1 << 26)) != 0 {
		rval |= SSE2
	}
	if (c & 1) != 0 {
		rval |= SSE3
	}
	if (c & 0x00000200) != 0 {
		rval |= SSSE3
	}
	if (c & 0x00080000) != 0 {
		rval |= SSE4
	}
	if (c & 0x00100000) != 0 {
		rval |= SSE42
	}
	if (c & (1 << 25)) != 0 {
		rval |= AESNI
	}
	if (c & (1 << 1)) != 0 {
		rval |= CLMUL
	}
	if (c & (1 << 28)) != 0 {
		// This field does not indicate that Hyper-Threading
		// Technology has been enabled for this specific processor.
		// To determine if Hyper-Threading Technology is supported,
		// check the value returned in EBX[23:16]
		v := (b >> 16) & 255
		if v > 0 {
			rval |= HTT
		}
	}

	// Check OXSAVE and AVX bits
	if (c & 0x18000000) == 0x18000000 {
		// Check for OS support
		eax, _ := xgetbv(0)
		if (eax & 0x6) == 0x6 {
			rval |= AVX
			if (c & 0x00001000) != 0 {
				rval |= FMA3
			}
		}
	}

	// Check AVX2, AVX2 requires OS support, but BMI1/2 don't.
	if mfi >= 7 {
		_, ebx, _, _ := cpuid(7)
		if (rval&AVX) != 0 && (ebx&0x00000020) != 0 {
			rval |= AVX2
		}
		if (ebx & 0x00000008) != 0 {
			rval |= BMI1
			if (ebx & 0x00000100) != 0 {
				rval |= BMI2
			}
		}
	}

	if maxExtendedFunction() >= 0x80000001 {
		_, _, c, d := cpuid(0x80000001)
		if (c & 0x00000020) != 0 {
			rval |= LZCNT
		}
		if (d & (1 << 31)) != 0 {
			rval |= AMD3DNOW
		}
		if (d & (1 << 30)) != 0 {
			rval |= AMD3DNOWEXT
		}
		if (d & (1 << 23)) != 0 {
			rval |= MMX
		}
		if (d & (1 << 22)) != 0 {
			rval |= MMXEXT
		}
		/* Allow for selectively disabling SSE2 functions on AMD processors
		   with SSE2 support but not SSE4a. This includes Athlon64, some
		   Opteron, and some Sempron processors. MMX, SSE, or 3DNow! are faster
		   than SSE2 often enough to utilize this special-case flag.
		   AV_CPU_FLAG_SSE2 and AV_CPU_FLAG_SSE2SLOW are both set in this case
		   so that SSE2 is used unless explicitly disabled by checking
		   AV_CPU_FLAG_SSE2SLOW. */
		if vendorID() != Intel &&
			rval&SSE2 != 0 && (c&0x00000040) == 0 {
			rval |= SSE2SLOW
		}

		/* XOP and FMA4 use the AVX instruction coding scheme, so they can't be
		 * used unless the OS has AVX support. */
		if (rval & AVX) != 0 {
			if (c & 0x00000800) != 0 {
				rval |= XOP
			}
			if (c & 0x00010000) != 0 {
				rval |= FMA4
			}
		}

		if vendorID() == Intel {
			family, model := familyModel()
			if family == 6 && (model == 9 || model == 13 || model == 14) {
				/* 6/9 (pentium-m "banias"), 6/13 (pentium-m "dothan"), and
				 * 6/14 (core1 "yonah") theoretically support sse2, but it's
				 * usually slower than mmx. */
				if (rval & SSE2) != 0 {
					rval |= SSE2SLOW
				}
				if (rval & SSE3) != 0 {
					rval |= SSE3SLOW
				}
			}
			/* The Atom processor has SSSE3 support, which is useful in many cases,
			 * but sometimes the SSSE3 version is slower than the SSE2 equivalent
			 * on the Atom, but is generally faster on other processors supporting
			 * SSSE3. This flag allows for selectively disabling certain SSSE3
			 * functions on the Atom. */
			if family == 6 && model == 28 {
				rval |= ATOM
			}
		}
	}
	return Flags(rval)
}

func valAsString(values ...uint32) []byte {
	r := make([]byte, 4*len(values))
	for i, v := range values {
		dst := r[i*4:]
		dst[0] = byte(v & 0xff)
		dst[1] = byte((v >> 8) & 0xff)
		dst[2] = byte((v >> 16) & 0xff)
		dst[3] = byte((v >> 24) & 0xff)
		switch {
		case dst[0] == 0:
			return r[:i*4]
		case r[1] == 0:
			return r[:i*4+1]
		case r[2] == 0:
			return r[:i*4+2]
		case r[3] == 0:
			return r[:i*4+3]
		}
	}
	return r
}
