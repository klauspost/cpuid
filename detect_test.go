// Copyright (c) 2026 Klaus Post, released under MIT License. See LICENSE file.

package cpuid

import (
	"archive/zip"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

type syntheticDetector struct {
	leaves    map[uint32][4]uint32
	subleaves map[[2]uint32][4]uint32
	xcr0      uint32
}

func (s *syntheticDetector) CPUID(op uint32) (eax, ebx, ecx, edx uint32) {
	if val, ok := s.leaves[op]; ok {
		return val[0], val[1], val[2], val[3]
	}
	if val, ok := s.subleaves[[2]uint32{op, 0}]; ok {
		return val[0], val[1], val[2], val[3]
	}
	return 0, 0, 0, 0
}

func (s *syntheticDetector) CPUIDX(op, op2 uint32) (eax, ebx, ecx, edx uint32) {
	if val, ok := s.subleaves[[2]uint32{op, op2}]; ok {
		return val[0], val[1], val[2], val[3]
	}
	if op2 == 0 {
		return s.CPUID(op)
	}
	return 0, 0, 0, 0
}

func (s *syntheticDetector) XGetBV(index uint32) (eax, edx uint32) {
	if index == 0 {
		return s.xcr0, 0
	}
	return 0, 0
}

func stringToRegs(s string) [4]uint32 {
	b := []byte(s)
	for len(b) < 16 {
		b = append(b, 0)
	}
	return [4]uint32{
		binary.LittleEndian.Uint32(b[0:4]),
		binary.LittleEndian.Uint32(b[4:8]),
		binary.LittleEndian.Uint32(b[8:12]),
		binary.LittleEndian.Uint32(b[12:16]),
	}
}

func TestDetectExternalX86_SyntheticIntel(t *testing.T) {
	brandPart1 := stringToRegs("Intel(R) Core(TM")
	brandPart2 := stringToRegs(") i7-7700K CPU @")
	brandPart3 := stringToRegs(" 4.20GHz\x00\x00\x00\x00\x00\x00\x00\x00")

	detector := &syntheticDetector{
		leaves: map[uint32][4]uint32{
			// Leaf 0: maxFunc=0x16, "GenuineIntel" (ebx="Genu", edx="ineI", ecx="ntel")
			0: {0x16, 0x756e6547, 0x6c65746e, 0x49656e69},
			// Leaf 1: Family 6, Model 158 (0x9E), Stepping 9
			// eax: (0x9 << 16) | (0x6 << 8) | (0xe << 4) | 0x9 = 0x000906e9
			// ebx: CLFLUSH=8 -> 64 bytes cache line (8*8)
			// ecx: SSE3(0), SSSE3(9), SSE4.1(19), SSE4.2(20), AESNI(25), XSAVE(26), OSXSAVE(27), AVX(28), FMA3(12)
			// edx: FPU(0), CMOV(15), MMX(23), FXSR(24), SSE(25), SSE2(26)
			1: {
				0x000906e9,
				0x08100800,
				(1 << 0) | (1 << 9) | (1 << 12) | (1 << 19) | (1 << 20) | (1 << 25) | (1 << 26) | (1 << 27) | (1 << 28),
				(1 << 0) | (1 << 15) | (1 << 23) | (1 << 24) | (1 << 25) | (1 << 26),
			},
			0x80000000: {0x80000008, 0, 0, 0},
			0x80000002: brandPart1,
			0x80000003: brandPart2,
			0x80000004: brandPart3,
		},
		subleaves: map[[2]uint32][4]uint32{
			// Leaf 7 Subleaf 0: AVX2(bit 5), BMI1(bit 3), BMI2(bit 8)
			{7, 0}: {0, (1 << 3) | (1 << 5) | (1 << 8), 0, 0},
		},
		xcr0: 0x7, // AVX OS support enabled
	}

	info := DetectExternalX86(detector)

	if info.VendorID != Intel {
		t.Fatalf("VendorID mismatch: got %v, want %v", info.VendorID, Intel)
	}
	if info.VendorString != "GenuineIntel" {
		t.Fatalf("VendorString mismatch: got %q, want %q", info.VendorString, "GenuineIntel")
	}
	if info.BrandName != "Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz" {
		t.Fatalf("BrandName mismatch: got %q", info.BrandName)
	}
	if info.Family != 6 || info.Model != 158 || info.Stepping != 9 {
		t.Fatalf("Family/Model/Stepping mismatch: got %d/%d/%d, want 6/158/9", info.Family, info.Model, info.Stepping)
	}
	if info.CacheLine != 64 {
		t.Fatalf("CacheLine mismatch: got %d, want 64", info.CacheLine)
	}
	if info.Hz != 4200000000 {
		t.Fatalf("Hz mismatch: got %d, want 4200000000", info.Hz)
	}

	requiredFeatures := []FeatureID{SSE, SSE2, SSE3, SSE4, SSE42, AVX, AVX2, BMI1, BMI2, AESNI}
	for _, f := range requiredFeatures {
		if !info.Supports(f) {
			t.Fatalf("Expected feature %s to be supported", f.String())
		}
	}
}

func TestDetectExternalX86_SyntheticAMD(t *testing.T) {
	brandPart1 := stringToRegs("AMD Ryzen 9 5950")
	brandPart2 := stringToRegs("X 16-Core Proces")
	brandPart3 := stringToRegs("sor             ")

	detector := &syntheticDetector{
		leaves: map[uint32][4]uint32{
			// Leaf 0: maxFunc=0x10, "AuthenticAMD" (ebx="Auth", edx="enti", ecx="cAMD")
			0: {0x10, 0x68747541, 0x444d4163, 0x69746e65},
			// Leaf 1: Family 25 (0x19), Model 33 (0x21), Stepping 0
			// eax: (0xa << 20) | (0x2 << 16) | (0xf << 8) | (0x1 << 4) | 0x0 = 0x00a20f10
			1: {
				0x00a20f10,
				0x04200800,
				(1 << 0) | (1 << 26) | (1 << 27) | (1 << 28),
				(1 << 0) | (1 << 25) | (1 << 26),
			},
			0x80000000: {0x80000023, 0, 0, 0},
			0x80000002: brandPart1,
			0x80000003: brandPart2,
			0x80000004: brandPart3,
		},
		subleaves: map[[2]uint32][4]uint32{
			{7, 0}: {0, (1 << 3) | (1 << 5) | (1 << 8), 0, 0},
		},
		xcr0: 0x7,
	}

	info := DetectExternalX86(detector)

	if info.VendorID != AMD {
		t.Fatalf("VendorID mismatch: got %v, want %v", info.VendorID, AMD)
	}
	if info.VendorString != "AuthenticAMD" {
		t.Fatalf("VendorString mismatch: got %q, want %q", info.VendorString, "AuthenticAMD")
	}
	if info.BrandName != "AMD Ryzen 9 5950X 16-Core Processor" {
		t.Fatalf("BrandName mismatch: got %q", info.BrandName)
	}
	if info.Family != 25 {
		t.Fatalf("Family mismatch: got %d, want 25", info.Family)
	}
	if !info.Supports(AVX) || !info.Supports(AVX2) {
		t.Fatalf("Expected AVX and AVX2 to be supported")
	}
}

func TestDetectExternalX86_XGetBV_OSDisabled(t *testing.T) {
	// Hardware supports AVX/AVX2 and OSXSAVE, but XGetBV(0) returns 0 (OS did not enable AVX state)
	detector := &syntheticDetector{
		leaves: map[uint32][4]uint32{
			0: {0x16, 0x756e6547, 0x6c65746e, 0x49656e69},
			1: {
				0x000906e9,
				0x08100800,
				(1 << 26) | (1 << 27) | (1 << 28), // XSAVE, OSXSAVE, AVX
				(1 << 25) | (1 << 26),
			},
			0x80000000: {0x80000004, 0, 0, 0},
		},
		subleaves: map[[2]uint32][4]uint32{
			{7, 0}: {0, (1 << 5), 0, 0}, // AVX2
		},
		xcr0: 0, // OS has NOT enabled AVX registers!
	}

	info := DetectExternalX86(detector)

	if info.Supports(AVX) {
		t.Fatalf("AVX should NOT be supported when XGetBV returns 0")
	}
	if info.Supports(AVX2) {
		t.Fatalf("AVX2 should NOT be supported when XGetBV returns 0")
	}
}

func TestDetectExternalX86_Nil(t *testing.T) {
	info := DetectExternalX86(nil)
	if info.VendorID != VendorUnknown {
		t.Fatalf("Expected VendorUnknown for nil detector, got %v", info.VendorID)
	}
	if info.BrandName != "" {
		t.Fatalf("Expected empty BrandName for nil detector, got %q", info.BrandName)
	}
}

type zipMockDetector struct {
	fakeID fakecpuid
}

func (z *zipMockDetector) CPUID(op uint32) (eax, ebx, ecx, edx uint32) {
	first, ok := z.fakeID[op]
	if !ok || len(first) == 0 {
		return 0, 0, 0, 0
	}
	return first[0][0], first[0][1], first[0][2], first[0][3]
}

func (z *zipMockDetector) CPUIDX(op, op2 uint32) (eax, ebx, ecx, edx uint32) {
	first, ok := z.fakeID[op]
	if !ok || int(op2) >= len(first) {
		return 0, 0, 0, 0
	}
	return first[op2][0], first[op2][1], first[op2][2], first[op2][3]
}

func (z *zipMockDetector) XGetBV(index uint32) (eax, edx uint32) {
	first, ok := z.fakeID[1]
	if !ok || len(first) == 0 {
		return 0, 0
	}
	if (first[0][2] & (1 << 26)) == 0 {
		return 0, 0
	}
	return math.MaxUint32, math.MaxUint32
}

func loadDetectorFromZip(t *testing.T, filename string) *zipMockDetector {
	zr, err := zip.OpenReader("testdata/cpuid_data.zip")
	if err != nil {
		t.Skip("No testdata:", err)
		return nil
	}
	defer zr.Close()

	for _, f := range zr.File {
		if filepath.Base(f.Name) == filename {
			rc, err := f.Open()
			if err != nil {
				t.Fatal(err)
			}
			defer rc.Close()
			content, err := io.ReadAll(rc)
			if err != nil {
				t.Fatal(err)
			}
			restore := mockCPU(content)
			fakeCopy := make(fakecpuid)
			for k, v := range fakeIDGlobal(content) {
				fakeCopy[k] = v
			}
			restore()
			return &zipMockDetector{fakeID: fakeCopy}
		}
	}
	t.Fatalf("File %s not found in zip", filename)
	return nil
}

func fakeIDGlobal(content []byte) fakecpuid {
	// Re-parse directly using mockCPU parsing logic
	lines := strings.Split(string(content), "\n")
	anyfound := false
	fakeID := make(fakecpuid)
	for _, line := range lines {
		line = strings.Trim(line, "\r\t ")
		if !strings.HasPrefix(line, "CPUID") {
			continue
		}
		if strings.HasPrefix(line, "CPUID 00000000") {
			if anyfound {
				break
			}
		}
		items := strings.Split(line, ":")
		if len(items) < 2 {
			if len(line) == 51 || len(line) == 50 {
				items = []string{line[0:14], line[15:]}
			} else {
				items = strings.Split(line, "\t")
				if len(items) != 2 {
					continue
				}
			}
		}
		items = items[0:2]
		vals := strings.Trim(items[1], "\r\n ")

		var idV uint32
		n, err := fmt.Sscanf(items[0], "CPUID %x", &idV)
		if err != nil || n != 1 {
			continue
		}
		existing, ok := fakeID[idV]
		if !ok {
			existing = make([][]uint32, 0)
		}

		values := make([]uint32, 4)
		n, err = fmt.Sscanf(vals, "%x-%x-%x-%x", &values[0], &values[1], &values[2], &values[3])
		if n != 4 || err != nil {
			n, err = fmt.Sscanf(vals, "%x %x %x %x", &values[0], &values[1], &values[2], &values[3])
			if n != 4 || err != nil {
				continue
			}
		}

		existing = append(existing, values)
		fakeID[idV] = existing
		anyfound = true
	}
	return fakeID
}

func TestDetectExternalX86_ZipData(t *testing.T) {
	detector := loadDetectorFromZip(t, "GenuineIntel00406E3_Skylake_CPUID.txt")
	if detector == nil {
		return
	}
	info := DetectExternalX86(detector)
	if info.VendorID != Intel {
		t.Fatalf("Expected Intel vendor, got %v", info.VendorID)
	}
	if info.VendorString != "GenuineIntel" {
		t.Fatalf("Expected GenuineIntel, got %s", info.VendorString)
	}
	if info.BrandName == "" {
		t.Fatalf("Expected non-empty BrandName")
	}
	if !info.Supports(AVX2) {
		t.Fatalf("Expected Skylake to support AVX2")
	}
}

func TestDetectExternalX86_Concurrent(t *testing.T) {
	d1 := &syntheticDetector{
		leaves: map[uint32][4]uint32{
			0: {1, 0x756e6547, 0x6c65746e, 0x49656e69},
			1: {0x000906e9, 0x08100800, 0, 0},
		},
	}
	d2 := &syntheticDetector{
		leaves: map[uint32][4]uint32{
			0: {1, 0x68747541, 0x444d4163, 0x69746e65},
			1: {0x00a20f10, 0x04200800, 0, 0},
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			info := DetectExternalX86(d1)
			if info.VendorID != Intel {
				t.Errorf("Concurrent: expected Intel, got %v", info.VendorID)
			}
		}()
		go func() {
			defer wg.Done()
			info := DetectExternalX86(d2)
			if info.VendorID != AMD {
				t.Errorf("Concurrent: expected AMD, got %v", info.VendorID)
			}
		}()
	}
	wg.Wait()
}
