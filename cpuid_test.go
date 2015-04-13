// Copyright (c) 2015 Klaus Post, released under MIT License. See LICENSE file.

package cpuid

import (
	"fmt"
	"testing"
)

// There is no real way to test a CPU identifier, since results will
// obviously differ on each machine.
func TestCPUID(t *testing.T) {
	n := maxExtendedFunction()
	t.Logf("MAX:%x\n", n)
	t.Log("Name:", CPU.BrandName)
	t.Log("PhysicalCores:", CPU.PhysicalCores)
	t.Log("ThreadsPerCore:", CPU.ThreadsPerCore)
	t.Log("LogicalCores:", CPU.LogicalCores)
	t.Log("Family", CPU.Family, "Mt odel:", CPU.Model)
	t.Log("Features:", CPU.Features)
	t.Log("Cacheline bytes:", CPU.CacheLine)
	if CPU.SSE2() {
		t.Log("Yay - we have SSE 2")
	}
}

func Example() {
	// Print basic CPU information:
	fmt.Println("Name:", CPU.BrandName)
	fmt.Println("PhysicalCores:", CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", CPU.LogicalCores)
	fmt.Println("Family", CPU.Family, "Model:", CPU.Model)
	fmt.Println("Features:", CPU.Features)
	fmt.Println("Cacheline bytes:", CPU.CacheLine)

	// Test if we have a specific feature:
	if CPU.SSE() {
		fmt.Println("We have Streaming SIMD Extensions")
	}
}

// Generated here: http://play.golang.org/p/31X7Inkrpp

// TestCmov tests Cmov() function
func TestCmov(t *testing.T) {
	got := CPU.Cmov()
	expected := CPU.Features&CMOV == CMOV
	if got != expected {
		t.Fatalf("Cmov: expected %v, got %v", expected, got)
	}
	t.Log("CMOV Support:", got)
}

// TestAmd3dnow tests Amd3dnow() function
func TestAmd3dnow(t *testing.T) {
	got := CPU.Amd3dnow()
	expected := CPU.Features&AMD3DNOW == AMD3DNOW
	if got != expected {
		t.Fatalf("Amd3dnow: expected %v, got %v", expected, got)
	}
	t.Log("AMD3DNOW Support:", got)
}

// TestAmd3dnowExt tests Amd3dnowExt() function
func TestAmd3dnowExt(t *testing.T) {
	got := CPU.Amd3dnowExt()
	expected := CPU.Features&AMD3DNOWEXT == AMD3DNOWEXT
	if got != expected {
		t.Fatalf("Amd3dnowExt: expected %v, got %v", expected, got)
	}
	t.Log("AMD3DNOWEXT Support:", got)
}

// TestMMX tests MMX() function
func TestMMX(t *testing.T) {
	got := CPU.MMX()
	expected := CPU.Features&MMX == MMX
	if got != expected {
		t.Fatalf("MMX: expected %v, got %v", expected, got)
	}
	t.Log("MMX Support:", got)
}

// TestMMXext tests MMXext() function
func TestMMXext(t *testing.T) {
	got := CPU.MMXExt()
	expected := CPU.Features&MMXEXT == MMXEXT
	if got != expected {
		t.Fatalf("MMXExt: expected %v, got %v", expected, got)
	}
	t.Log("MMXEXT Support:", got)
}

// TestSSE tests SSE() function
func TestSSE(t *testing.T) {
	got := CPU.SSE()
	expected := CPU.Features&SSE == SSE
	if got != expected {
		t.Fatalf("SSE: expected %v, got %v", expected, got)
	}
	t.Log("SSE Support:", got)
}

// TestSSE2 tests SSE2() function
func TestSSE2(t *testing.T) {
	got := CPU.SSE2()
	expected := CPU.Features&SSE2 == SSE2
	if got != expected {
		t.Fatalf("SSE2: expected %v, got %v", expected, got)
	}
	t.Log("SSE2 Support:", got)
}

// TestSSE3 tests SSE3() function
func TestSSE3(t *testing.T) {
	got := CPU.SSE3()
	expected := CPU.Features&SSE3 == SSE3
	if got != expected {
		t.Fatalf("SSE3: expected %v, got %v", expected, got)
	}
	t.Log("SSE3 Support:", got)
}

// TestSSSE3 tests SSSE3() function
func TestSSSE3(t *testing.T) {
	got := CPU.SSSE3()
	expected := CPU.Features&SSSE3 == SSSE3
	if got != expected {
		t.Fatalf("SSSE3: expected %v, got %v", expected, got)
	}
	t.Log("SSSE3 Support:", got)
}

// TestSSE4 tests SSE4() function
func TestSSE4(t *testing.T) {
	got := CPU.SSE4()
	expected := CPU.Features&SSE4 == SSE4
	if got != expected {
		t.Fatalf("SSE4: expected %v, got %v", expected, got)
	}
	t.Log("SSE4 Support:", got)
}

// TestSSE42 tests SSE42() function
func TestSSE42(t *testing.T) {
	got := CPU.SSE42()
	expected := CPU.Features&SSE42 == SSE42
	if got != expected {
		t.Fatalf("SSE42: expected %v, got %v", expected, got)
	}
	t.Log("SSE42 Support:", got)
}

// TestAVX tests AVX() function
func TestAVX(t *testing.T) {
	got := CPU.AVX()
	expected := CPU.Features&AVX == AVX
	if got != expected {
		t.Fatalf("AVX: expected %v, got %v", expected, got)
	}
	t.Log("AVX Support:", got)
}

// TestAVX2 tests AVX2() function
func TestAVX2(t *testing.T) {
	got := CPU.AVX2()
	expected := CPU.Features&AVX2 == AVX2
	if got != expected {
		t.Fatalf("AVX2: expected %v, got %v", expected, got)
	}
	t.Log("AVX2 Support:", got)
}

// TestFMA3 tests FMA3() function
func TestFMA3(t *testing.T) {
	got := CPU.FMA3()
	expected := CPU.Features&FMA3 == FMA3
	if got != expected {
		t.Fatalf("FMA3: expected %v, got %v", expected, got)
	}
	t.Log("FMA3 Support:", got)
}

// TestFMA4 tests FMA4() function
func TestFMA4(t *testing.T) {
	got := CPU.FMA4()
	expected := CPU.Features&FMA4 == FMA4
	if got != expected {
		t.Fatalf("FMA4: expected %v, got %v", expected, got)
	}
	t.Log("FMA4 Support:", got)
}

// TestXOP tests XOP() function
func TestXOP(t *testing.T) {
	got := CPU.XOP()
	expected := CPU.Features&XOP == XOP
	if got != expected {
		t.Fatalf("XOP: expected %v, got %v", expected, got)
	}
	t.Log("XOP Support:", got)
}

// TestBMI1 tests BMI1() function
func TestBMI1(t *testing.T) {
	got := CPU.BMI1()
	expected := CPU.Features&BMI1 == BMI1
	if got != expected {
		t.Fatalf("BMI1: expected %v, got %v", expected, got)
	}
	t.Log("BMI1 Support:", got)
}

// TestBMI2 tests BMI2() function
func TestBMI2(t *testing.T) {
	got := CPU.BMI2()
	expected := CPU.Features&BMI2 == BMI2
	if got != expected {
		t.Fatalf("BMI2: expected %v, got %v", expected, got)
	}
	t.Log("BMI2 Support:", got)
}

// TestTBM tests TBM() function
func TestTBM(t *testing.T) {
	got := CPU.TBM()
	expected := CPU.Features&TBM == TBM
	if got != expected {
		t.Fatalf("TBM: expected %v, got %v", expected, got)
	}
	t.Log("TBM Support:", got)
}

// TestLzcnt tests Lzcnt() function
func TestLzcnt(t *testing.T) {
	got := CPU.Lzcnt()
	expected := CPU.Features&LZCNT == LZCNT
	if got != expected {
		t.Fatalf("Lzcnt: expected %v, got %v", expected, got)
	}
	t.Log("LZCNT Support:", got)
}

// TestAesNi tests AesNi() function
func TestAesNi(t *testing.T) {
	got := CPU.AesNi()
	expected := CPU.Features&AESNI == AESNI
	if got != expected {
		t.Fatalf("AesNi: expected %v, got %v", expected, got)
	}
	t.Log("AESNI Support:", got)
}

// TestClmul tests Clmul() function
func TestClmul(t *testing.T) {
	got := CPU.Clmul()
	expected := CPU.Features&CLMUL == CLMUL
	if got != expected {
		t.Fatalf("Clmul: expected %v, got %v", expected, got)
	}
	t.Log("CLMUL Support:", got)
}

// TestSSE2Slow tests SSE2Slow() function
func TestSSE2Slow(t *testing.T) {
	got := CPU.SSE2Slow()
	expected := CPU.Features&SSE2SLOW == SSE2SLOW
	if got != expected {
		t.Fatalf("SSE2Slow: expected %v, got %v", expected, got)
	}
	t.Log("SSE2SLOW Support:", got)
}

// TestSSE3Slow tests SSE3slow() function
func TestSSE3Slow(t *testing.T) {
	got := CPU.SSE3Slow()
	expected := CPU.Features&SSE3SLOW == SSE3SLOW
	if got != expected {
		t.Fatalf("SSE3slow: expected %v, got %v", expected, got)
	}
	t.Log("SSE3SLOW Support:", got)
}

// TestAtom tests Atom() function
func TestAtom(t *testing.T) {
	got := CPU.Atom()
	expected := CPU.Features&ATOM == ATOM
	if got != expected {
		t.Fatalf("Atom: expected %v, got %v", expected, got)
	}
	t.Log("ATOM Support:", got)
}
