// Copyright (c) 2015 Klaus Post, released under MIT License. See LICENSE file.

// Package cpuid provides information about the CPU running the current program.
//
// CPU features are detected on startup, and kept for fast access through the life of the application.
// Currently x86 / x64 (AMD64) as well as arm64 is supported.
//
// You can access the CPU information by accessing the shared CPU variable of the cpuid library.
//
// Package home: https://github.com/klauspost/cpuid
package cpuid

import (
	"math"
	"strings"
)

// AMD refererence: https://www.amd.com/system/files/TechDocs/25481.pdf
// and Processor Programming Reference (PPR)

// Vendor is a representation of a CPU vendor.
type Vendor int

const (
	Other Vendor = iota
	Intel
	AMD
	VIA
	Transmeta
	NSC
	KVM  // Kernel-based Virtual Machine
	MSVM // Microsoft Hyper-V or Windows Virtual PC
	VMware
	XenHVM
	Bhyve
	Hygon
	SiS
	RDC
)

const (
	CMOV               = iota // i686 CMOV
	NX                        // NX (No-Execute) bit
	AMD3DNOW                  // AMD 3DNOW
	AMD3DNOWEXT               // AMD 3DNowExt
	MMX                       // standard MMX
	MMXEXT                    // SSE integer functions or AMD MMX ext
	SSE                       // SSE functions
	SSE2                      // P4 SSE functions
	SSE3                      // Prescott SSE3 functions
	SSSE3                     // Conroe SSSE3 functions
	SSE4                      // Penryn SSE4.1 functions
	SSE4A                     // AMD Barcelona microarchitecture SSE4a instructions
	SSE42                     // Nehalem SSE4.2 functions
	AVX                       // AVX functions
	AVX2                      // AVX2 functions
	FMA3                      // Intel FMA 3
	FMA4                      // Bulldozer FMA4 functions
	XOP                       // Bulldozer XOP functions
	F16C                      // Half-precision floating-point conversion
	BMI1                      // Bit Manipulation Instruction Set 1
	BMI2                      // Bit Manipulation Instruction Set 2
	TBM                       // AMD Trailing Bit Manipulation
	LZCNT                     // LZCNT instruction
	POPCNT                    // POPCNT instruction
	AESNI                     // Advanced Encryption Standard New Instructions
	CLMUL                     // Carry-less Multiplication
	HTT                       // Hyperthreading (enabled)
	HLE                       // Hardware Lock Elision
	RTM                       // Restricted Transactional Memory
	RDRAND                    // RDRAND instruction is available
	RDSEED                    // RDSEED instruction is available
	ADX                       // Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
	SHA                       // Intel SHA Extensions
	AVX512F                   // AVX-512 Foundation
	AVX512DQ                  // AVX-512 Doubleword and Quadword Instructions
	AVX512IFMA                // AVX-512 Integer Fused Multiply-Add Instructions
	AVX512PF                  // AVX-512 Prefetch Instructions
	AVX512ER                  // AVX-512 Exponential and Reciprocal Instructions
	AVX512CD                  // AVX-512 Conflict Detection Instructions
	AVX512BW                  // AVX-512 Byte and Word Instructions
	AVX512VL                  // AVX-512 Vector Length Extensions
	AVX512VBMI                // AVX-512 Vector Bit Manipulation Instructions
	AVX512VBMI2               // AVX-512 Vector Bit Manipulation Instructions, Version 2
	AVX512VNNI                // AVX-512 Vector Neural Network Instructions
	AVX512VPOPCNTDQ           // AVX-512 Vector Population Count Doubleword and Quadword
	GFNI                      // Galois Field New Instructions
	VAES                      // Vector AES
	AVX512BITALG              // AVX-512 Bit Algorithms
	VPCLMULQDQ                // Carry-Less Multiplication Quadword
	AVX512BF16                // AVX-512 BFLOAT16 Instructions
	AVX512VP2INTERSECT        // AVX-512 Intersect for D/Q
	MPX                       // Intel MPX (Memory Protection Extensions)
	ERMS                      // Enhanced REP MOVSB/STOSB
	RDTSCP                    // RDTSCP Instruction
	CX16                      // CMPXCHG16B Instruction
	SGX                       // Software Guard Extensions
	SGXLC                     // Software Guard Extensions Launch Control
	IBPB                      // Indirect Branch Restricted Speculation (IBRS) and Indirect Branch Predictor Barrier (IBPB)
	STIBP                     // Single Thread Indirect Branch Predictors
	VMX                       // Virtual Machine Extensions

	// Performance indicators
	SSE2SLOW // SSE2 is supported, but usually not faster
	SSE3SLOW // SSE3 is supported, but usually not faster
	ATOM     // Atom processor, some SSSE3 instructions are slower

	AMXBF16 // Tile computational operations on BFLOAT16 numbers
	AMXTILE // Tile architecture
	AMXINT8 // Tile computational operations on 8-bit integers

	HYPERVISOR // This bit has been reserved by Intel & AMD for use by hypervisors

	WAITPKG   // TPAUSE, UMONITOR, UMWAIT
	SERIALIZE // Serialize Instruction Execution
	TSXLDTRK  // Intel TSX Suspend Load Address Tracking
	WBNOINVD  // Write Back and Do Not Invalidate Cache
	MOVDIRI   // Move Doubleword as Direct Store
	MOVDIR64B // Move 64 Bytes as Direct Store
	ENQCMD    // Enqueue Command
	CLDEMOTE  // Cache Line Demote

	// Keep it last. It automatically defines the size of []FlagSet
	LASTID
)

var flagNames = map[Flags]string{
	CMOV:               "CMOV",               // i686 CMOV
	NX:                 "NX",                 // NX (No-Execute) bit
	AMD3DNOW:           "AMD3DNOW",           // AMD 3DNOW
	AMD3DNOWEXT:        "AMD3DNOWEXT",        // AMD 3DNowExt
	MMX:                "MMX",                // Standard MMX
	MMXEXT:             "MMXEXT",             // SSE integer functions or AMD MMX ext
	SSE:                "SSE",                // SSE functions
	SSE2:               "SSE2",               // P4 SSE2 functions
	SSE3:               "SSE3",               // Prescott SSE3 functions
	SSSE3:              "SSSE3",              // Conroe SSSE3 functions
	SSE4:               "SSE4.1",             // Penryn SSE4.1 functions
	SSE4A:              "SSE4A",              // AMD Barcelona microarchitecture SSE4a instructions
	SSE42:              "SSE4.2",             // Nehalem SSE4.2 functions
	AVX:                "AVX",                // AVX functions
	AVX2:               "AVX2",               // AVX functions
	FMA3:               "FMA3",               // Intel FMA 3
	FMA4:               "FMA4",               // Bulldozer FMA4 functions
	XOP:                "XOP",                // Bulldozer XOP functions
	F16C:               "F16C",               // Half-precision floating-point conversion
	BMI1:               "BMI1",               // Bit Manipulation Instruction Set 1
	BMI2:               "BMI2",               // Bit Manipulation Instruction Set 2
	TBM:                "TBM",                // AMD Trailing Bit Manipulation
	LZCNT:              "LZCNT",              // LZCNT instruction
	POPCNT:             "POPCNT",             // POPCNT instruction
	AESNI:              "AESNI",              // Advanced Encryption Standard New Instructions
	CLMUL:              "CLMUL",              // Carry-less Multiplication
	HTT:                "HTT",                // Hyperthreading (enabled)
	HLE:                "HLE",                // Hardware Lock Elision
	RTM:                "RTM",                // Restricted Transactional Memory
	RDRAND:             "RDRAND",             // RDRAND instruction is available
	RDSEED:             "RDSEED",             // RDSEED instruction is available
	ADX:                "ADX",                // Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
	SHA:                "SHA",                // Intel SHA Extensions
	AVX512F:            "AVX512F",            // AVX-512 Foundation
	AVX512DQ:           "AVX512DQ",           // AVX-512 Doubleword and Quadword Instructions
	AVX512IFMA:         "AVX512IFMA",         // AVX-512 Integer Fused Multiply-Add Instructions
	AVX512PF:           "AVX512PF",           // AVX-512 Prefetch Instructions
	AVX512ER:           "AVX512ER",           // AVX-512 Exponential and Reciprocal Instructions
	AVX512CD:           "AVX512CD",           // AVX-512 Conflict Detection Instructions
	AVX512BW:           "AVX512BW",           // AVX-512 Byte and Word Instructions
	AVX512VL:           "AVX512VL",           // AVX-512 Vector Length Extensions
	AVX512VBMI:         "AVX512VBMI",         // AVX-512 Vector Bit Manipulation Instructions
	AVX512VBMI2:        "AVX512VBMI2",        // AVX-512 Vector Bit Manipulation Instructions, Version 2
	AVX512VNNI:         "AVX512VNNI",         // AVX-512 Vector Neural Network Instructions
	AVX512VPOPCNTDQ:    "AVX512VPOPCNTDQ",    // AVX-512 Vector Population Count Doubleword and Quadword
	GFNI:               "GFNI",               // Galois Field New Instructions
	VAES:               "VAES",               // Vector AES
	AVX512BITALG:       "AVX512BITALG",       // AVX-512 Bit Algorithms
	VPCLMULQDQ:         "VPCLMULQDQ",         // Carry-Less Multiplication Quadword
	AVX512BF16:         "AVX512BF16",         // AVX-512 BFLOAT16 Instruction
	AVX512VP2INTERSECT: "AVX512VP2INTERSECT", // AVX-512 Intersect for D/Q
	MPX:                "MPX",                // Intel MPX (Memory Protection Extensions)
	ERMS:               "ERMS",               // Enhanced REP MOVSB/STOSB
	RDTSCP:             "RDTSCP",             // RDTSCP Instruction
	CX16:               "CX16",               // CMPXCHG16B Instruction
	SGX:                "SGX",                // Software Guard Extensions
	SGXLC:              "SGXLC",              // Software Guard Extensions Launch Control
	IBPB:               "IBPB",               // Indirect Branch Restricted Speculation and Indirect Branch Predictor Barrier
	STIBP:              "STIBP",              // Single Thread Indirect Branch Predictors
	VMX:                "VMX",                // Virtual Machine Extensions

	// Performance indicators
	SSE2SLOW: "SSE2SLOW", // SSE2 supported, but usually not faster
	SSE3SLOW: "SSE3SLOW", // SSE3 supported, but usually not faster
	ATOM:     "ATOM",     // Atom processor, some SSSE3 instructions are slower

	AMXBF16: "AMXBF16", // Tile computational operations on BFLOAT16 numbers
	AMXTILE: "AMXTILE", // Tile architecture
	AMXINT8: "AMXINT8", // Tile computational operations on 8-bit integers

	HYPERVISOR: "HYPERVISOR", // This bit has been reserved by Intel & AMD for use by hypervisors

	WAITPKG:   "WAITPKG",   // TPAUSE, UMONITOR, UMWAIT
	SERIALIZE: "SERIALIZE", // Serialize Instruction Execution
	TSXLDTRK:  "TSXLDTRK",  // Intel TSX Suspend Load Address Tracking
	WBNOINVD:  "WBNOINVD",  // Write Back and Do Not Invalidate Cache
	MOVDIRI:   "MOVDIRI",   // Move Doubleword as Direct Store
	MOVDIR64B: "MOVDIR64B", // Move 64 Bytes as Direct Store
	ENQCMD:    "ENQCMD",    // Enqueue Command
	CLDEMOTE:  "CLDEMOTE",  // Cache Line Demote
}

/* all special features for arm64 should be defined here */
const (
	/* extension instructions */
	FP ArmFlags = 1 << iota
	ASIMD
	EVTSTRM
	AES
	PMULL
	SHA1
	SHA2
	CRC32
	ATOMICS
	FPHP
	ASIMDHP
	ARMCPUID
	ASIMDRDM
	JSCVT
	FCMA
	LRCPC
	DCPOP
	SHA3
	SM3
	SM4
	ASIMDDP
	SHA512
	SVE
	GPA
)

var flagNamesArm = map[ArmFlags]string{
	FP:       "FP",       // Single-precision and double-precision floating point
	ASIMD:    "ASIMD",    // Advanced SIMD
	EVTSTRM:  "EVTSTRM",  // Generic timer
	AES:      "AES",      // AES instructions
	PMULL:    "PMULL",    // Polynomial Multiply instructions (PMULL/PMULL2)
	SHA1:     "SHA1",     // SHA-1 instructions (SHA1C, etc)
	SHA2:     "SHA2",     // SHA-2 instructions (SHA256H, etc)
	CRC32:    "CRC32",    // CRC32/CRC32C instructions
	ATOMICS:  "ATOMICS",  // Large System Extensions (LSE)
	FPHP:     "FPHP",     // Half-precision floating point
	ASIMDHP:  "ASIMDHP",  // Advanced SIMD half-precision floating point
	ARMCPUID: "CPUID",    // Some CPU ID registers readable at user-level
	ASIMDRDM: "ASIMDRDM", // Rounding Double Multiply Accumulate/Subtract (SQRDMLAH/SQRDMLSH)
	JSCVT:    "JSCVT",    // Javascript-style double->int convert (FJCVTZS)
	FCMA:     "FCMA",     // Floatin point complex number addition and multiplication
	LRCPC:    "LRCPC",    // Weaker release consistency (LDAPR, etc)
	DCPOP:    "DCPOP",    // Data cache clean to Point of Persistence (DC CVAP)
	SHA3:     "SHA3",     // SHA-3 instructions (EOR3, RAXI, XAR, BCAX)
	SM3:      "SM3",      // SM3 instructions
	SM4:      "SM4",      // SM4 instructions
	ASIMDDP:  "ASIMDDP",  // SIMD Dot Product
	SHA512:   "SHA512",   // SHA512 instructions
	SVE:      "SVE",      // Scalable Vector Extension
	GPA:      "GPA",      // Generic Pointer Authentication
}

// CPUInfo contains information about the detected system CPU.
type CPUInfo struct {
	BrandName      string   // Brand name reported by the CPU
	VendorID       Vendor   // Comparable CPU vendor ID
	VendorString   string   // Raw vendor string.
	Features       Flags    // Features of the CPU (x64) (deprecated)
	featureSet     FlagSet  // Features of the CPU
	Arm            ArmFlags // Features of the CPU (arm)
	PhysicalCores  int      // Number of physical processor cores in your CPU. Will be 0 if undetectable.
	ThreadsPerCore int      // Number of threads per physical core. Will be 1 if undetectable.
	LogicalCores   int      // Number of physical cores times threads that can run on each core through the use of hyperthreading. Will be 0 if undetectable.
	Family         int      // CPU family number
	Model          int      // CPU model number
	CacheLine      int      // Cache line size in bytes. Will be 0 if undetectable.
	Hz             int64    // Clock speed, if known
	Cache          struct {
		L1I int // L1 Instruction Cache (per core or shared). Will be -1 if undetected
		L1D int // L1 Data Cache (per core or shared). Will be -1 if undetected
		L2  int // L2 Cache (per core or shared). Will be -1 if undetected
		L3  int // L3 Cache (per core, per ccx or shared). Will be -1 if undetected
	}
	SGX       SGXSupport
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
// Use this as the primary entry point to you data.
var CPU CPUInfo

func init() {
	initCPU()
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
	// Set defaults
	CPU.ThreadsPerCore = 1
	CPU.Cache.L1I = -1
	CPU.Cache.L1D = -1
	CPU.Cache.L2 = -1
	CPU.Cache.L3 = -1
	addInfo(&CPU)
}

// Generated here: http://play.golang.org/p/BxFH2Gdc0G

// Cmov indicates support of CMOV instructions
func (c CPUInfo) Cmov() bool {
	return c.featureSet.inSet(CMOV)
}

// Amd3dnow indicates support of AMD 3DNOW! instructions
func (c CPUInfo) Amd3dnow() bool {
	return c.featureSet.inSet(AMD3DNOW)
}

// Amd3dnowExt indicates support of AMD 3DNOW! Extended instructions
func (c CPUInfo) Amd3dnowExt() bool {
	return c.featureSet.inSet(AMD3DNOWEXT)
}

// VMX indicates support of VMX
func (c CPUInfo) VMX() bool {
	return c.featureSet.inSet(VMX)
}

// MMX indicates support of MMX instructions
func (c CPUInfo) MMX() bool {
	return c.featureSet.inSet(MMX)
}

// MMXExt indicates support of MMXEXT instructions
// (SSE integer functions or AMD MMX ext)
func (c CPUInfo) MMXExt() bool {
	return c.featureSet.inSet(MMXEXT)
}

// SSE indicates support of SSE instructions
func (c CPUInfo) SSE() bool {
	return c.featureSet.inSet(SSE)
}

// SSE2 indicates support of SSE 2 instructions
func (c CPUInfo) SSE2() bool {
	return c.featureSet.inSet(SSE2)
}

// SSE3 indicates support of SSE 3 instructions
func (c CPUInfo) SSE3() bool {
	return c.featureSet.inSet(SSE3)
}

// SSSE3 indicates support of SSSE 3 instructions
func (c CPUInfo) SSSE3() bool {
	return c.featureSet.inSet(SSSE3)
}

// SSE4 indicates support of SSE 4 (also called SSE 4.1) instructions
func (c CPUInfo) SSE4() bool {
	return c.featureSet.inSet(SSE4)
}

// SSE42 indicates support of SSE4.2 instructions
func (c CPUInfo) SSE42() bool {
	return c.featureSet.inSet(SSE42)
}

// AVX indicates support of AVX instructions
// and operating system support of AVX instructions
func (c CPUInfo) AVX() bool {
	return c.featureSet.inSet(AVX)
}

// AVX2 indicates support of AVX2 instructions
func (c CPUInfo) AVX2() bool {
	return c.featureSet.inSet(AVX2)
}

// FMA3 indicates support of FMA3 instructions
func (c CPUInfo) FMA3() bool {
	return c.featureSet.inSet(FMA3)
}

// FMA4 indicates support of FMA4 instructions
func (c CPUInfo) FMA4() bool {
	return c.featureSet.inSet(FMA4)
}

// XOP indicates support of XOP instructions
func (c CPUInfo) XOP() bool {
	return c.featureSet.inSet(XOP)
}

// F16C indicates support of F16C instructions
func (c CPUInfo) F16C() bool {
	return c.featureSet.inSet(F16C)
}

// BMI1 indicates support of BMI1 instructions
func (c CPUInfo) BMI1() bool {
	return c.featureSet.inSet(BMI1)
}

// BMI2 indicates support of BMI2 instructions
func (c CPUInfo) BMI2() bool {
	return c.featureSet.inSet(BMI2)
}

// TBM indicates support of TBM instructions
// (AMD Trailing Bit Manipulation)
func (c CPUInfo) TBM() bool {
	return c.featureSet.inSet(TBM)
}

// Lzcnt indicates support of LZCNT instruction
func (c CPUInfo) Lzcnt() bool {
	return c.featureSet.inSet(LZCNT)
}

// Popcnt indicates support of POPCNT instruction
func (c CPUInfo) Popcnt() bool {
	return c.featureSet.inSet(POPCNT)
}

// HTT indicates the processor has Hyperthreading enabled
func (c CPUInfo) HTT() bool {
	return c.featureSet.inSet(HTT)
}

// SSE2Slow indicates that SSE2 may be slow on this processor
func (c CPUInfo) SSE2Slow() bool {
	return c.featureSet.inSet(SSE2SLOW)
}

// SSE3Slow indicates that SSE3 may be slow on this processor
func (c CPUInfo) SSE3Slow() bool {
	return c.featureSet.inSet(SSE3SLOW)
}

// AesNi indicates support of AES-NI instructions
// (Advanced Encryption Standard New Instructions)
func (c CPUInfo) AesNi() bool {
	return c.featureSet.inSet(AESNI)
}

// Clmul indicates support of CLMUL instructions
// (Carry-less Multiplication)
func (c CPUInfo) Clmul() bool {
	return c.featureSet.inSet(CLMUL)
}

// NX indicates support of NX (No-Execute) bit
func (c CPUInfo) NX() bool {
	return c.featureSet.inSet(NX)
}

// SSE4A indicates support of AMD Barcelona microarchitecture SSE4a instructions
func (c CPUInfo) SSE4A() bool {
	return c.featureSet.inSet(SSE4A)
}

// HLE indicates support of Hardware Lock Elision
func (c CPUInfo) HLE() bool {
	return c.featureSet.inSet(HLE)
}

// RTM indicates support of Restricted Transactional Memory
func (c CPUInfo) RTM() bool {
	return c.featureSet.inSet(RTM)
}

// Rdrand indicates support of RDRAND instruction is available
func (c CPUInfo) Rdrand() bool {
	return c.featureSet.inSet(RDRAND)
}

// Rdseed indicates support of RDSEED instruction is available
func (c CPUInfo) Rdseed() bool {
	return c.featureSet.inSet(RDSEED)
}

// ADX indicates support of Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
func (c CPUInfo) ADX() bool {
	return c.featureSet.inSet(ADX)
}

// SHA indicates support of Intel SHA Extensions
func (c CPUInfo) SHA() bool {
	return c.featureSet.inSet(SHA)
}

// AVX512F indicates support of AVX-512 Foundation
func (c CPUInfo) AVX512F() bool {
	return c.featureSet.inSet(AVX512F)
}

// AVX512DQ indicates support of AVX-512 Doubleword and Quadword Instructions
func (c CPUInfo) AVX512DQ() bool {
	return c.featureSet.inSet(AVX512DQ)
}

// AVX512IFMA indicates support of AVX-512 Integer Fused Multiply-Add Instructions
func (c CPUInfo) AVX512IFMA() bool {
	return c.featureSet.inSet(AVX512IFMA)
}

// AVX512PF indicates support of AVX-512 Prefetch Instructions
func (c CPUInfo) AVX512PF() bool {
	return c.featureSet.inSet(AVX512PF)
}

// AVX512ER indicates support of AVX-512 Exponential and Reciprocal Instructions
func (c CPUInfo) AVX512ER() bool {
	return c.featureSet.inSet(AVX512ER)
}

// AVX512CD indicates support of AVX-512 Conflict Detection Instructions
func (c CPUInfo) AVX512CD() bool {
	return c.featureSet.inSet(AVX512CD)
}

// AVX512BW indicates support of AVX-512 Byte and Word Instructions
func (c CPUInfo) AVX512BW() bool {
	return c.featureSet.inSet(AVX512BW)
}

// AVX512VL indicates support of AVX-512 Vector Length Extensions
func (c CPUInfo) AVX512VL() bool {
	return c.featureSet.inSet(AVX512VL)
}

// AVX512VBMI indicates support of AVX-512 Vector Bit Manipulation Instructions
func (c CPUInfo) AVX512VBMI() bool {
	return c.featureSet.inSet(AVX512VBMI)
}

// AVX512VBMI2 indicates support of AVX-512 Vector Bit Manipulation Instructions, Version 2
func (c CPUInfo) AVX512VBMI2() bool {
	return c.featureSet.inSet(AVX512VBMI2)
}

// AVX512VNNI indicates support of AVX-512 Vector Neural Network Instructions
func (c CPUInfo) AVX512VNNI() bool {
	return c.featureSet.inSet(AVX512VNNI)
}

// AVX512VPOPCNTDQ indicates support of AVX-512 Vector Population Count Doubleword and Quadword
func (c CPUInfo) AVX512VPOPCNTDQ() bool {
	return c.featureSet.inSet(AVX512VPOPCNTDQ)
}

// GFNI indicates support of Galois Field New Instructions
func (c CPUInfo) GFNI() bool {
	return c.featureSet.inSet(GFNI)
}

// VAES indicates support of Vector AES
func (c CPUInfo) VAES() bool {
	return c.featureSet.inSet(VAES)
}

// AVX512BITALG indicates support of AVX-512 Bit Algorithms
func (c CPUInfo) AVX512BITALG() bool {
	return c.featureSet.inSet(AVX512BITALG)
}

// VPCLMULQDQ indicates support of Carry-Less Multiplication Quadword
func (c CPUInfo) VPCLMULQDQ() bool {
	return c.featureSet.inSet(VPCLMULQDQ)
}

// AVX512BF16 indicates support of AVX-512 BFLOAT16 Instruction
func (c CPUInfo) AVX512BF16() bool {
	return c.featureSet.inSet(AVX512BF16)
}

// AVX512VP2INTERSECT indicates support of AVX-512 Intersect for D/Q
func (c CPUInfo) AVX512VP2INTERSECT() bool {
	return c.featureSet.inSet(AVX512VP2INTERSECT)
}

// AMXBF16 indicates support of Tile computational operations on BFLOAT16 numbers
func (c CPUInfo) AMXBF16() bool {
	return c.featureSet.inSet(AMXBF16)
}

// AMXTILE indicates support of Tile architecture
func (c CPUInfo) AMXTILE() bool {
	return c.featureSet.inSet(AMXTILE)
}

// AMXINT8 indicates support of Tile computational operations on 8-bit integers
func (c CPUInfo) AMXINT8() bool {
	return c.featureSet.inSet(AMXINT8)
}

// WAITPKG indicates support of TPAUSE, UMONITOR, UMWAIT
func (c CPUInfo) WAITPKG() bool {
	return c.featureSet.inSet(WAITPKG)
}

// SERIALIZE indicates support of Serialize Instruction Execution
func (c CPUInfo) SERIALIZE() bool {
	return c.featureSet.inSet(SERIALIZE)
}

// TSXLDTRK indicates support of Intel TSX Suspend Load Address Tracking
func (c CPUInfo) TSXLDTRK() bool {
	return c.featureSet.inSet(TSXLDTRK)
}

// WBNOINVD indicates support of Write Back and Do Not Invalidate Cache
func (c CPUInfo) WBNOINVD() bool {
	return c.featureSet.inSet(WBNOINVD)
}

// MOVDIRI indicates support of Move Doubleword as Direct Store
func (c CPUInfo) MOVDIRI() bool {
	return c.featureSet.inSet(MOVDIRI)
}

// MOVDIR64B indicates support of Move 64 Bytes as Direct Store
func (c CPUInfo) MOVDIR64B() bool {
	return c.featureSet.inSet(MOVDIR64B)
}

// ENQCMD indicates support of Enqueue Command
func (c CPUInfo) ENQCMD() bool {
	return c.featureSet.inSet(ENQCMD)
}

// CLDEMOTE indicates support of Cache Line Demote
func (c CPUInfo) CLDEMOTE() bool {
	return c.featureSet.inSet(CLDEMOTE)
}

// MPX indicates support of Intel MPX (Memory Protection Extensions)
func (c CPUInfo) MPX() bool {
	return c.featureSet.inSet(MPX)
}

// ERMS indicates support of Enhanced REP MOVSB/STOSB
func (c CPUInfo) ERMS() bool {
	return c.featureSet.inSet(ERMS)
}

// RDTSCP Instruction is available.
func (c CPUInfo) RDTSCP() bool {
	return c.featureSet.inSet(RDTSCP)
}

// CX16 indicates if CMPXCHG16B instruction is available.
func (c CPUInfo) CX16() bool {
	return c.featureSet.inSet(CX16)
}

// TSX is split into HLE (Hardware Lock Elision) and RTM (Restricted Transactional Memory) detection.
// So TSX simply checks that.
func (c CPUInfo) TSX() bool {
	return c.featureSet.inSet(HLE) && c.featureSet.inSet(RTM)
}

// Atom indicates an Atom processor
func (c CPUInfo) Atom() bool {
	return c.featureSet.inSet(ATOM)
}

// Intel returns true if vendor is recognized as Intel
func (c CPUInfo) Intel() bool {
	return c.VendorID == Intel
}

// AMD returns true if vendor is recognized as AMD
func (c CPUInfo) AMD() bool {
	return c.VendorID == AMD
}

// Hygon returns true if vendor is recognized as Hygon
func (c CPUInfo) Hygon() bool {
	return c.VendorID == Hygon
}

// Transmeta returns true if vendor is recognized as Transmeta
func (c CPUInfo) Transmeta() bool {
	return c.VendorID == Transmeta
}

// NSC returns true if vendor is recognized as National Semiconductor
func (c CPUInfo) NSC() bool {
	return c.VendorID == NSC
}

// VIA returns true if vendor is recognized as VIA
func (c CPUInfo) VIA() bool {
	return c.VendorID == VIA
}

func (c CPUInfo) FeatureSet() []string {
	s := make([]string, 0)
	for _, f := range c.featureSet.Strings() {
		s = append(s, f)
	}
	return s
}

// RTCounter returns the 64-bit time-stamp counter
// Uses the RDTSCP instruction. The value 0 is returned
// if the CPU does not support the instruction.
func (c CPUInfo) RTCounter() uint64 {
	if !c.RDTSCP() {
		return 0
	}
	a, _, _, d := rdtscpAsm()
	return uint64(a) | (uint64(d) << 32)
}

// Ia32TscAux returns the IA32_TSC_AUX part of the RDTSCP.
// This variable is OS dependent, but on Linux contains information
// about the current cpu/core the code is running on.
// If the RDTSCP instruction isn't supported on the CPU, the value 0 is returned.
func (c CPUInfo) Ia32TscAux() uint32 {
	if !c.RDTSCP() {
		return 0
	}
	_, _, ecx, _ := rdtscpAsm()
	return ecx
}

// LogicalCPU will return the Logical CPU the code is currently executing on.
// This is likely to change when the OS re-schedules the running thread
// to another CPU.
// If the current core cannot be detected, -1 will be returned.
func (c CPUInfo) LogicalCPU() int {
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
// a virtual machine.
func (c CPUInfo) VM() bool {
	return CPU.featureSet.inSet(HYPERVISOR)
}

// Flags contains detected cpu features and characteristics
type Flags uint64

// ArmFlags contains detected ARM cpu features and characteristics
type ArmFlags uint64

// FlagSet contains detected cpu features and characteristics in an array of Flags
type FlagSet []Flags

// String returns a string representation of the detected
// CPU features.
func (f Flags) String() string {
	return strings.Join(f.Strings(), ",")
}

func (s FlagSet) inSet(offset uint) bool {
	if len(s) == 0 {
		return false
	}
	return s[offset>>6]&(1<<(offset&63)) != 0
}

func (s FlagSet) set(offset uint) {
	if len(s) == 0 {
		return
	}
	s[offset>>6] |= 1 << (offset & 63)
}

// Strings returns an array of the detected features.
func (f Flags) Strings() []string {
	s := FlagSet{f}
	return s.Strings()
}

// Strings returns an array of the detected features for FlagsSet.
func (s FlagSet) Strings() []string {
	if len(s) == 0 {
		return []string{""}
	}
	bits := uint(len(s) * 64)
	r := make([]string, 0, bits)
	for i := uint(0); i < bits; i++ {
		if s.inSet(i) {
			r = append(r, flagNames[Flags(i)])
		}
	}
	return r
}

// String returns a string representation of the detected
// CPU features.
func (f ArmFlags) String() string {
	return strings.Join(f.Strings(), ",")
}

// Strings returns an array of the detected features.
func (f ArmFlags) Strings() []string {
	r := make([]string, 0, 20)
	for i := uint(0); i < 64; i++ {
		key := ArmFlags(1 << i)
		val := flagNamesArm[key]
		if f&key != 0 {
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

	if mfi < 0x4 || (vend != Intel && vend != AMD) {
		return 1
	}

	if mfi < 0xb {
		if vend != Intel {
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
	case Intel:
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
	case AMD, Hygon:
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
	case Intel:
		return logicalCores() / threadsPerCore()
	case AMD, Hygon:
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
var vendorMapping = map[string]Vendor{
	"AMDisbetter!": AMD,
	"AuthenticAMD": AMD,
	"CentaurHauls": VIA,
	"GenuineIntel": Intel,
	"TransmetaCPU": Transmeta,
	"GenuineTMx86": Transmeta,
	"Geode by NSC": NSC,
	"VIA VIA VIA ": VIA,
	"KVMKVMKVMKVM": KVM,
	"Microsoft Hv": MSVM,
	"VMwareVMware": VMware,
	"XenVMMXenVMM": XenHVM,
	"bhyve bhyve ": Bhyve,
	"HygonGenuine": Hygon,
	"Vortex86 SoC": SiS,
	"SiS SiS SiS ": SiS,
	"RiseRiseRise": SiS,
	"Genuine  RDC": RDC,
}

func vendorID() (Vendor, string) {
	_, b, c, d := cpuid(0)
	v := string(valAsString(b, d, c))
	vend, ok := vendorMapping[v]
	if !ok {
		return Other, v
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

func (c *CPUInfo) cacheSize() {
	c.Cache.L1D = -1
	c.Cache.L1I = -1
	c.Cache.L2 = -1
	c.Cache.L3 = -1
	vendor, _ := vendorID()
	switch vendor {
	case Intel:
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
					c.Cache.L1D = size
				} else if cacheType == 2 {
					// 2 = Instruction Cache
					c.Cache.L1I = size
				} else {
					if c.Cache.L1D < 0 {
						c.Cache.L1I = size
					}
					if c.Cache.L1I < 0 {
						c.Cache.L1I = size
					}
				}
			case 2:
				c.Cache.L2 = size
			case 3:
				c.Cache.L3 = size
			}
		}
	case AMD, Hygon:
		// Untested.
		if maxExtendedFunction() < 0x80000005 {
			return
		}
		_, _, ecx, edx := cpuid(0x80000005)
		c.Cache.L1D = int(((ecx >> 24) & 0xFF) * 1024)
		c.Cache.L1I = int(((edx >> 24) & 0xFF) * 1024)

		if maxExtendedFunction() < 0x80000006 {
			return
		}
		_, _, ecx, _ = cpuid(0x80000006)
		c.Cache.L2 = int(((ecx >> 16) & 0xFFFF) * 1024)

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
					c.Cache.L1D = size
				case 2:
					// Inst cache
					c.Cache.L1I = size
				default:
					if c.Cache.L1D < 0 {
						c.Cache.L1I = size
					}
					if c.Cache.L1I < 0 {
						c.Cache.L1I = size
					}
				}
			case 2:
				c.Cache.L2 = size
			case 3:
				c.Cache.L3 = size
			}
		}
	}

	return
}

type SGXEPCSection struct {
	BaseAddress uint64
	EPCSize     uint64
}

type SGXSupport struct {
	Available           bool
	LaunchControl       bool
	SGX1Supported       bool
	SGX2Supported       bool
	MaxEnclaveSizeNot64 int64
	MaxEnclaveSize64    int64
	EPCSections         []SGXEPCSection
}

func hasSGX(available, lc bool) (rval SGXSupport) {
	rval.Available = available

	if !available {
		return
	}

	rval.LaunchControl = lc

	a, _, _, d := cpuidex(0x12, 0)
	rval.SGX1Supported = a&0x01 != 0
	rval.SGX2Supported = a&0x02 != 0
	rval.MaxEnclaveSizeNot64 = 1 << (d & 0xFF)     // pow 2
	rval.MaxEnclaveSize64 = 1 << ((d >> 8) & 0xFF) // pow 2
	rval.EPCSections = make([]SGXEPCSection, 0)

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

			section := SGXEPCSection{BaseAddress: baseAddress, EPCSize: size}
			rval.EPCSections = append(rval.EPCSections, section)
		}
	}

	return
}

func support() FlagSet {
	mfi := maxFunctionID()
	vend, _ := vendorID()
	if mfi < 0x1 {
		return nil
	}

	fs := make(FlagSet, LASTID/64+1)

	_, _, c, d := cpuid(1)
	if (d & (1 << 15)) != 0 {
		fs.set(CMOV)
	}
	if (d & (1 << 23)) != 0 {
		fs.set(MMX)
	}
	if (d & (1 << 25)) != 0 {
		fs.set(MMXEXT)
	}
	if (d & (1 << 25)) != 0 {
		fs.set(SSE)
	}
	if (d & (1 << 26)) != 0 {
		fs.set(SSE2)
	}
	if (c & 1) != 0 {
		fs.set(SSE3)
	}
	if (c & (1 << 5)) != 0 {
		fs.set(VMX)
	}
	if (c & 0x00000200) != 0 {
		fs.set(SSSE3)
	}
	if (c & 0x00080000) != 0 {
		fs.set(SSE4)
	}
	if (c & 0x00100000) != 0 {
		fs.set(SSE42)
	}
	if (c & (1 << 25)) != 0 {
		fs.set(AESNI)
	}
	if (c & (1 << 1)) != 0 {
		fs.set(CLMUL)
	}
	if c&(1<<23) != 0 {
		fs.set(POPCNT)
	}
	if c&(1<<30) != 0 {
		fs.set(RDRAND)
	}
	// This bit has been reserved by Intel & AMD for use by hypervisors,
	// and indicates the presence of a hypervisor.
	if c&(1<<31) != 0 {
		fs.set(HYPERVISOR)
	}
	if c&(1<<29) != 0 {
		fs.set(F16C)
	}
	if c&(1<<13) != 0 {
		fs.set(CX16)
	}
	if vend == Intel && (d&(1<<28)) != 0 && mfi >= 4 {
		if threadsPerCore() > 1 {
			fs.set(HTT)
		}
	}
	if vend == AMD && (d&(1<<28)) != 0 && mfi >= 4 {
		if threadsPerCore() > 1 {
			fs.set(HTT)
		}
	}
	// Check XGETBV, OXSAVE and AVX bits
	if c&(1<<26) != 0 && c&(1<<27) != 0 && c&(1<<28) != 0 {
		// Check for OS support
		eax, _ := xgetbv(0)
		if (eax & 0x6) == 0x6 {
			fs.set(AVX)
			if (c & 0x00001000) != 0 {
				fs.set(FMA3)
			}
		}
	}

	// Check AVX2, AVX2 requires OS support, but BMI1/2 don't.
	if mfi >= 7 {
		_, ebx, ecx, edx := cpuidex(7, 0)
		eax1, _, _, _ := cpuidex(7, 1)
		if fs.inSet(AVX) && (ebx&0x00000020) != 0 {
			fs.set(AVX2)
		}
		// CPUID.(EAX=7, ECX=0).EBX
		if (ebx & 0x00000008) != 0 {
			fs.set(BMI1)
			if (ebx & 0x00000100) != 0 {
				fs.set(BMI2)
			}
		}
		if ebx&(1<<2) != 0 {
			fs.set(SGX)
		}
		if ebx&(1<<4) != 0 {
			fs.set(HLE)
		}
		if ebx&(1<<9) != 0 {
			fs.set(ERMS)
		}
		if ebx&(1<<11) != 0 {
			fs.set(RTM)
		}
		if ebx&(1<<14) != 0 {
			fs.set(MPX)
		}
		if ebx&(1<<18) != 0 {
			fs.set(RDSEED)
		}
		if ebx&(1<<19) != 0 {
			fs.set(ADX)
		}
		if ebx&(1<<29) != 0 {
			fs.set(SHA)
		}
		// CPUID.(EAX=7, ECX=0).ECX
		if ecx&(1<<5) != 0 {
			fs.set(WAITPKG)
		}
		if ecx&(1<<25) != 0 {
			fs.set(CLDEMOTE)
		}
		if ecx&(1<<27) != 0 {
			fs.set(MOVDIRI)
		}
		if ecx&(1<<28) != 0 {
			fs.set(MOVDIR64B)
		}
		if ecx&(1<<29) != 0 {
			fs.set(ENQCMD)
		}
		if ecx&(1<<30) != 0 {
			fs.set(SGXLC)
		}
		// CPUID.(EAX=7, ECX=0).EDX
		if edx&(1<<14) != 0 {
			fs.set(SERIALIZE)
		}
		if edx&(1<<16) != 0 {
			fs.set(TSXLDTRK)
		}
		if edx&(1<<26) != 0 {
			fs.set(IBPB)
		}
		if edx&(1<<27) != 0 {
			fs.set(STIBP)
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
					fs.set(AVX512F)
				}
				if ebx&(1<<17) != 0 {
					fs.set(AVX512DQ)
				}
				if ebx&(1<<21) != 0 {
					fs.set(AVX512IFMA)
				}
				if ebx&(1<<26) != 0 {
					fs.set(AVX512PF)
				}
				if ebx&(1<<27) != 0 {
					fs.set(AVX512ER)
				}
				if ebx&(1<<28) != 0 {
					fs.set(AVX512CD)
				}
				if ebx&(1<<30) != 0 {
					fs.set(AVX512BW)
				}
				if ebx&(1<<31) != 0 {
					fs.set(AVX512VL)
				}
				// ecx
				if ecx&(1<<1) != 0 {
					fs.set(AVX512VBMI)
				}
				if ecx&(1<<6) != 0 {
					fs.set(AVX512VBMI2)
				}
				if ecx&(1<<8) != 0 {
					fs.set(GFNI)
				}
				if ecx&(1<<9) != 0 {
					fs.set(VAES)
				}
				if ecx&(1<<10) != 0 {
					fs.set(VPCLMULQDQ)
				}
				if ecx&(1<<11) != 0 {
					fs.set(AVX512VNNI)
				}
				if ecx&(1<<12) != 0 {
					fs.set(AVX512BITALG)
				}
				if ecx&(1<<14) != 0 {
					fs.set(AVX512VPOPCNTDQ)
				}
				// edx
				if edx&(1<<8) != 0 {
					fs.set(AVX512VP2INTERSECT)
				}
				if edx&(1<<22) != 0 {
					fs.set(AMXBF16)
				}
				if edx&(1<<24) != 0 {
					fs.set(AMXTILE)
				}
				if edx&(1<<25) != 0 {
					fs.set(AMXINT8)
				}
				// eax1 = CPUID.(EAX=7, ECX=1).EAX
				if eax1&(1<<5) != 0 {
					fs.set(AVX512BF16)
				}
			}
		}
	}

	if maxExtendedFunction() >= 0x80000001 {
		_, _, c, d := cpuid(0x80000001)
		if (c & (1 << 5)) != 0 {
			fs.set(LZCNT)
			fs.set(POPCNT)
		}
		if (d & (1 << 31)) != 0 {
			fs.set(AMD3DNOW)
		}
		if (d & (1 << 30)) != 0 {
			fs.set(AMD3DNOWEXT)
		}
		if (d & (1 << 23)) != 0 {
			fs.set(MMX)
		}
		if (d & (1 << 22)) != 0 {
			fs.set(MMXEXT)
		}
		if (c & (1 << 6)) != 0 {
			fs.set(SSE4A)
		}
		if d&(1<<20) != 0 {
			fs.set(NX)
		}
		if d&(1<<27) != 0 {
			fs.set(RDTSCP)
		}

		/* Allow for selectively disabling SSE2 functions on AMD processors
		   with SSE2 support but not SSE4a. This includes Athlon64, some
		   Opteron, and some Sempron processors. MMX, SSE, or 3DNow! are faster
		   than SSE2 often enough to utilize this special-case flag.
		   AV_CPU_FLAG_SSE2 and AV_CPU_FLAG_SSE2SLOW are both set in this case
		   so that SSE2 is used unless explicitly disabled by checking
		   AV_CPU_FLAG_SSE2SLOW. */
		if vend != Intel &&
			fs.inSet(SSE2) && (c&0x00000040) == 0 {
			fs.set(SSE2SLOW)
		}

		/* XOP and FMA4 use the AVX instruction coding scheme, so they can't be
		 * used unless the OS has AVX support. */
		if fs.inSet(AVX) {
			if (c & 0x00000800) != 0 {
				fs.set(XOP)
			}
			if (c & 0x00010000) != 0 {
				fs.set(FMA4)
			}
		}

		if vend == Intel {
			family, model := familyModel()
			if family == 6 && (model == 9 || model == 13 || model == 14) {
				/* 6/9 (pentium-m "banias"), 6/13 (pentium-m "dothan"), and
				 * 6/14 (core1 "yonah") theoretically support sse2, but it's
				 * usually slower than mmx. */
				if fs.inSet(SSE2) {
					fs.set(SSE2SLOW)
				}
				if fs.inSet(SSE3) {
					fs.set(SSE3SLOW)
				}
			}
			/* The Atom processor has SSSE3 support, which is useful in many cases,
			 * but sometimes the SSSE3 version is slower than the SSE2 equivalent
			 * on the Atom, but is generally faster on other processors supporting
			 * SSSE3. This flag allows for selectively disabling certain SSSE3
			 * functions on the Atom. */
			if family == 6 && model == 28 {
				fs.set(ATOM)
			}
		}
	}
	if maxExtendedFunction() >= 0x80000008 {
		_, b, _, _ := cpuid(0x80000008)
		if (b & (1 << 9)) != 0 {
			fs.set(WBNOINVD)
		}
	}

	return fs
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
func (c CPUInfo) ArmFP() bool {
	return c.Arm&FP != 0
}

// Advanced SIMD
func (c CPUInfo) ArmASIMD() bool {
	return c.Arm&ASIMD != 0
}

// Generic timer
func (c CPUInfo) ArmEVTSTRM() bool {
	return c.Arm&EVTSTRM != 0
}

// AES instructions
func (c CPUInfo) ArmAES() bool {
	return c.Arm&AES != 0
}

// Polynomial Multiply instructions (PMULL/PMULL2)
func (c CPUInfo) ArmPMULL() bool {
	return c.Arm&PMULL != 0
}

// SHA-1 instructions (SHA1C, etc)
func (c CPUInfo) ArmSHA1() bool {
	return c.Arm&SHA1 != 0
}

// SHA-2 instructions (SHA256H, etc)
func (c CPUInfo) ArmSHA2() bool {
	return c.Arm&SHA2 != 0
}

// CRC32/CRC32C instructions
func (c CPUInfo) ArmCRC32() bool {
	return c.Arm&CRC32 != 0
}

// Large System Extensions (LSE)
func (c CPUInfo) ArmATOMICS() bool {
	return c.Arm&ATOMICS != 0
}

// Half-precision floating point
func (c CPUInfo) ArmFPHP() bool {
	return c.Arm&FPHP != 0
}

// Advanced SIMD half-precision floating point
func (c CPUInfo) ArmASIMDHP() bool {
	return c.Arm&ASIMDHP != 0
}

// Rounding Double Multiply Accumulate/Subtract (SQRDMLAH/SQRDMLSH)
func (c CPUInfo) ArmASIMDRDM() bool {
	return c.Arm&ASIMDRDM != 0
}

// Javascript-style double->int convert (FJCVTZS)
func (c CPUInfo) ArmJSCVT() bool {
	return c.Arm&JSCVT != 0
}

// Floatin point complex number addition and multiplication
func (c CPUInfo) ArmFCMA() bool {
	return c.Arm&FCMA != 0
}

// Weaker release consistency (LDAPR, etc)
func (c CPUInfo) ArmLRCPC() bool {
	return c.Arm&LRCPC != 0
}

// Data cache clean to Point of Persistence (DC CVAP)
func (c CPUInfo) ArmDCPOP() bool {
	return c.Arm&DCPOP != 0
}

// SHA-3 instructions (EOR3, RAXI, XAR, BCAX)
func (c CPUInfo) ArmSHA3() bool {
	return c.Arm&SHA3 != 0
}

// SM3 instructions
func (c CPUInfo) ArmSM3() bool {
	return c.Arm&SM3 != 0
}

// SM4 instructions
func (c CPUInfo) ArmSM4() bool {
	return c.Arm&SM4 != 0
}

// SIMD Dot Product
func (c CPUInfo) ArmASIMDDP() bool {
	return c.Arm&ASIMDDP != 0
}

// SHA512 instructions
func (c CPUInfo) ArmSHA512() bool {
	return c.Arm&SHA512 != 0
}

// Scalable Vector Extension
func (c CPUInfo) ArmSVE() bool {
	return c.Arm&SVE != 0
}

// Generic Pointer Authentication
func (c CPUInfo) ArmGPA() bool {
	return c.Arm&GPA != 0
}
