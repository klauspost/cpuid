// Generated, DO NOT EDIT,
// but copy it to your own project and rename the package.
// See more at http://github.com/klauspost/cpuid

package cpuid

import (
	"math"
	"strings"
)

// AMD refererence: https://www.amd.com/system/files/TechDocs/25481.pdf
// and Processor Programming Reference (PPR)

// Vendor is a representation of a CPU vendor.
type vendor int

const (
	other vendor = iota
	intel
	amd
	via
	transmeta
	nsc
	kvm  // Kernel-based Virtual Machine
	msvm // Microsoft Hyper-V or Windows Virtual PC
	vmware
	xenhvm
	bhyve
	hygon
	sis
	rdc
)

const (
	cmov               = 1 << iota // i686 CMOV
	nx                             // NX (No-Execute) bit
	amd3dnow                       // AMD 3DNOW
	amd3dnowext                    // AMD 3DNowExt
	mmx                            // standard MMX
	mmxext                         // SSE integer functions or AMD MMX ext
	sse                            // SSE functions
	sse2                           // P4 SSE functions
	sse3                           // Prescott SSE3 functions
	ssse3                          // Conroe SSSE3 functions
	sse4                           // Penryn SSE4.1 functions
	sse4a                          // AMD Barcelona microarchitecture SSE4a instructions
	sse42                          // Nehalem SSE4.2 functions
	avx                            // AVX functions
	avx2                           // AVX2 functions
	fma3                           // Intel FMA 3
	fma4                           // Bulldozer FMA4 functions
	xop                            // Bulldozer XOP functions
	f16c                           // Half-precision floating-point conversion
	bmi1                           // Bit Manipulation Instruction Set 1
	bmi2                           // Bit Manipulation Instruction Set 2
	tbm                            // AMD Trailing Bit Manipulation
	lzcnt                          // LZCNT instruction
	popcnt                         // POPCNT instruction
	aesni                          // Advanced Encryption Standard New Instructions
	clmul                          // Carry-less Multiplication
	htt                            // Hyperthreading (enabled)
	hle                            // Hardware Lock Elision
	rtm                            // Restricted Transactional Memory
	rdrand                         // RDRAND instruction is available
	rdseed                         // RDSEED instruction is available
	adx                            // Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
	sha                            // Intel SHA Extensions
	avx512f                        // AVX-512 Foundation
	avx512dq                       // AVX-512 Doubleword and Quadword Instructions
	avx512ifma                     // AVX-512 Integer Fused Multiply-Add Instructions
	avx512pf                       // AVX-512 Prefetch Instructions
	avx512er                       // AVX-512 Exponential and Reciprocal Instructions
	avx512cd                       // AVX-512 Conflict Detection Instructions
	avx512bw                       // AVX-512 Byte and Word Instructions
	avx512vl                       // AVX-512 Vector Length Extensions
	avx512vbmi                     // AVX-512 Vector Bit Manipulation Instructions
	avx512vbmi2                    // AVX-512 Vector Bit Manipulation Instructions, Version 2
	avx512vnni                     // AVX-512 Vector Neural Network Instructions
	avx512vpopcntdq                // AVX-512 Vector Population Count Doubleword and Quadword
	gfni                           // Galois Field New Instructions
	vaes                           // Vector AES
	avx512bitalg                   // AVX-512 Bit Algorithms
	vpclmulqdq                     // Carry-Less Multiplication Quadword
	avx512bf16                     // AVX-512 BFLOAT16 Instructions
	avx512vp2intersect             // AVX-512 Intersect for D/Q
	mpx                            // Intel MPX (Memory Protection Extensions)
	erms                           // Enhanced REP MOVSB/STOSB
	rdtscp                         // RDTSCP Instruction
	cx16                           // CMPXCHG16B Instruction
	sgx                            // Software Guard Extensions
	sgxlc                          // Software Guard Extensions Launch Control
	ibpb                           // Indirect Branch Restricted Speculation (IBRS) and Indirect Branch Predictor Barrier (IBPB)
	stibp                          // Single Thread Indirect Branch Predictors
	vmx                            // Virtual Machine Extensions

	// Performance indicators
	sse2slow // SSE2 is supported, but usually not faster
	sse3slow // SSE3 is supported, but usually not faster
	atom     // Atom processor, some SSSE3 instructions are slower
)

var flagNames = map[flags]string{
	cmov:               "CMOV",               // i686 CMOV
	nx:                 "NX",                 // NX (No-Execute) bit
	amd3dnow:           "AMD3DNOW",           // AMD 3DNOW
	amd3dnowext:        "AMD3DNOWEXT",        // AMD 3DNowExt
	mmx:                "MMX",                // Standard MMX
	mmxext:             "MMXEXT",             // SSE integer functions or AMD MMX ext
	sse:                "SSE",                // SSE functions
	sse2:               "SSE2",               // P4 SSE2 functions
	sse3:               "SSE3",               // Prescott SSE3 functions
	ssse3:              "SSSE3",              // Conroe SSSE3 functions
	sse4:               "SSE4.1",             // Penryn SSE4.1 functions
	sse4a:              "SSE4A",              // AMD Barcelona microarchitecture SSE4a instructions
	sse42:              "SSE4.2",             // Nehalem SSE4.2 functions
	avx:                "AVX",                // AVX functions
	avx2:               "AVX2",               // AVX functions
	fma3:               "FMA3",               // Intel FMA 3
	fma4:               "FMA4",               // Bulldozer FMA4 functions
	xop:                "XOP",                // Bulldozer XOP functions
	f16c:               "F16C",               // Half-precision floating-point conversion
	bmi1:               "BMI1",               // Bit Manipulation Instruction Set 1
	bmi2:               "BMI2",               // Bit Manipulation Instruction Set 2
	tbm:                "TBM",                // AMD Trailing Bit Manipulation
	lzcnt:              "LZCNT",              // LZCNT instruction
	popcnt:             "POPCNT",             // POPCNT instruction
	aesni:              "AESNI",              // Advanced Encryption Standard New Instructions
	clmul:              "CLMUL",              // Carry-less Multiplication
	htt:                "HTT",                // Hyperthreading (enabled)
	hle:                "HLE",                // Hardware Lock Elision
	rtm:                "RTM",                // Restricted Transactional Memory
	rdrand:             "RDRAND",             // RDRAND instruction is available
	rdseed:             "RDSEED",             // RDSEED instruction is available
	adx:                "ADX",                // Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
	sha:                "SHA",                // Intel SHA Extensions
	avx512f:            "AVX512F",            // AVX-512 Foundation
	avx512dq:           "AVX512DQ",           // AVX-512 Doubleword and Quadword Instructions
	avx512ifma:         "AVX512IFMA",         // AVX-512 Integer Fused Multiply-Add Instructions
	avx512pf:           "AVX512PF",           // AVX-512 Prefetch Instructions
	avx512er:           "AVX512ER",           // AVX-512 Exponential and Reciprocal Instructions
	avx512cd:           "AVX512CD",           // AVX-512 Conflict Detection Instructions
	avx512bw:           "AVX512BW",           // AVX-512 Byte and Word Instructions
	avx512vl:           "AVX512VL",           // AVX-512 Vector Length Extensions
	avx512vbmi:         "AVX512VBMI",         // AVX-512 Vector Bit Manipulation Instructions
	avx512vbmi2:        "AVX512VBMI2",        // AVX-512 Vector Bit Manipulation Instructions, Version 2
	avx512vnni:         "AVX512VNNI",         // AVX-512 Vector Neural Network Instructions
	avx512vpopcntdq:    "AVX512VPOPCNTDQ",    // AVX-512 Vector Population Count Doubleword and Quadword
	gfni:               "GFNI",               // Galois Field New Instructions
	vaes:               "VAES",               // Vector AES
	avx512bitalg:       "AVX512BITALG",       // AVX-512 Bit Algorithms
	vpclmulqdq:         "VPCLMULQDQ",         // Carry-Less Multiplication Quadword
	avx512bf16:         "AVX512BF16",         // AVX-512 BFLOAT16 Instruction
	avx512vp2intersect: "AVX512VP2INTERSECT", // AVX-512 Intersect for D/Q
	mpx:                "MPX",                // Intel MPX (Memory Protection Extensions)
	erms:               "ERMS",               // Enhanced REP MOVSB/STOSB
	rdtscp:             "RDTSCP",             // RDTSCP Instruction
	cx16:               "CX16",               // CMPXCHG16B Instruction
	sgx:                "SGX",                // Software Guard Extensions
	sgxlc:              "SGXLC",              // Software Guard Extensions Launch Control
	ibpb:               "IBPB",               // Indirect Branch Restricted Speculation and Indirect Branch Predictor Barrier
	stibp:              "STIBP",              // Single Thread Indirect Branch Predictors
	vmx:                "VMX",                // Virtual Machine Extensions

	// Performance indicators
	sse2slow: "SSE2SLOW", // SSE2 supported, but usually not faster
	sse3slow: "SSE3SLOW", // SSE3 supported, but usually not faster
	atom:     "ATOM",     // Atom processor, some SSSE3 instructions are slower

}

/* all special features for arm64 should be defined here */
const (
	/* extension instructions */
	arm_fp = 1 << iota
	arm_asimd
	arm_evtstrm
	arm_aes
	arm_pmull
	arm_sha1
	arm_sha2
	arm_crc32
	arm_atomics
	arm_fphp
	arm_asimdhp
	arm_cpuid
	arm_asimdrdm
	arm_jscvt
	arm_fcma
	arm_lrcpc
	arm_dcpop
	arm_sha3
	arm_sm3
	arm_sm4
	arm_asimddp
	arm_sha512
	arm_sve
	arm_gpa
)

var flagNamesArm = map[flags]string{
	arm_fp:       "FP",       // Single-precision and double-precision floating point
	arm_asimd:    "ASIMD",    // Advanced SIMD
	arm_evtstrm:  "EVTSTRM",  // Generic timer
	arm_aes:      "AES",      // AES instructions
	arm_pmull:    "PMULL",    // Polynomial Multiply instructions (PMULL/PMULL2)
	arm_sha1:     "SHA1",     // SHA-1 instructions (SHA1C, etc)
	arm_sha2:     "SHA2",     // SHA-2 instructions (SHA256H, etc)
	arm_crc32:    "CRC32",    // CRC32/CRC32C instructions
	arm_atomics:  "ATOMICS",  // Large System Extensions (LSE)
	arm_fphp:     "FPHP",     // Half-precision floating point
	arm_asimdhp:  "ASIMDHP",  // Advanced SIMD half-precision floating point
	arm_cpuid:    "CPUID",    // Some CPU ID registers readable at user-level
	arm_asimdrdm: "ASIMDRDM", // Rounding Double Multiply Accumulate/Subtract (SQRDMLAH/SQRDMLSH)
	arm_jscvt:    "JSCVT",    // Javascript-style double->int convert (FJCVTZS)
	arm_fcma:     "FCMA",     // Floatin point complex number addition and multiplication
	arm_lrcpc:    "LRCPC",    // Weaker release consistency (LDAPR, etc)
	arm_dcpop:    "DCPOP",    // Data cache clean to Point of Persistence (DC CVAP)
	arm_sha3:     "SHA3",     // SHA-3 instructions (EOR3, RAXI, XAR, BCAX)
	arm_sm3:      "SM3",      // SM3 instructions
	arm_sm4:      "SM4",      // SM4 instructions
	arm_asimddp:  "ASIMDDP",  // SIMD Dot Product
	arm_sha512:   "SHA512",   // SHA512 instructions
	arm_sve:      "SVE",      // Scalable Vector Extension
	arm_gpa:      "GPA",      // Generic Pointer Authentication
}

// CPUInfo contains information about the detected system CPU.
type cpuInfo struct {
	brandname      string // Brand name reported by the CPU
	vendorid       vendor // Comparable CPU vendor ID
	vendorstring   string // Raw vendor string.
	features       flags  // Features of the CPU (x64)
	arm            flags  // Features of the CPU (arm)
	physicalcores  int    // Number of physical processor cores in your CPU. Will be 0 if undetectable.
	threadspercore int    // Number of threads per physical core. Will be 1 if undetectable.
	logicalcores   int    // Number of physical cores times threads that can run on each core through the use of hyperthreading. Will be 0 if undetectable.
	family         int    // CPU family number
	model          int    // CPU model number
	cacheline      int    // Cache line size in bytes. Will be 0 if undetectable.
	hz             int64  // Clock speed, if known
	cache          struct {
		l1i int // L1 Instruction Cache (per core or shared). Will be -1 if undetected
		l1d int // L1 Data Cache (per core or shared). Will be -1 if undetected
		l2  int // L2 Cache (per core or shared). Will be -1 if undetected
		l3  int // L3 Cache (per core, per ccx or shared). Will be -1 if undetected
	}
	sgx       sgxsupport
	maxFunc   uint32
	maxExFunc uint32
}

var cpuid func(op uint32) (eax, ebx, ecx, edx uint32)
var cpuidex func(op, op2 uint32) (eax, ebx, ecx, edx uint32)
var xgetbv func(index uint32) (eax, edx uint32)
var rdtscpAsm func() (eax, ebx, ecx, edx uint32)

// CPU contains information about the CPU as detected on startup,
// or when Detect last was called.
//
// Use this as the primary entry point to you data,
// this way queries are
var cpu cpuInfo

func init() {
	initCPU()
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
	cpu.maxFunc = maxFunctionID()
	cpu.maxExFunc = maxExtendedFunction()
	cpu.brandname = brandName()
	cpu.cacheline = cacheLine()
	cpu.family, cpu.model = familyModel()
	cpu.features = support()
	cpu.arm = flags(supportArm())
	cpu.sgx = hasSGX(cpu.features&sgx != 0, cpu.features&sgxlc != 0)
	cpu.threadspercore = threadsPerCore()
	cpu.logicalcores = logicalCores()
	cpu.physicalcores = physicalCores()
	cpu.vendorid, cpu.vendorstring = vendorID()
	cpu.hz = hertz(cpu.brandname)
	cpu.cacheSize()
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

// VMX indicates support of VMX
func (c cpuInfo) vmx() bool {
	return c.features&vmx != 0
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

// Popcnt indicates support of POPCNT instruction
func (c cpuInfo) popcnt() bool {
	return c.features&popcnt != 0
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

// AVX512VBMI2 indicates support of AVX-512 Vector Bit Manipulation Instructions, Version 2
func (c cpuInfo) avx512vbmi2() bool {
	return c.features&avx512vbmi2 != 0
}

// AVX512VNNI indicates support of AVX-512 Vector Neural Network Instructions
func (c cpuInfo) avx512vnni() bool {
	return c.features&avx512vnni != 0
}

// AVX512VPOPCNTDQ indicates support of AVX-512 Vector Population Count Doubleword and Quadword
func (c cpuInfo) avx512vpopcntdq() bool {
	return c.features&avx512vpopcntdq != 0
}

// GFNI indicates support of Galois Field New Instructions
func (c cpuInfo) gfni() bool {
	return c.features&gfni != 0
}

// VAES indicates support of Vector AES
func (c cpuInfo) vaes() bool {
	return c.features&vaes != 0
}

// AVX512BITALG indicates support of AVX-512 Bit Algorithms
func (c cpuInfo) avx512bitalg() bool {
	return c.features&avx512bitalg != 0
}

// VPCLMULQDQ indicates support of Carry-Less Multiplication Quadword
func (c cpuInfo) vpclmulqdq() bool {
	return c.features&vpclmulqdq != 0
}

// AVX512BF16 indicates support of
func (c cpuInfo) avx512bf16() bool {
	return c.features&avx512bf16 != 0
}

// AVX512VP2INTERSECT indicates support of
func (c cpuInfo) avx512vp2intersect() bool {
	return c.features&avx512vp2intersect != 0
}

// MPX indicates support of Intel MPX (Memory Protection Extensions)
func (c cpuInfo) mpx() bool {
	return c.features&mpx != 0
}

// ERMS indicates support of Enhanced REP MOVSB/STOSB
func (c cpuInfo) erms() bool {
	return c.features&erms != 0
}

// RDTSCP Instruction is available.
func (c cpuInfo) rdtscp() bool {
	return c.features&rdtscp != 0
}

// CX16 indicates if CMPXCHG16B instruction is available.
func (c cpuInfo) cx16() bool {
	return c.features&cx16 != 0
}

// TSX is split into HLE (Hardware Lock Elision) and RTM (Restricted Transactional Memory) detection.
// So TSX simply checks that.
func (c cpuInfo) tsx() bool {
	return c.features&(hle|rtm) == hle|rtm
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

// Hygon returns true if vendor is recognized as Hygon
func (c cpuInfo) hygon() bool {
	return c.vendorid == hygon
}

// Transmeta returns true if vendor is recognized as Transmeta
func (c cpuInfo) transmeta() bool {
	return c.vendorid == transmeta
}

// NSC returns true if vendor is recognized as National Semiconductor
func (c cpuInfo) nsc() bool {
	return c.vendorid == nsc
}

// VIA returns true if vendor is recognized as VIA
func (c cpuInfo) via() bool {
	return c.vendorid == via
}

// RTCounter returns the 64-bit time-stamp counter
// Uses the RDTSCP instruction. The value 0 is returned
// if the CPU does not support the instruction.
func (c cpuInfo) rtcounter() uint64 {
	if !c.rdtscp() {
		return 0
	}
	a, _, _, d := rdtscpAsm()
	return uint64(a) | (uint64(d) << 32)
}

// Ia32TscAux returns the IA32_TSC_AUX part of the RDTSCP.
// This variable is OS dependent, but on Linux contains information
// about the current cpu/core the code is running on.
// If the RDTSCP instruction isn't supported on the CPU, the value 0 is returned.
func (c cpuInfo) ia32tscaux() uint32 {
	if !c.rdtscp() {
		return 0
	}
	_, _, ecx, _ := rdtscpAsm()
	return ecx
}

// LogicalCPU will return the Logical CPU the code is currently executing on.
// This is likely to change when the OS re-schedules the running thread
// to another CPU.
// If the current core cannot be detected, -1 will be returned.
func (c cpuInfo) logicalcpu() int {
	if c.maxFunc < 1 {
		return -1
	}
	_, ebx, _, _ := cpuid(1)
	return int(ebx >> 24)
}

// hertz tries to compute the clock speed of the CPU. If leaf 15 is
// supported, use it, otherwise parse the brand string. Yes, really.
func hertz(model string) int64 {
	mfi := maxFunctionID()
	if mfi >= 0x15 {
		eax, ebx, ecx, _ := cpuid(0x15)
		if eax != 0 && ebx != 0 && ecx != 0 {
			return int64((int64(ecx) * int64(ebx)) / int64(eax))
		}
	}
	// computeHz determines the official rated speed of a CPU from its brand
	// string. This insanity is *actually the official documented way to do
	// this according to Intel*, prior to leaf 0x15 existing. The official
	// documentation only shows this working for exactly `x.xx` or `xxxx`
	// cases, e.g., `2.50GHz` or `1300MHz`; this parser will accept other
	// sizes.
	hz := strings.LastIndex(model, "Hz")
	if hz < 3 {
		return -1
	}
	var multiplier int64
	switch model[hz-1] {
	case 'M':
		multiplier = 1000 * 1000
	case 'G':
		multiplier = 1000 * 1000 * 1000
	case 'T':
		multiplier = 1000 * 1000 * 1000 * 1000
	}
	if multiplier == 0 {
		return -1
	}
	freq := int64(0)
	divisor := int64(0)
	decimalShift := int64(1)
	var i int
	for i = hz - 2; i >= 0 && model[i] != ' '; i-- {
		if model[i] >= '0' && model[i] <= '9' {
			freq += int64(model[i]-'0') * decimalShift
			decimalShift *= 10
		} else if model[i] == '.' {
			if divisor != 0 {
				return -1
			}
			divisor = decimalShift
		} else {
			return -1
		}
	}
	// we didn't find a space
	if i < 0 {
		return -1
	}
	if divisor != 0 {
		return (freq * multiplier) / divisor
	}
	return freq * multiplier
}

// VM Will return true if the cpu id indicates we are in
// a virtual machine. This is only a hint, and will very likely
// have many false negatives.
func (c cpuInfo) vm() bool {
	switch c.vendorid {
	case msvm, kvm, vmware, xenhvm, bhyve:
		return true
	}
	return false
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
		v := make([]uint32, 0, 48)
		for i := uint32(0); i < 3; i++ {
			a, b, c, d := cpuid(0x80000002 + i)
			v = append(v, a, b, c, d)
		}
		return strings.Trim(string(valAsString(v...)), " ")
	}
	return "unknown"
}

func threadsPerCore() int {
	mfi := maxFunctionID()
	vend, _ := vendorID()

	if mfi < 0x4 || (vend != intel && vend != amd) {
		return 1
	}

	if mfi < 0xb {
		if vend != intel {
			return 1
		}
		_, b, _, d := cpuid(1)
		if (d & (1 << 28)) != 0 {
			// v will contain logical core count
			v := (b >> 16) & 255
			if v > 1 {
				a4, _, _, _ := cpuid(4)
				// physical cores
				v2 := (a4 >> 26) + 1
				if v2 > 0 {
					return int(v) / int(v2)
				}
			}
		}
		return 1
	}
	_, b, _, _ := cpuidex(0xb, 0)
	if b&0xffff == 0 {
		return 1
	}
	return int(b & 0xffff)
}

func logicalCores() int {
	mfi := maxFunctionID()
	v, _ := vendorID()
	switch v {
	case intel:
		// Use this on old Intel processors
		if mfi < 0xb {
			if mfi < 1 {
				return 0
			}
			// CPUID.1:EBX[23:16] represents the maximum number of addressable IDs (initial APIC ID)
			// that can be assigned to logical processors in a physical package.
			// The value may not be the same as the number of logical processors that are present in the hardware of a physical package.
			_, ebx, _, _ := cpuid(1)
			logical := (ebx >> 16) & 0xff
			return int(logical)
		}
		_, b, _, _ := cpuidex(0xb, 1)
		return int(b & 0xffff)
	case amd, hygon:
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
	v, _ := vendorID()
	switch v {
	case intel:
		return logicalCores() / threadsPerCore()
	case amd, hygon:
		lc := logicalCores()
		tpc := threadsPerCore()
		if lc > 0 && tpc > 0 {
			return lc / tpc
		}
		// The following is inaccurate on AMD EPYC 7742 64-Core Processor

		if maxExtendedFunction() >= 0x80000008 {
			_, _, c, _ := cpuid(0x80000008)
			return int(c&0xff) + 1
		}
	}
	return 0
}

// Except from http://en.wikipedia.org/wiki/CPUID#EAX.3D0:_Get_vendor_ID
var vendorMapping = map[string]vendor{
	"AMDisbetter!": amd,
	"AuthenticAMD": amd,
	"CentaurHauls": via,
	"GenuineIntel": intel,
	"TransmetaCPU": transmeta,
	"GenuineTMx86": transmeta,
	"Geode by NSC": nsc,
	"VIA VIA VIA ": via,
	"KVMKVMKVMKVM": kvm,
	"Microsoft Hv": msvm,
	"VMwareVMware": vmware,
	"XenVMMXenVMM": xenhvm,
	"bhyve bhyve ": bhyve,
	"HygonGenuine": hygon,
	"Vortex86 SoC": sis,
	"SiS SiS SiS ": sis,
	"RiseRiseRise": sis,
	"Genuine  RDC": rdc,
}

func vendorID() (vendor, string) {
	_, b, c, d := cpuid(0)
	v := string(valAsString(b, d, c))
	vend, ok := vendorMapping[v]
	if !ok {
		return other, v
	}
	return vend, v
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

func (c *cpuInfo) cacheSize() {
	c.cache.l1d = -1
	c.cache.l1i = -1
	c.cache.l2 = -1
	c.cache.l3 = -1
	vendor, _ := vendorID()
	switch vendor {
	case intel:
		if maxFunctionID() < 4 {
			return
		}
		for i := uint32(0); ; i++ {
			eax, ebx, ecx, _ := cpuidex(4, i)
			cacheType := eax & 15
			if cacheType == 0 {
				break
			}
			cacheLevel := (eax >> 5) & 7
			coherency := int(ebx&0xfff) + 1
			partitions := int((ebx>>12)&0x3ff) + 1
			associativity := int((ebx>>22)&0x3ff) + 1
			sets := int(ecx) + 1
			size := associativity * partitions * coherency * sets
			switch cacheLevel {
			case 1:
				if cacheType == 1 {
					// 1 = Data Cache
					c.cache.l1d = size
				} else if cacheType == 2 {
					// 2 = Instruction Cache
					c.cache.l1i = size
				} else {
					if c.cache.l1d < 0 {
						c.cache.l1i = size
					}
					if c.cache.l1i < 0 {
						c.cache.l1i = size
					}
				}
			case 2:
				c.cache.l2 = size
			case 3:
				c.cache.l3 = size
			}
		}
	case amd, hygon:
		// Untested.
		if maxExtendedFunction() < 0x80000005 {
			return
		}
		_, _, ecx, edx := cpuid(0x80000005)
		c.cache.l1d = int(((ecx >> 24) & 0xFF) * 1024)
		c.cache.l1i = int(((edx >> 24) & 0xFF) * 1024)

		if maxExtendedFunction() < 0x80000006 {
			return
		}
		_, _, ecx, _ = cpuid(0x80000006)
		c.cache.l2 = int(((ecx >> 16) & 0xFFFF) * 1024)

		// CPUID Fn8000_001D_EAX_x[N:0] Cache Properties
		if maxExtendedFunction() < 0x8000001D {
			return
		}
		for i := uint32(0); i < math.MaxUint32; i++ {
			eax, ebx, ecx, _ := cpuidex(0x8000001D, i)

			level := (eax >> 5) & 7
			cacheNumSets := ecx + 1
			cacheLineSize := 1 + (ebx & 2047)
			cachePhysPartitions := 1 + ((ebx >> 12) & 511)
			cacheNumWays := 1 + ((ebx >> 22) & 511)

			typ := eax & 15
			size := int(cacheNumSets * cacheLineSize * cachePhysPartitions * cacheNumWays)
			if typ == 0 {
				return
			}

			switch level {
			case 1:
				switch typ {
				case 1:
					// Data cache
					c.cache.l1d = size
				case 2:
					// Inst cache
					c.cache.l1i = size
				default:
					if c.cache.l1d < 0 {
						c.cache.l1i = size
					}
					if c.cache.l1i < 0 {
						c.cache.l1i = size
					}
				}
			case 2:
				c.cache.l2 = size
			case 3:
				c.cache.l3 = size
			}
		}
	}

	return
}

type sgxepcsection struct {
	baseaddress uint64
	epcsize     uint64
}

type sgxsupport struct {
	available           bool
	launchcontrol       bool
	sgx1supported       bool
	sgx2supported       bool
	maxenclavesizenot64 int64
	maxenclavesize64    int64
	epcsections         []sgxepcsection
}

func hasSGX(available, lc bool) (rval sgxsupport) {
	rval.available = available

	if !available {
		return
	}

	rval.launchcontrol = lc

	a, _, _, d := cpuidex(0x12, 0)
	rval.sgx1supported = a&0x01 != 0
	rval.sgx2supported = a&0x02 != 0
	rval.maxenclavesizenot64 = 1 << (d & 0xFF)     // pow 2
	rval.maxenclavesize64 = 1 << ((d >> 8) & 0xFF) // pow 2
	rval.epcsections = make([]sgxepcsection, 0)

	for subleaf := uint32(2); subleaf < 2+8; subleaf++ {
		eax, ebx, ecx, edx := cpuidex(0x12, subleaf)
		leafType := eax & 0xf

		if leafType == 0 {
			// Invalid subleaf, stop iterating
			break
		} else if leafType == 1 {
			// EPC Section subleaf
			baseAddress := uint64(eax&0xfffff000) + (uint64(ebx&0x000fffff) << 32)
			size := uint64(ecx&0xfffff000) + (uint64(edx&0x000fffff) << 32)

			section := sgxepcsection{baseaddress: baseAddress, epcsize: size}
			rval.epcsections = append(rval.epcsections, section)
		}
	}

	return
}

func support() flags {
	mfi := maxFunctionID()
	vend, _ := vendorID()
	if mfi < 0x1 {
		return 0
	}
	rval := uint64(0)
	_, _, c, d := cpuid(1)
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
	if (c & (1 << 5)) != 0 {
		rval |= vmx
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
	if c&(1<<23) != 0 {
		rval |= popcnt
	}
	if c&(1<<30) != 0 {
		rval |= rdrand
	}
	if c&(1<<29) != 0 {
		rval |= f16c
	}
	if c&(1<<13) != 0 {
		rval |= cx16
	}
	if vend == intel && (d&(1<<28)) != 0 && mfi >= 4 {
		if threadsPerCore() > 1 {
			rval |= htt
		}
	}
	if vend == amd && (d&(1<<28)) != 0 && mfi >= 4 {
		if threadsPerCore() > 1 {
			rval |= htt
		}
	}
	// Check XGETBV, OXSAVE and AVX bits
	if c&(1<<26) != 0 && c&(1<<27) != 0 && c&(1<<28) != 0 {
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
		_, ebx, ecx, edx := cpuidex(7, 0)
		eax1, _, _, _ := cpuidex(7, 1)
		if (rval&avx) != 0 && (ebx&0x00000020) != 0 {
			rval |= avx2
		}
		if (ebx & 0x00000008) != 0 {
			rval |= bmi1
			if (ebx & 0x00000100) != 0 {
				rval |= bmi2
			}
		}
		if ebx&(1<<2) != 0 {
			rval |= sgx
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
		if edx&(1<<26) != 0 {
			rval |= ibpb
		}
		if ecx&(1<<30) != 0 {
			rval |= sgxlc
		}
		if edx&(1<<27) != 0 {
			rval |= stibp
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
				if ecx&(1<<6) != 0 {
					rval |= avx512vbmi2
				}
				if ecx&(1<<8) != 0 {
					rval |= gfni
				}
				if ecx&(1<<9) != 0 {
					rval |= vaes
				}
				if ecx&(1<<10) != 0 {
					rval |= vpclmulqdq
				}
				if ecx&(1<<11) != 0 {
					rval |= avx512vnni
				}
				if ecx&(1<<12) != 0 {
					rval |= avx512bitalg
				}
				if ecx&(1<<14) != 0 {
					rval |= avx512vpopcntdq
				}
				// edx
				if edx&(1<<8) != 0 {
					rval |= avx512vp2intersect
				}
				// cpuid eax 07h,ecx=1
				if eax1&(1<<5) != 0 {
					rval |= avx512bf16
				}
			}
		}
	}

	if maxExtendedFunction() >= 0x80000001 {
		_, _, c, d := cpuid(0x80000001)
		if (c & (1 << 5)) != 0 {
			rval |= lzcnt
			rval |= popcnt
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
		if d&(1<<27) != 0 {
			rval |= rdtscp
		}

		/* Allow for selectively disabling SSE2 functions on AMD processors
		   with SSE2 support but not SSE4a. This includes Athlon64, some
		   Opteron, and some Sempron processors. MMX, SSE, or 3DNow! are faster
		   than SSE2 often enough to utilize this special-case flag.
		   AV_CPU_FLAG_SSE2 and AV_CPU_FLAG_SSE2SLOW are both set in this case
		   so that SSE2 is used unless explicitly disabled by checking
		   AV_CPU_FLAG_SSE2SLOW. */
		if vend != intel &&
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

		if vend == intel {
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
		case dst[1] == 0:
			return r[:i*4+1]
		case dst[2] == 0:
			return r[:i*4+2]
		case dst[3] == 0:
			return r[:i*4+3]
		}
	}
	return r
}

// Single-precision and double-precision floating point
func (c cpuInfo) armfp() bool {
	return c.arm&arm_fp != 0
}

// Advanced SIMD
func (c cpuInfo) armasimd() bool {
	return c.arm&arm_asimd != 0
}

// Generic timer
func (c cpuInfo) armevtstrm() bool {
	return c.arm&arm_evtstrm != 0
}

// AES instructions
func (c cpuInfo) armaes() bool {
	return c.arm&arm_aes != 0
}

// Polynomial Multiply instructions (PMULL/PMULL2)
func (c cpuInfo) armpmull() bool {
	return c.arm&arm_pmull != 0
}

// SHA-1 instructions (SHA1C, etc)
func (c cpuInfo) armsha1() bool {
	return c.arm&arm_sha1 != 0
}

// SHA-2 instructions (SHA256H, etc)
func (c cpuInfo) armsha2() bool {
	return c.arm&arm_sha2 != 0
}

// CRC32/CRC32C instructions
func (c cpuInfo) armcrc32() bool {
	return c.arm&arm_crc32 != 0
}

// Large System Extensions (LSE)
func (c cpuInfo) armatomics() bool {
	return c.arm&arm_atomics != 0
}

// Half-precision floating point
func (c cpuInfo) armfphp() bool {
	return c.arm&arm_fphp != 0
}

// Advanced SIMD half-precision floating point
func (c cpuInfo) armasimdhp() bool {
	return c.arm&arm_asimdhp != 0
}

// Rounding Double Multiply Accumulate/Subtract (SQRDMLAH/SQRDMLSH)
func (c cpuInfo) armasimdrdm() bool {
	return c.arm&arm_asimdrdm != 0
}

// Javascript-style double->int convert (FJCVTZS)
func (c cpuInfo) armjscvt() bool {
	return c.arm&arm_jscvt != 0
}

// Floatin point complex number addition and multiplication
func (c cpuInfo) armfcma() bool {
	return c.arm&arm_fcma != 0
}

// Weaker release consistency (LDAPR, etc)
func (c cpuInfo) armlrcpc() bool {
	return c.arm&arm_lrcpc != 0
}

// Data cache clean to Point of Persistence (DC CVAP)
func (c cpuInfo) armdcpop() bool {
	return c.arm&arm_dcpop != 0
}

// SHA-3 instructions (EOR3, RAXI, XAR, BCAX)
func (c cpuInfo) armsha3() bool {
	return c.arm&arm_sha3 != 0
}

// SM3 instructions
func (c cpuInfo) armsm3() bool {
	return c.arm&arm_sm3 != 0
}

// SM4 instructions
func (c cpuInfo) armsm4() bool {
	return c.arm&arm_sm4 != 0
}

// SIMD Dot Product
func (c cpuInfo) armasimddp() bool {
	return c.arm&arm_asimddp != 0
}

// SHA512 instructions
func (c cpuInfo) armsha512() bool {
	return c.arm&arm_sha512 != 0
}

// Scalable Vector Extension
func (c cpuInfo) armsve() bool {
	return c.arm&arm_sve != 0
}

// Generic Pointer Authentication
func (c cpuInfo) armgpa() bool {
	return c.arm&arm_gpa != 0
}
