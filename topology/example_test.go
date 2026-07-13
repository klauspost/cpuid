// Copyright (c) 2026 Klaus Post, released under MIT License. See LICENSE file.

package topology_test

import (
	"fmt"

	"github.com/klauspost/cpuid/v2"
	"github.com/klauspost/cpuid/v2/internal/samplecpu"
	"github.com/klauspost/cpuid/v2/topology"
)

// Example shows how to obtain and inspect the topology of the running machine.
func Example() {
	// ScanTopology pins the calling goroutine to every online CPU to build an
	// exact tree. For a cheaper approximation that does no pinning, use
	// cpuid.CPU.Topology() instead. On unsupported platforms ScanTopology
	// returns a best-effort snapshot together with a non-nil error.
	sys, _ := cpuid.ScanTopology()

	fmt.Printf("%s: %d package(s), %d cores, %d threads\n",
		sys.Vendor, sys.NumPackages(), sys.NumCores(), sys.NumThreads())
	fmt.Printf("performance cores: %d, efficiency cores: %d\n",
		sys.CoresByType()[topology.Performance], sys.CoresByType()[topology.Efficiency])
	fmt.Printf("total L3: %d bytes\n", sys.TotalCache(3, topology.Unified))

	// Output varies by machine, so it is not checked here. The method examples
	// below use a fixed synthetic topology for reproducible output.
}

// ExampleSystem_TotalCache totals a cache level below any node: the whole
// system, or a single group (CCD/tile). Each physical instance is counted once.
func ExampleSystem_TotalCache() {
	sys := samplecpu.Topology()
	fmt.Printf("L3 total:       %d MB\n", sys.TotalCache(3, topology.Unified)>>20)
	fmt.Printf("L3 in group #0: %d MB\n", sys.Packages[0].Groups[0].TotalCache(3, topology.Unified)>>20)
	fmt.Printf("L2 total:       %d MB\n", sys.TotalCache(2, topology.Unified)>>20)
	// Output:
	// L3 total:       32 MB
	// L3 in group #0: 16 MB
	// L2 total:       6 MB
}

// ExampleSystem_CoresByType counts cores per type, e.g. to tell how many
// performance and efficiency cores a hybrid CPU has.
func ExampleSystem_CoresByType() {
	sys := samplecpu.Topology()
	byType := sys.CoresByType()
	fmt.Printf("performance: %d\n", byType[topology.Performance])
	fmt.Printf("efficiency:  %d\n", byType[topology.Efficiency])
	fmt.Printf("threads:     %d\n", sys.NumThreads())
	// Output:
	// performance: 2
	// efficiency:  4
	// threads:     8
}

// ExampleSystem_String prints the topology as an indented tree.
func ExampleSystem_String() {
	fmt.Print(samplecpu.Topology())
	// Output:
	// System: GenuineIntel (scanned) — 1 package(s), 6 core(s), 8 thread(s)
	//   Package #0
	//     Group #0  [L3 16MB]
	//       Core #0 (Performance) threads: 0, 1  [L1d 32KB L1i 32KB L2 1MB]
	//       Core #1 (Performance) threads: 2, 3  [L1d 32KB L1i 32KB L2 1MB]
	//     Group #1  [L3 16MB]
	//       Core #2 (Efficiency) threads: 4  [L1d 32KB L1i 32KB L2 1MB]
	//       Core #3 (Efficiency) threads: 5  [L1d 32KB L1i 32KB L2 1MB]
	//       Core #4 (Efficiency) threads: 6  [L1d 32KB L1i 32KB L2 1MB]
	//       Core #5 (Efficiency) threads: 7  [L1d 32KB L1i 32KB L2 1MB]
	//   Cores by type: Performance=2 Efficiency=4
	//   Total cache: L1D 192KB, L1I 192KB, L2 6MB, L3 32MB
}
