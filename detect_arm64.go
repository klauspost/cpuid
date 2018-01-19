// Copyright (c) 2018 Robin Lu (lubinsz@gmail.com), released under MIT License. See LICENSE file.
// +build arm64,linux,cgo
package cpuid

/*
#include <sys/auxv.h>

#define HWCAP_CPUID	(1 << 11)

unsigned long gethwcap() {
        return getauxval(AT_HWCAP);
}

static uint64_t getMIDR() {
	register uint64_t midr = UINT64_MAX;
	unsigned long hwcap = getauxval(AT_HWCAP);

	if (hwcap & HWCAP_CPUID)
		asm volatile ("mrs %0, midr_el1" : "=r"(midr));
	else
		midr = 0;

	return (uint64_t)midr;
}
*/
import "C"

const (
	MIDR_PARTNUM_SHIFT      = 4
	MIDR_PARTNUM_MASK       = 0xfff << MIDR_PARTNUM_SHIFT
	MIDR_ARCHITECTURE_SHIFT = 16
	MIDR_ARCHITECTURE_MASK  = 0xf << MIDR_ARCHITECTURE_SHIFT
	MIDR_VARIANT_SHIFT      = 20
	MIDR_VARIANT_MASK       = 0xf << MIDR_VARIANT_SHIFT
	MIDR_IMPLEMENTOR_SHIFT  = 24
	MIDR_IMPLEMENTOR_MASK   = 0xff << MIDR_IMPLEMENTOR_SHIFT
)

type cpuinfo_arm64 struct {
	HwCap   uint64
	MidrReg uint64
	/* reserve */
	ZvaSize uint32
}

var (
	cpuInfo_arm64 cpuinfo_arm64
)

func initArm64() {
	cpuInfo_arm64.HwCap = uint64(C.gethwcap())
	cpuInfo_arm64.MidrReg = uint64(C.getMIDR())
}

/*
 * A little tricky here.
 * There is no such registers of cpuid(eax, ebx, ecx, edx) in Arm64 platform.
 * So we only borrow the name of these functions.
 * In armCpuid(),
 * eax = part
 * ebx = architecture
 * ecx = variant
 * edx = implementor
 */
func armCpuid(op uint32) (eax, ebx, ecx, edx uint32) {
	midr := cpuInfo_arm64.MidrReg

	eax = uint32((midr & MIDR_PARTNUM_MASK) >> MIDR_PARTNUM_SHIFT)
	ebx = uint32((midr & MIDR_ARCHITECTURE_MASK) >> MIDR_ARCHITECTURE_SHIFT)
	ecx = uint32((midr & MIDR_VARIANT_MASK) >> MIDR_VARIANT_SHIFT)
	edx = uint32((midr & MIDR_IMPLEMENTOR_MASK) >> MIDR_IMPLEMENTOR_SHIFT)

	return uint32(eax), uint32(ebx), uint32(ecx), uint32(edx)
}

/*
 * eax = hwcap
 */
func armCpuidex(op, op2 uint32) (eax, ebx, ecx, edx uint32) {
	return uint32(cpuInfo_arm64.HwCap), 0, 0, 0
}

func armXgetbv(index uint32) (eax, edx uint32) {
	return 0, 0
}

func armRdtscpAsm() (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}

func initCPU() {
	/* read hwcap & midr register */
	initArm64()

	cpuid = armCpuid
	cpuidex = armCpuidex

	/* set 0 as default */
	xgetbv = armXgetbv
	rdtscpAsm = armRdtscpAsm
}
