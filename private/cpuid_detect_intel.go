// Generated, DO NOT EDIT,
// but copy it to your own project and rename the package.
// See more at http://github.com/klauspost/cpuid

//+build 386,!gccgo,!noasm amd64,!gccgo,!noasm,!appengine

package cpuid

func asmCpuid(op uint32) (eax, ebx, ecx, edx uint32)
func asmCpuidex(op, op2 uint32) (eax, ebx, ecx, edx uint32)
func asmXgetbv(index uint32) (eax, edx uint32)
func asmRdtscpAsm() (eax, ebx, ecx, edx uint32)

func initCPU() {
	cpuid = asmCpuid
	cpuidex = asmCpuidex
	xgetbv = asmXgetbv
	rdtscpAsm = asmRdtscpAsm
}

func addInfo(c *cpuInfo) {
	c.maxFunc = maxFunctionID()
	c.maxExFunc = maxExtendedFunction()
	c.brandname = brandName()
	c.cacheline = cacheLine()
	c.family, c.model = familyModel()
	c.features = support()
	c.sgx = hasSGX(c.features&sgx != 0, c.features&sgxlc != 0)
	c.threadspercore = threadsPerCore()
	c.logicalcores = logicalCores()
	c.physicalcores = physicalCores()
	c.vendorid, c.vendorstring = vendorID()
	c.hz = hertz(c.brandname)
	c.cacheSize()
}
