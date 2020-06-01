// Generated, DO NOT EDIT,
// but copy it to your own project and rename the package.
// See more at http://github.com/klauspost/cpuid

//+build !amd64,!386,!arm64 gccgo noasm appengine

package cpuid

func initCPU() {
	cpuid = func(uint32) (a, b, c, d uint32) { return 0, 0, 0, 0 }
	cpuidex = func(x, y uint32) (a, b, c, d uint32) { return 0, 0, 0, 0 }
	xgetbv = func(uint32) (a, b uint32) { return 0, 0 }
	rdtscpAsm = func() (a, b, c, d uint32) { return 0, 0, 0, 0 }
}

func addInfo(info *cpuInfo) {}
