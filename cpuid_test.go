// Copyright (c) 2015 Klaus Post, released under MIT License. See LICENSE file.

package cpuid

import (
	"fmt"
	"strings"
	"testing"
)

func TestLastID(t *testing.T) {
	if lastID.String() != "lastID" {
		t.Fatal("stringer not updated, run go generate")
	}
}

func TestLastVendorID(t *testing.T) {
	if lastVendor.String() != "lastVendor" {
		t.Fatal("stringer not updated, run go generate")
	}
}

// There is no real way to test a CPU identifier, since results will
// obviously differ on each machine.
func TestCPUID(t *testing.T) {
	Detect()
	n := maxFunctionID()
	t.Logf("Max Function:0x%x", n)
	n = maxExtendedFunction()
	t.Logf("Max Extended Function:0x%x", n)
	t.Log("VendorString:", CPU.VendorString)
	t.Log("VendorID:", CPU.VendorID)
	t.Log("Name:", CPU.BrandName)
	t.Log("PhysicalCores:", CPU.PhysicalCores)
	t.Log("ThreadsPerCore:", CPU.ThreadsPerCore)
	t.Log("LogicalCores:", CPU.LogicalCores)
	t.Log("Family", CPU.Family, "Model:", CPU.Model)
	t.Log("Features:", fmt.Sprintf(strings.Join(CPU.FeatureSet(), ",")))
	t.Log("Cacheline bytes:", CPU.CacheLine)
	t.Log("L1 Instruction Cache:", CPU.Cache.L1I, "bytes")
	t.Log("L1 Data Cache:", CPU.Cache.L1D, "bytes")
	t.Log("L2 Cache:", CPU.Cache.L2, "bytes")
	t.Log("L3 Cache:", CPU.Cache.L3, "bytes")
	t.Log("Hz:", CPU.Hz, "Hz")
	t.Log("VM:", CPU.VM())
	t.Log("BoostFreq:", CPU.BoostFreq, "Hz")
}

func TestExample(t *testing.T) {
	Detect()
	// Print basic CPU information:
	fmt.Println("Name:", CPU.BrandName)
	fmt.Println("PhysicalCores:", CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", CPU.LogicalCores)
	fmt.Println("Family", CPU.Family, "Model:", CPU.Model, "Vendor ID:", CPU.VendorID)
	fmt.Println("Features:", fmt.Sprintf(strings.Join(CPU.FeatureSet(), ",")))
	fmt.Println("Cacheline bytes:", CPU.CacheLine)
	fmt.Println("L1 Data Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", CPU.Hz, "hz")

	// Test if we have these specific features:
	if CPU.Supports(SSE, SSE2) {
		fmt.Println("We have Streaming SIMD 2 Extensions")
	}
}
func TestDumpCPUID(t *testing.T) {
	n := int(maxFunctionID())
	for i := 0; i <= n; i++ {
		a, b, c, d := cpuidex(uint32(i), 0)
		t.Logf("CPUID %08x: %08x-%08x-%08x-%08x", i, a, b, c, d)
		ex := uint32(1)
		for {
			a2, b2, c2, d2 := cpuidex(uint32(i), ex)
			if a2 == a && b2 == b && d2 == d || ex > 50 || a2 == 0 {
				break
			}
			t.Logf("CPUID %08x: %08x-%08x-%08x-%08x", i, a2, b2, c2, d2)
			a, b, c, d = a2, b2, c2, d2
			ex++
		}
	}
	n2 := maxExtendedFunction()
	for i := uint32(0x80000000); i <= n2; i++ {
		a, b, c, d := cpuid(i)
		t.Logf("CPUID %08x: %08x-%08x-%08x-%08x", i, a, b, c, d)
	}
}

func Example() {
	// Print basic CPU information:
	fmt.Println("Name:", CPU.BrandName)
	fmt.Println("PhysicalCores:", CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", CPU.LogicalCores)
	fmt.Println("Family", CPU.Family, "Model:", CPU.Model)
	fmt.Println("Features:", CPU.FeatureSet())
	fmt.Println("Cacheline bytes:", CPU.CacheLine)
}

func TestBrandNameZero(t *testing.T) {
	if len(CPU.BrandName) > 0 {
		// Cut out last byte
		last := []byte(CPU.BrandName[len(CPU.BrandName)-1:])
		if last[0] == 0 {
			t.Fatal("last byte was zero")
		} else if last[0] == 32 {
			t.Fatal("whitespace wasn't trimmed")
		}
	}
}

// TestSGX tests SGX detection
func TestSGX(t *testing.T) {
	got := CPU.SGX.Available
	expected := CPU.featureSet.inSet(SGX)
	if got != expected {
		t.Fatalf("SGX: expected %v, got %v", expected, got)
	}
	t.Log("SGX Support:", got)

	if CPU.SGX.Available {
		var total uint64 = 0
		leaves := false
		for _, s := range CPU.SGX.EPCSections {
			t.Logf("SGX EPC section: base address 0x%x, size %v", s.BaseAddress, s.EPCSize)
			total += s.EPCSize
			leaves = true
		}
		if leaves && total == 0 {
			t.Fatal("SGX enabled without any available EPC memory")
		}
	}
}

func TestHas(t *testing.T) {
	Detect()
	defer Detect()
	feats := CPU.FeatureSet()
	for _, feat := range feats {
		f := ParseFeature(feat)
		if f == UNKNOWN {
			t.Error("Got unknown feature:", feat)
			continue
		}
		if !CPU.Has(f) {
			t.Error("CPU.Has returned false, want true")
		}
		if !CPU.Supports(f) {
			t.Error("CPU.Supports returned false, want true")
		}
		// Disable it.
		CPU.Disable(f)
		if CPU.Has(f) {
			t.Error("CPU.Has returned true, want false")
		}
		if CPU.Supports(f) {
			t.Error("CPU.Supports returned true, want false")
		}
		// Reenable
		CPU.Enable(f)
		if !CPU.Has(f) {
			t.Error("CPU.Has returned false, want true")
		}
		if !CPU.Supports(f) {
			t.Error("CPU.Supports returned false, want true")
		}
	}
}

// TestSGXLC tests SGX Launch Control detection
func TestSGXLC(t *testing.T) {
	got := CPU.SGX.LaunchControl
	expected := CPU.featureSet.inSet(SGXLC)
	if got != expected {
		t.Fatalf("SGX: expected %v, got %v", expected, got)
	}
	t.Log("SGX Launch Control Support:", got)
}

// Test VM function
func TestVM(t *testing.T) {
	got := CPU.VM()
	expected := CPU.featureSet.inSet(HYPERVISOR)
	if got != expected {
		t.Fatalf("TestVM: expected %v, got %v", expected, got)
	}
	t.Log("TestVM:", got)
}

// Test RTCounter function
func TestRtCounter(t *testing.T) {
	a := CPU.RTCounter()
	b := CPU.RTCounter()
	t.Log("CPU Counter:", a, b, b-a)
}

// Prints the value of Ia32TscAux()
func TestIa32TscAux(t *testing.T) {
	ecx := CPU.Ia32TscAux()
	t.Logf("Ia32TscAux:0x%x\n", ecx)
	if ecx != 0 {
		chip := (ecx & 0xFFF000) >> 12
		core := ecx & 0xFFF
		t.Log("Likely chip, core:", chip, core)
	}
}

func TestThreadsPerCoreNZ(t *testing.T) {
	if CPU.ThreadsPerCore == 0 {
		t.Fatal("threads per core is zero")
	}
}

// Prints the value of LogicalCPU()
func TestLogicalCPU(t *testing.T) {
	t.Log("Currently executing on cpu:", CPU.LogicalCPU())
}

func TestMaxFunction(t *testing.T) {
	expect := maxFunctionID()
	if CPU.maxFunc != expect {
		t.Fatal("Max function does not match, expected", expect, "but got", CPU.maxFunc)
	}
	expect = maxExtendedFunction()
	if CPU.maxExFunc != expect {
		t.Fatal("Max Extended function does not match, expected", expect, "but got", CPU.maxFunc)
	}
}

// This example will calculate the chip/core number on Linux
// Linux encodes numa id (<<12) and core id (8bit) into TSC_AUX.
func ExampleCPUInfo_Ia32TscAux() {
	ecx := CPU.Ia32TscAux()
	if ecx == 0 {
		fmt.Println("Unknown CPU ID")
		return
	}
	chip := (ecx & 0xFFF000) >> 12
	core := ecx & 0xFFF
	fmt.Println("Chip, Core:", chip, core)
}
