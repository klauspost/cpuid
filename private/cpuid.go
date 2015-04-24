// Generated, DO NOT EDIT,
// but copy it to your own project and rename the package.
// See more at http://github.com/klauspost/cpuid

package cpuid

import "strings"

// Vendor is a representation of a CPU vendor.
type vendor int

const (
	other	vendor	= iota
	intel
	amd
)

const (
	cmov		= 1 << iota	// i686 CMOV
	nx				// NX (No-Execute) bit
	amd3dnow			// AMD 3DNOW
	amd3dnowext			// AMD 3DNowExt
	mmx				// standard MMX
	mmxext				// SSE integer functions or AMD MMX ext
	sse				// SSE functions
	sse2				// P4 SSE functions
	sse3				// Prescott SSE3 functions
	ssse3				// Conroe SSSE3 functions
	sse4				// Penryn SSE4.1 functions
	sse4a				// AMD Barcelona microarchitecture SSE4a instructions
	sse42				// Nehalem SSE4.2 functions
	avx				// AVX functions
	avx2				// AVX2 functions
	fma3				// Intel FMA 3
	fma4				// Bulldozer FMA4 functions
	xop				// Bulldozer XOP functions
	f16c				// Half-precision floating-point conversion
	bmi1				// Bit Manipulation Instruction Set 1
	bmi2				// Bit Manipulation Instruction Set 2
	tbm				// AMD Trailing Bit Manipulation
	lzcnt				// LZCNT instruction
	aesni				// Advanced Encryption Standard New Instructions
	clmul				// Carry-less Multiplication
	htt				// Hyperthreading (enabled)
	hle				// Hardware Lock Elision
	rtm				// Restricted Transactional Memory
	rdrand				// RDRAND instruction is available
	rdseed				// RDSEED instruction is available
	adx				// Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
	sha				// Intel SHA Extensions
	avx512f				// AVX-512 Foundation
	avx512dq			// AVX-512 Doubleword and Quadword Instructions
	avx512ifma			// AVX-512 Integer Fused Multiply-Add Instructions
	avx512pf			// AVX-512 Prefetch Instructions
	avx512er			// AVX-512 Exponential and Reciprocal Instructions
	avx512cd			// AVX-512 Conflict Detection Instructions
	avx512bw			// AVX-512 Byte and Word Instructions
	avx512vl			// AVX-512 Vector Length Extensions
	avx512vbmi			// AVX-512 Vector Bit Manipulation Instructions
	mpx				// Intel MPX (Memory Protection Extensions)
	erms				// Enhanced REP MOVSB/STOSB

	// Performance indicators
	sse2slow	// SSE2 is supported, but usually not faster
	sse3slow	// SSE3 is supported, but usually not faster
	atom		// Atom processor, some SSSE3 instructions are slower
)

var flagNames = map[flags]string{
	cmov:		"CMOV",		// i686 CMOV
	nx:		"NX",		// NX (No-Execute) bit
	amd3dnow:	"AMD3DNOW",	// AMD 3DNOW
	amd3dnowext:	"AMD3DNOWEXT",	// AMD 3DNowExt
	mmx:		"MMX",		// Standard MMX
	mmxext:		"MMXEXT",	// SSE integer functions or AMD MMX ext
	sse:		"SSE",		// SSE functions
	sse2:		"SSE2",		// P4 SSE2 functions
	sse3:		"SSE3",		// Prescott SSE3 functions
	ssse3:		"SSSE3",	// Conroe SSSE3 functions
	sse4:		"SSE4.1",	// Penryn SSE4.1 functions
	sse42:		"SSE4.2",	// Nehalem SSE4.2 functions
	avx:		"AVX",		// AVX functions
	avx2:		"AVX2",		// AVX functions
	fma3:		"FMA3",		// Intel FMA 3
	fma4:		"FMA4",		// Bulldozer FMA4 functions
	xop:		"XOP",		// Bulldozer XOP functions
	f16c:		"F16C",		// Half-precision floating-point conversion
	bmi1:		"BMI1",		// Bit Manipulation Instruction Set 1
	bmi2:		"BMI2",		// Bit Manipulation Instruction Set 2
	tbm:		"TBM",		// AMD Trailing Bit Manipulation
	lzcnt:		"LZCNT",	// LZCNT instruction
	aesni:		"AESNI",	// Advanced Encryption Standard New Instructions
	clmul:		"CLMUL",	// Carry-less Multiplication
	htt:		"HTT",		// Hyperthreading (enabled)
	hle:		"HLE",		// Hardware Lock Elision
	rtm:		"RTM",		// Restricted Transactional Memory
	rdrand:		"RDRAND",	// RDRAND instruction is available
	rdseed:		"RDSEED",	// RDSEED instruction is available
	adx:		"ADX",		// Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
	sha:		"SHA",		// Intel SHA Extensions
	avx512f:	"AVX512F",	// AVX-512 Foundation
	avx512dq:	"AVX512DQ",	// AVX-512 Doubleword and Quadword Instructions
	avx512ifma:	"AVX512IFMA",	// AVX-512 Integer Fused Multiply-Add Instructions
	avx512pf:	"AVX512PF",	// AVX-512 Prefetch Instructions
	avx512er:	"AVX512ER",	// AVX-512 Exponential and Reciprocal Instructions
	avx512cd:	"AVX512CD",	// AVX-512 Conflict Detection Instructions
	avx512bw:	"AVX512BW",	// AVX-512 Byte and Word Instructions
	avx512vl:	"AVX512VL",	// AVX-512 Vector Length Extensions
	avx512vbmi:	"AVX512VBMI",	// AVX-512 Vector Bit Manipulation Instructions
	mpx:		"MPX",		// Intel MPX (Memory Protection Extensions)
	erms:		"ERMS",		// Enhanced REP MOVSB/STOSB

	// Performance indicators
	sse2slow:	"SSE2SLOW",	// SSE2 supported, but usually not faster
	sse3slow:	"SSE3SLOW",	// SSE3 supported, but usually not faster
	atom:		"ATOM",		// Atom processor, some SSSE3 instructions are slower

}

// CPUInfo contains information about the detected system CPU.
type cpuInfo struct {
	brandname	string	// Brand name reported by the CPU
	vendorid	vendor	// Comparable CPU vendor ID
	features	flags	// Features of the CPU
	physicalcores	int	// Number of physical processor cores in your CPU. Will be 0 if undetectable.
	threadspercore	int	// Number of threads per physical core. Will be 1 if undetectable.
	logicalcores	int	// Number of physical cores times threads that can run on each core through the use of hyperthreading. Will be 0 if undetectable.
	family		int	// CPU family number
	model		int	// CPU model number
	cacheline	int	// Cache line size in bytes. Will be 0 if undetectable.
}

// CPU contains information about the CPU as detected on startup,
// or when Detect last was called.
//
// Use this as the primary entry point to you data,
// this way queries are
var cpu cpuInfo

func init() {
	detect()
}

// Detect will re-detect current CPU info.
// This will replace the content of the exported CPU variable.
//
// Unless you expect the CPU to change while you are running your program
// you should not need to call this function.
// If you call this, you must ensure that no other goroutine is accessing the
// exported CPU variable.
func detect() {
	cpu.brandname = brandName()
	cpu.cacheline = cacheLine()
	cpu.family, cpu.model = familyModel()
	cpu.features = support()
	cpu.threadspercore = threadsPerCore()
	cpu.logicalcores = logicalCores()
	cpu.physicalcores = physicalCores()
	cpu.vendorid = vendorID()
}

// Generated here: http://play.golang.org/p/BxFH2Gdc0G

// Cmov indicates support of CMOV instructions
func (c cpuInfo) cmov() bool {
	return c.features&cmov != 0
}

// Amd3dnow indicates support of AMD 3DNOW! instructions
func (c cpuInfo) amd3dnow() bool {
	return c.features&amd3dnow != 0
}

// Amd3dnowExt indicates support of AMD 3DNOW! Extended instructions
func (c cpuInfo) amd3dnowext() bool {
	return c.features&amd3dnowext != 0
}

// MMX indicates support of MMX instructions
func (c cpuInfo) mmx() bool {
	return c.features&mmx != 0
}

// MMXExt indicates support of MMXEXT instructions
// (SSE integer functions or AMD MMX ext)
func (c cpuInfo) mmxext() bool {
	return c.features&mmxext != 0
}

// SSE indicates support of SSE instructions
func (c cpuInfo) sse() bool {
	return c.features&sse != 0
}

// SSE2 indicates support of SSE 2 instructions
func (c cpuInfo) sse2() bool {
	return c.features&sse2 != 0
}

// SSE3 indicates support of SSE 3 instructions
func (c cpuInfo) sse3() bool {
	return c.features&sse3 != 0
}

// SSSE3 indicates support of SSSE 3 instructions
func (c cpuInfo) ssse3() bool {
	return c.features&ssse3 != 0
}

// SSE4 indicates support of SSE 4 (also called SSE 4.1) instructions
func (c cpuInfo) sse4() bool {
	return c.features&sse4 != 0
}

// SSE42 indicates support of SSE4.2 instructions
func (c cpuInfo) sse42() bool {
	return c.features&sse42 != 0
}

// AVX indicates support of AVX instructions
// and operating system support of AVX instructions
func (c cpuInfo) avx() bool {
	return c.features&avx != 0
}

// AVX2 indicates support of AVX2 instructions
func (c cpuInfo) avx2() bool {
	return c.features&avx2 != 0
}

// FMA3 indicates support of FMA3 instructions
func (c cpuInfo) fma3() bool {
	return c.features&fma3 != 0
}

// FMA4 indicates support of FMA4 instructions
func (c cpuInfo) fma4() bool {
	return c.features&fma4 != 0
}

// XOP indicates support of XOP instructions
func (c cpuInfo) xop() bool {
	return c.features&xop != 0
}

// F16C indicates support of F16C instructions
func (c cpuInfo) f16c() bool {
	return c.features&f16c != 0
}

// BMI1 indicates support of BMI1 instructions
func (c cpuInfo) bmi1() bool {
	return c.features&bmi1 != 0
}

// BMI2 indicates support of BMI2 instructions
func (c cpuInfo) bmi2() bool {
	return c.features&bmi2 != 0
}

// TBM indicates support of TBM instructions
// (AMD Trailing Bit Manipulation)
func (c cpuInfo) tbm() bool {
	return c.features&tbm != 0
}

// Lzcnt indicates support of LZCNT instruction
func (c cpuInfo) lzcnt() bool {
	return c.features&lzcnt != 0
}

// HTT indicates the processor has Hyperthreading enabled
func (c cpuInfo) htt() bool {
	return c.features&htt != 0
}

// SSE2Slow indicates that SSE2 may be slow on this processor
func (c cpuInfo) sse2slow() bool {
	return c.features&sse2slow != 0
}

// SSE3Slow indicates that SSE3 may be slow on this processor
func (c cpuInfo) sse3slow() bool {
	return c.features&sse3slow != 0
}

// AesNi indicates support of AES-NI instructions
// (Advanced Encryption Standard New Instructions)
func (c cpuInfo) aesni() bool {
	return c.features&aesni != 0
}

// Clmul indicates support of CLMUL instructions
// (Carry-less Multiplication)
func (c cpuInfo) clmul() bool {
	return c.features&clmul != 0
}

// NX indicates support of NX (No-Execute) bit
func (c cpuInfo) nx() bool {
	return c.features&nx != 0
}

// SSE4A indicates support of AMD Barcelona microarchitecture SSE4a instructions
func (c cpuInfo) sse4a() bool {
	return c.features&sse4a != 0
}

// HLE indicates support of Hardware Lock Elision
func (c cpuInfo) hle() bool {
	return c.features&hle != 0
}

// RTM indicates support of Restricted Transactional Memory
func (c cpuInfo) rtm() bool {
	return c.features&rtm != 0
}

// Rdrand indicates support of RDRAND instruction is available
func (c cpuInfo) rdrand() bool {
	return c.features&rdrand != 0
}

// Rdseed indicates support of RDSEED instruction is available
func (c cpuInfo) rdseed() bool {
	return c.features&rdseed != 0
}

// ADX indicates support of Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
func (c cpuInfo) adx() bool {
	return c.features&adx != 0
}

// SHA indicates support of Intel SHA Extensions
func (c cpuInfo) sha() bool {
	return c.features&sha != 0
}

// AVX512F indicates support of AVX-512 Foundation
func (c cpuInfo) avx512f() bool {
	return c.features&avx512f != 0
}

// AVX512DQ indicates support of AVX-512 Doubleword and Quadword Instructions
func (c cpuInfo) avx512dq() bool {
	return c.features&avx512dq != 0
}

// AVX512IFMA indicates support of AVX-512 Integer Fused Multiply-Add Instructions
func (c cpuInfo) avx512ifma() bool {
	return c.features&avx512ifma != 0
}

// AVX512PF indicates support of AVX-512 Prefetch Instructions
func (c cpuInfo) avx512pf() bool {
	return c.features&avx512pf != 0
}

// AVX512ER indicates support of AVX-512 Exponential and Reciprocal Instructions
func (c cpuInfo) avx512er() bool {
	return c.features&avx512er != 0
}

// AVX512CD indicates support of AVX-512 Conflict Detection Instructions
func (c cpuInfo) avx512cd() bool {
	return c.features&avx512cd != 0
}

// AVX512BW indicates support of AVX-512 Byte and Word Instructions
func (c cpuInfo) avx512bw() bool {
	return c.features&avx512bw != 0
}

// AVX512VL indicates support of AVX-512 Vector Length Extensions
func (c cpuInfo) avx512vl() bool {
	return c.features&avx512vl != 0
}

// AVX512VBMI indicates support of AVX-512 Vector Bit Manipulation Instructions
func (c cpuInfo) avx512vbmi() bool {
	return c.features&avx512vbmi != 0
}

// MPX indicates support of Intel MPX (Memory Protection Extensions)
func (c cpuInfo) mpx() bool {
	return c.features&mpx != 0
}

// ERMS indicates support of Enhanced REP MOVSB/STOSB
func (c cpuInfo) erms() bool {
	return c.features&erms != 0
}

// Atom indicates an Atom processor
func (c cpuInfo) atom() bool {
	return c.features&atom != 0
}

// Intel returns true if vendor is recognized as Intel
func (c cpuInfo) intel() bool {
	return c.vendorid == intel
}

// AMD returns true if vendor is recognized as AMD
func (c cpuInfo) amd() bool {
	return c.vendorid == amd
}

// Flags contains detected cpu features and caracteristics
type flags uint64

// String returns a string representation of the detected
// CPU features.
func (f flags) String() string {
	return strings.Join(f.strings(), ",")
}

// Strings returns and array of the detected features.
func (f flags) strings() []string {
	s := support()
	r := make([]string, 0, 20)
	for i := uint(0); i < 64; i++ {
		key := flags(1 << i)
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
	case intel:
		// Use this on old Intel processors
		if maxFunctionID() < 0xb {
			// CPUID.1:EBX[23:16] represents the maximum number of addressable IDs (initial APIC ID)
			// that can be assigned to logical processors in a physical package.
			// The value may not be the same as the number of logical processors that are present in the hardware of a physical package.
			_, ebx, _, _ := cpuid(1)
			ebx <<= 16
			if ebx != 0 {
				return int(ebx & 0xff)
			}
			return 0
		}
		_, b, _, _ := cpuidex(0xb, 1)
		return int(b & 0xffff)
	case amd:
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
	case intel:
		return logicalCores() / threadsPerCore()
	case amd:
		_, _, c, _ := cpuid(0x80000008)
		return int(c&0xff) + 1
	default:
		return 0
	}
}

func vendorID() vendor {
	_, b, c, d := cpuid(0)
	if b == 0x756e6547 && d == 0x49656e69 && c == 0x6c65746e {
		return intel
	} else if b == 0x68747541 && d == 0x69746e65 && c == 0x444d4163 {
		return amd
	} else {
		return other
	}
}

func cacheLine() int {
	if maxFunctionID() < 0x1 {
		return 0
	}

	_, ebx, _, _ := cpuid(1)
	cache := (ebx & 0xff00) >> 5	// cflush size
	if cache == 0 && maxExtendedFunction() >= 0x80000006 {
		_, _, ecx, _ := cpuid(0x80000006)
		cache = ecx & 0xff	// cacheline size
	}
	// TODO: Read from Cache and TLB Information
	return int(cache)
}

func support() flags {
	mfi := maxFunctionID()
	if mfi < 0x1 {
		return 0
	}
	rval := uint64(0)
	_, b, c, d := cpuid(1)
	if (d & (1 << 15)) != 0 {
		rval |= cmov
	}
	if (d & (1 << 23)) != 0 {
		rval |= mmx
	}
	if (d & (1 << 25)) != 0 {
		rval |= mmxext
	}
	if (d & (1 << 25)) != 0 {
		rval |= sse
	}
	if (d & (1 << 26)) != 0 {
		rval |= sse2
	}
	if (c & 1) != 0 {
		rval |= sse3
	}
	if (c & 0x00000200) != 0 {
		rval |= ssse3
	}
	if (c & 0x00080000) != 0 {
		rval |= sse4
	}
	if (c & 0x00100000) != 0 {
		rval |= sse42
	}
	if (c & (1 << 25)) != 0 {
		rval |= aesni
	}
	if (c & (1 << 1)) != 0 {
		rval |= clmul
	}
	if c&(1<<30) != 0 {
		rval |= rdrand
	}
	if c&(1<<29) != 0 {
		rval |= f16c
	}
	if (c & (1 << 28)) != 0 {
		// This field does not indicate that Hyper-Threading
		// Technology has been enabled for this specific processor.
		// To determine if Hyper-Threading Technology is supported,
		// check the value returned in EBX[23:16]
		v := (b >> 16) & 255
		if v > 0 {
			rval |= htt
		}
	}

	// Check OXSAVE and AVX bits
	if (c & 0x18000000) == 0x18000000 {
		// Check for OS support
		eax, _ := xgetbv(0)
		if (eax & 0x6) == 0x6 {
			rval |= avx
			if (c & 0x00001000) != 0 {
				rval |= fma3
			}
		}
	}

	// Check AVX2, AVX2 requires OS support, but BMI1/2 don't.
	if mfi >= 7 {
		_, ebx, ecx, _ := cpuidex(7, 0)
		if (rval&avx) != 0 && (ebx&0x00000020) != 0 {
			rval |= avx2
		}
		if (ebx & 0x00000008) != 0 {
			rval |= bmi1
			if (ebx & 0x00000100) != 0 {
				rval |= bmi2
			}
		}
		if ebx&(1<<4) != 0 {
			rval |= hle
		}
		if ebx&(1<<9) != 0 {
			rval |= erms
		}
		if ebx&(1<<11) != 0 {
			rval |= rtm
		}
		if ebx&(1<<14) != 0 {
			rval |= mpx
		}
		if ebx&(1<<18) != 0 {
			rval |= rdseed
		}
		if ebx&(1<<19) != 0 {
			rval |= adx
		}
		if ebx&(1<<29) != 0 {
			rval |= sha
		}

		// Only detect AVX-512 features if XGETBV is supported
		if c&((1<<26)|(1<<27)) == (1<<26)|(1<<27) {
			// Check for OS support
			eax, _ := xgetbv(0)

			// Verify that XCR0[7:5] = ‘111b’ (OPMASK state, upper 256-bit of ZMM0-ZMM15 and
			// ZMM16-ZMM31 state are enabled by OS)
			/// and that XCR0[2:1] = ‘11b’ (XMM state and YMM state are enabled by OS).
			if (eax>>5)&7 == 7 && (eax>>1)&3 == 3 {
				if ebx&(1<<16) != 0 {
					rval |= avx512f
				}
				if ebx&(1<<17) != 0 {
					rval |= avx512dq
				}
				if ebx&(1<<21) != 0 {
					rval |= avx512ifma
				}
				if ebx&(1<<26) != 0 {
					rval |= avx512pf
				}
				if ebx&(1<<27) != 0 {
					rval |= avx512er
				}
				if ebx&(1<<28) != 0 {
					rval |= avx512cd
				}
				if ebx&(1<<30) != 0 {
					rval |= avx512bw
				}
				if ebx&(1<<31) != 0 {
					rval |= avx512vl
				}
				// ecx
				if ecx&(1<<1) != 0 {
					rval |= avx512vbmi
				}
			}
		}
	}

	if maxExtendedFunction() >= 0x80000001 {
		_, _, c, d := cpuid(0x80000001)
		if (c & 0x00000020) != 0 {
			rval |= lzcnt
		}
		if (d & (1 << 31)) != 0 {
			rval |= amd3dnow
		}
		if (d & (1 << 30)) != 0 {
			rval |= amd3dnowext
		}
		if (d & (1 << 23)) != 0 {
			rval |= mmx
		}
		if (d & (1 << 22)) != 0 {
			rval |= mmxext
		}
		if (c & (1 << 6)) != 0 {
			rval |= sse4a
		}
		if d&(1<<20) != 0 {
			rval |= nx
		}

		/* Allow for selectively disabling SSE2 functions on AMD processors
		   with SSE2 support but not SSE4a. This includes Athlon64, some
		   Opteron, and some Sempron processors. MMX, SSE, or 3DNow! are faster
		   than SSE2 often enough to utilize this special-case flag.
		   AV_CPU_FLAG_SSE2 and AV_CPU_FLAG_SSE2SLOW are both set in this case
		   so that SSE2 is used unless explicitly disabled by checking
		   AV_CPU_FLAG_SSE2SLOW. */
		if vendorID() != intel &&
			rval&sse2 != 0 && (c&0x00000040) == 0 {
			rval |= sse2slow
		}

		/* XOP and FMA4 use the AVX instruction coding scheme, so they can't be
		 * used unless the OS has AVX support. */
		if (rval & avx) != 0 {
			if (c & 0x00000800) != 0 {
				rval |= xop
			}
			if (c & 0x00010000) != 0 {
				rval |= fma4
			}
		}

		if vendorID() == intel {
			family, model := familyModel()
			if family == 6 && (model == 9 || model == 13 || model == 14) {
				/* 6/9 (pentium-m "banias"), 6/13 (pentium-m "dothan"), and
				 * 6/14 (core1 "yonah") theoretically support sse2, but it's
				 * usually slower than mmx. */
				if (rval & sse2) != 0 {
					rval |= sse2slow
				}
				if (rval & sse3) != 0 {
					rval |= sse3slow
				}
			}
			/* The Atom processor has SSSE3 support, which is useful in many cases,
			 * but sometimes the SSSE3 version is slower than the SSE2 equivalent
			 * on the Atom, but is generally faster on other processors supporting
			 * SSSE3. This flag allows for selectively disabling certain SSSE3
			 * functions on the Atom. */
			if family == 6 && model == 28 {
				rval |= atom
			}
		}
	}
	return flags(rval)
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
