// Copyright (c) 2026 Klaus Post, released under MIT License. See LICENSE file.

package cpuid

import (
	"strings"
	"sync"
)

var detectExternalMu sync.Mutex

// X86Detector provides an interface for external or mock sources of x86 CPUID and XGETBV data.
type X86Detector interface {
	CPUID(op uint32) (eax, ebx, ecx, edx uint32)
	CPUIDX(op, op2 uint32) (eax, ebx, ecx, edx uint32)
	XGetBV(index uint32) (eax, edx uint32)
}

// DetectExternalX86 returns CPUInfo for an x86 CPU by querying the provided X86Detector.
func DetectExternalX86(d X86Detector) CPUInfo {
	if d == nil {
		return CPUInfo{}
	}

	detectExternalMu.Lock()
	defer detectExternalMu.Unlock()

	oldCPUID := cpuid
	oldCPUIDX := cpuidex
	oldXGetBV := xgetbv
	oldDarwin := darwinHasAVX512

	defer func() {
		cpuid = oldCPUID
		cpuidex = oldCPUIDX
		xgetbv = oldXGetBV
		darwinHasAVX512 = oldDarwin
	}()

	var lastXGetBVEax uint32
	cpuid = d.CPUID
	cpuidex = d.CPUIDX
	xgetbv = func(index uint32) (eax, edx uint32) {
		eax, edx = d.XGetBV(index)
		if index == 0 {
			lastXGetBVEax = eax
		}
		return eax, edx
	}
	darwinHasAVX512 = func() bool {
		return (lastXGetBVEax>>5)&7 == 7 && (lastXGetBVEax>>1)&3 == 3
	}

	var c CPUInfo
	c.ThreadsPerCore = 1
	c.Cache.L1I = -1
	c.Cache.L1D = -1
	c.Cache.L2 = -1
	c.Cache.L3 = -1

	addInfoX86(&c)

	if disableFlag != nil {
		s := strings.SplitSeq(*disableFlag, ",")
		for feat := range s {
			feat := ParseFeature(strings.TrimSpace(feat))
			if feat != UNKNOWN {
				c.featureSet.unset(feat)
			}
		}
	}

	return c
}

func addInfoX86(c *CPUInfo) {
	c.maxFunc = maxFunctionID()
	c.maxExFunc = maxExtendedFunction()
	c.BrandName = brandName()
	c.CacheLine = cacheLine()
	c.Family, c.Model, c.Stepping = familyModel()
	c.featureSet = support()
	c.SGX = hasSGX(c.featureSet.inSet(SGX), c.featureSet.inSet(SGXLC))
	c.AMDMemEncryption = hasAMDMemEncryption(c.featureSet.inSet(SME) || c.featureSet.inSet(SEV))
	c.ThreadsPerCore = threadsPerCore()
	c.LogicalCores = logicalCores()
	c.PhysicalCores = physicalCores()
	c.VendorID, c.VendorString = vendorID()
	c.HypervisorVendorID, c.HypervisorVendorString = hypervisorVendorID()
	c.AVX10Level = c.supportAVX10()
	c.cacheSize()
	c.frequencies()
	if c.maxFunc >= 0x0A {
		eax, ebx, _, edx := cpuid(0x0A)
		c.PMU = parseLeaf0AH(c, eax, ebx, edx)
	}
}
