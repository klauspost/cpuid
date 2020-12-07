// Copyright (c) 2020 Klaus Post, released under MIT License. See LICENSE file.

//+build arm64
//+build linux android

package cpuid

import (
	"runtime"
	_ "unsafe" //required for go:linkname
)

// HWCAP bits.
const (
	hwcap_FP       = 1 << 0
	hwcap_ASIMD    = 1 << 1
	hwcap_EVTSTRM  = 1 << 2
	hwcap_AES      = 1 << 3
	hwcap_PMULL    = 1 << 4
	hwcap_SHA1     = 1 << 5
	hwcap_SHA2     = 1 << 6
	hwcap_CRC32    = 1 << 7
	hwcap_ATOMICS  = 1 << 8
	hwcap_FPHP     = 1 << 9
	hwcap_ASIMDHP  = 1 << 10
	hwcap_CPUID    = 1 << 11
	hwcap_ASIMDRDM = 1 << 12
	hwcap_JSCVT    = 1 << 13
	hwcap_FCMA     = 1 << 14
	hwcap_LRCPC    = 1 << 15
	hwcap_DCPOP    = 1 << 16
	hwcap_SHA3     = 1 << 17
	hwcap_SM3      = 1 << 18
	hwcap_SM4      = 1 << 19
	hwcap_ASIMDDP  = 1 << 20
	hwcap_SHA512   = 1 << 21
	hwcap_SVE      = 1 << 22
	hwcap_ASIMDFHM = 1 << 23
)

//go:linkname hwcap interval/cpu.HWCap
var hwcap uint

func detectOS(c *CPUInfo) bool {
	// HWCap was populated by the runtime from the auxiliary vector.
	// Use HWCap information since reading aarch64 system registers
	// is not supported in user space on older linux kernels.
	c.featureSet.setIf(isSet(hwcap, hwcap_FP), FP)
	c.featureSet.setIf(isSet(hwcap, hwcap_ASIMD), ASIMD)
	c.featureSet.setIf(isSet(hwcap, hwcap_AES), AESARM)
	c.featureSet.setIf(isSet(hwcap, hwcap_PMULL), PMULL)
	c.featureSet.setIf(isSet(hwcap, hwcap_SHA1), SHA1)
	c.featureSet.setIf(isSet(hwcap, hwcap_SHA2), SHA2)
	c.featureSet.setIf(isSet(hwcap, hwcap_SHA3), SHA3)
	c.featureSet.setIf(isSet(hwcap, hwcap_CRC32), CRC32)
	c.featureSet.setIf(isSet(hwcap, hwcap_CPUID), ARMCPUID)
	c.featureSet.setIf(isSet(hwcap, hwcap_ASIMDDP), ASIMDDP)
	c.featureSet.setIf(isSet(hwcap, hwcap_ASIMDHP), ASIMDHP)
	c.featureSet.setIf(isSet(hwcap, hwcap_ASIMDRDM), ASIMDRDM)
	c.featureSet.setIf(isSet(hwcap, hwcap_DCPOP), DCPOP)
	c.featureSet.setIf(isSet(hwcap, hwcap_EVTSTRM), EVTSTRM)
	c.featureSet.setIf(isSet(hwcap, hwcap_FCMA), FCMA)
	c.featureSet.setIf(isSet(hwcap, hwcap_JSCVT), JSCVT)
	c.featureSet.setIf(isSet(hwcap, hwcap_LRCPC), LRCPC)
	c.featureSet.setIf(isSet(hwcap, hwcap_SHA512), SHA512)
	c.featureSet.setIf(isSet(hwcap, hwcap_SM3), SM3)
	c.featureSet.setIf(isSet(hwcap, hwcap_SM4), SM4)
	c.featureSet.setIf(isSet(hwcap, hwcap_SVE), SVE)

	// The Samsung S9+ kernel reports support for atomics, but not all cores
	// actually support them, resulting in SIGILL. See issue #28431.
	// TODO(elias.naur): Only disable the optimization on bad chipsets on android.
	c.featureSet.setIf(isSet(hwcap, hwcap_ATOMICS) && runtime.GOOS != "android", ATOMICS)

	return hwcap != 0
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}
