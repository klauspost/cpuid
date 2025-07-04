// Copyright (c) 2015 Klaus Post, released under MIT License. See LICENSE file.

//go:build (386 && !gccgo && !noasm && !appengine) || (amd64 && !gccgo && !noasm && !appengine)
// +build 386,!gccgo,!noasm,!appengine amd64,!gccgo,!noasm,!appengine

package cpuid

func asmCpuid(op uint32) (eax, ebx, ecx, edx uint32)
func asmCpuidex(op, op2 uint32) (eax, ebx, ecx, edx uint32)
func asmXgetbv(index uint32) (eax, edx uint32)
func asmRdtscpAsm() (eax, ebx, ecx, edx uint32)
func asmDarwinHasAVX512() bool

func initCPU() {
	cpuid = asmCpuid
	cpuidex = asmCpuidex
	xgetbv = asmXgetbv
	rdtscpAsm = asmRdtscpAsm
	darwinHasAVX512 = asmDarwinHasAVX512
}

func addInfo(c *CPUInfo, safe bool) {
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
		c.PMU = parseLeaf0AH(eax, ebx, edx)
	}
}

func getVectorLength() (vl, pl uint64) { return 0, 0 }

func parseLeaf0AH(eax, ebx, edx uint32) PerformanceMonitoringInfo {
	var info PerformanceMonitoringInfo

	info.VersionID = uint8(eax & 0xFF)
	info.NumGPCounters = uint8((eax >> 8) & 0xFF)
	info.GPPMCWidth = uint8((eax >> 16) & 0xFF)

	info.RawEBX = ebx
	info.RawEAX = eax
	info.RawEDX = edx
	info.SupportedFixedEvents = []FeatureID{}

	if info.VersionID > 1 { // This information is only valid if VersionID > 1
		info.NumFixedPMC = uint8(edx & 0x1F)          // Bits 4:0
		info.FixedPMCWidth = uint8((edx >> 5) & 0xFF) // Bits 12:5
	}
	if info.VersionID > 0 {
		// first 4 fixed events are always instructions retired, cycles, ref cycles and topdown slots
		if ebx == 0x0 && info.NumFixedPMC == 3 {
			info.SupportedFixedEvents = []FeatureID{PMU_FixedCounter_Instructions, PMU_FixedCounter_Cycles, PMU_FixedCounter_RefCycles}
		}
		if ebx == 0x0 && info.NumFixedPMC == 4 {
			info.SupportedFixedEvents = []FeatureID{PMU_FixedCounter_Instructions, PMU_FixedCounter_Cycles, PMU_FixedCounter_RefCycles, PMU_FixedCounter_Topdown_Slots}
		}
		if ebx != 0x0 {
			if ((ebx >> 0) & 1) == 0 {
				info.SupportedFixedEvents = append(info.SupportedFixedEvents, PMU_FixedCounter_Instructions)
			}
			if ((ebx >> 1) & 1) == 0 {
				info.SupportedFixedEvents = append(info.SupportedFixedEvents, PMU_FixedCounter_Cycles)
			}
			if ((ebx >> 2) & 1) == 0 {
				info.SupportedFixedEvents = append(info.SupportedFixedEvents, PMU_FixedCounter_RefCycles)
			}
			if ((ebx >> 3) & 1) == 0 {
				info.SupportedFixedEvents = append(info.SupportedFixedEvents, PMU_FixedCounter_Topdown_Slots)
			}
		}
	}

	return info
}
