// Copyright (c) 2026 Klaus Post, released under MIT License. See LICENSE file.

package cpuid

import (
	"archive/zip"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/klauspost/cpuid/v2/topology"
)

// gtCPU is the ground truth for one logical processor, parsed from the
// "allcpu: Package P / Core C / Thread T: ..., <type>" header the dumps carry.
type gtCPU struct {
	pkg, core, thread int
	typ               topology.CoreType
	valid             bool
}

type dumpCPU struct {
	fake fakecpuid
	gt   gtCPU
}

func parseAllcpu(s string, gt *gtCPU) {
	if n, _ := fmt.Sscanf(s, "allcpu: Package %d / Core %d / Thread %d", &gt.pkg, &gt.core, &gt.thread); n != 3 {
		return
	}
	gt.valid = true
	switch {
	case strings.Contains(s, "Intel Atom"):
		gt.typ = topology.Efficiency
	case strings.Contains(s, "Intel Core"):
		gt.typ = topology.Performance
	}
}

// parseCPUIDLine handles both dump formats: colon-separated
// ("CPUID 0000000B: aaaa-bbbb-cccc-dddd [SL 00]") and the older tab/space
// form ("CPUID 0000000B  \taaaa-bbbb-cccc-dddd").
func parseCPUIDLine(line string, fake fakecpuid) {
	rest := strings.TrimSpace(strings.TrimPrefix(line, "CPUID"))
	i := strings.IndexAny(rest, " \t:")
	if i < 0 {
		return
	}
	var leaf uint32
	if n, _ := fmt.Sscanf(rest[:i], "%x", &leaf); n != 1 {
		return
	}
	vals := strings.TrimLeft(rest[i:], " \t:")
	if j := strings.IndexByte(vals, '['); j >= 0 { // drop "[SL nn]" / "[name]" annotations
		vals = vals[:j]
	}
	v := make([]uint32, 4)
	if n, _ := fmt.Sscanf(vals, "%x-%x-%x-%x", &v[0], &v[1], &v[2], &v[3]); n != 4 {
		if n, _ := fmt.Sscanf(vals, "%x %x %x %x", &v[0], &v[1], &v[2], &v[3]); n != 4 {
			return
		}
	}
	fake[leaf] = append(fake[leaf], v)
}

// parseDumpBlocks splits an instlatx64/AIDA64 dump into one entry per logical
// processor. Each "CPUID Registers / Logical CPU #N" section becomes a block;
// other sections (MSR, …) are ignored. Dumps without per-CPU sections are
// returned as a single block.
func parseDumpBlocks(def []byte) []dumpCPU {
	lines := strings.Split(string(def), "\n")
	var out []dumpCPU
	var cur *dumpCPU
	for _, line := range lines {
		t := strings.TrimSpace(line)
		switch {
		case strings.Contains(t, "CPUID Registers"): // new CPU section (both formats)
			out = append(out, dumpCPU{fake: fakecpuid{}})
			cur = &out[len(out)-1]
			continue
		case strings.HasPrefix(t, "------[") || (strings.Contains(t, "Registers") && !strings.HasPrefix(t, "CPUID ")):
			cur = nil // MSR/MTRR/other section
			continue
		}
		if cur == nil {
			continue
		}
		switch {
		case strings.HasPrefix(t, "allcpu:"):
			parseAllcpu(t, &cur.gt)
		case strings.HasPrefix(t, "CPUID "):
			parseCPUIDLine(t, cur.fake)
		}
	}
	// Old flat dumps have no per-CPU sections: treat the whole file as one CPU.
	if len(out) == 0 {
		d := dumpCPU{fake: fakecpuid{}}
		for _, line := range lines {
			if t := strings.TrimSpace(line); strings.HasPrefix(t, "CPUID ") {
				parseCPUIDLine(t, d.fake)
			}
		}
		if len(d.fake) > 0 {
			out = append(out, d)
		}
	}
	// Drop empty blocks (e.g. a header with no register lines).
	filtered := out[:0]
	for _, d := range out {
		if len(d.fake) > 0 {
			filtered = append(filtered, d)
		}
	}
	return filtered
}

// installMock points the package CPUID functions at a single dump block.
func installMock(fake fakecpuid) func() {
	save := idfuncs{cpuid: cpuid, cpuidex: cpuidex, xgetbv: xgetbv}
	cpuid = func(op uint32) (a, b, c, d uint32) {
		if e := fake[op]; len(e) > 0 {
			return e[0][0], e[0][1], e[0][2], e[0][3]
		}
		return 0, 0, 0, 0
	}
	cpuidex = func(op, op2 uint32) (a, b, c, d uint32) {
		if e := fake[op]; int(op2) < len(e) {
			return e[op2][0], e[op2][1], e[op2][2], e[op2][3]
		}
		return 0, 0, 0, 0
	}
	xgetbv = func(uint32) (uint32, uint32) { return 0, 0 }
	return func() { cpuid, cpuidex, xgetbv = save.cpuid, save.cpuidex, save.xgetbv }
}

func buildTopologyFromDump(d []dumpCPU) ([]logicalCPU, *topology.System) {
	recs := make([]logicalCPU, 0, len(d))
	for i := range d {
		restore := installMock(d[i].fake)
		recs = append(recs, parseLogicalCPU())
		restore()
	}
	return recs, assembleTopology(recs)
}

// uniqueAPICs reports whether every record has a distinct APIC id. Some
// recorded dumps were captured without per-CPU affinity and repeat one id, so
// their real core/package layout cannot be reconstructed.
func uniqueAPICs(recs []logicalCPU) bool {
	seen := make(map[uint32]bool, len(recs))
	for _, r := range recs {
		if seen[r.apicID] {
			return false
		}
		seen[r.apicID] = true
	}
	return true
}

// TestTopologyDumps replays every recorded CPUID dump through the topology
// builder and checks the result against the dumps' own ground-truth headers.
func TestTopologyDumps(t *testing.T) {
	zr, err := zip.OpenReader("testdata/cpuid_data.zip")
	if err != nil {
		t.Skip("No testdata:", err)
	}
	defer zr.Close()
	defer Detect()

	for _, f := range zr.File {
		if !strings.HasSuffix(f.Name, ".txt") {
			continue
		}
		t.Run(filepath.Base(f.Name), func(t *testing.T) {
			rc, err := f.Open()
			if err != nil {
				t.Fatal(err)
			}
			content, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				t.Fatal(err)
			}
			cpus := parseDumpBlocks(content)
			if len(cpus) == 0 {
				t.Skip("no CPUID data")
			}
			recs, sys := buildTopologyFromDump(cpus)

			// One thread per recorded logical CPU, always.
			if got := sys.NumThreads(); got != len(cpus) {
				t.Errorf("NumThreads()=%d, want %d\n%s", got, len(cpus), sys)
			}
			// No cache instance may be counted twice: the sum over the tree
			// must equal the sum of every core-private plus group-shared cache.
			checkCacheOnce(t, sys)

			if !uniqueAPICs(recs) {
				t.Skip("dump repeats APIC ids (captured without per-CPU affinity)")
			}

			// Compare against ground-truth headers when present on every block.
			pkgs := map[int]bool{}
			cores := map[[2]int]bool{}
			typeOf := map[[2]int]topology.CoreType{}
			pos := map[[3]int]bool{}
			typed := false
			for i := range cpus {
				gt := cpus[i].gt
				if !gt.valid {
					return // no ground truth for this dump
				}
				tri := [3]int{gt.pkg, gt.core, gt.thread}
				if pos[tri] {
					t.Skipf("dump repeats logical position %v (inconsistent ground truth)", tri)
				}
				pos[tri] = true
				pkgs[gt.pkg] = true
				key := [2]int{gt.pkg, gt.core}
				cores[key] = true
				if gt.typ != topology.Unknown {
					typed = true
					typeOf[key] = gt.typ
				}
			}
			if got := sys.NumPackages(); got != len(pkgs) {
				t.Errorf("NumPackages()=%d, want %d", got, len(pkgs))
			}
			if got := sys.NumCores(); got != len(cores) {
				t.Errorf("NumCores()=%d, want %d\n%s", got, len(cores), sys)
			}
			if typed {
				want := map[topology.CoreType]int{}
				for _, ct := range typeOf {
					want[ct]++
				}
				got := sys.CoresByType()
				for ct, n := range want {
					if got[ct] != n {
						t.Errorf("CoresByType()[%v]=%d, want %d", ct, got[ct], n)
					}
				}
			}
		})
	}
}

// checkCacheOnce verifies TotalCache equals the independent sum of all caches
// stored in the tree, i.e. that no instance is attached (and counted) twice.
func checkCacheOnce(t *testing.T, sys *topology.System) {
	t.Helper()
	type lk struct {
		level int
		typ   topology.CacheType
	}
	want := map[lk]int{}
	var walk func(cs []topology.Cache)
	walk = func(cs []topology.Cache) {
		for _, c := range cs {
			want[lk{c.Level, c.Type}] += c.Size
		}
	}
	for pi := range sys.Packages {
		p := &sys.Packages[pi]
		walk(p.Caches)
		for gi := range p.Groups {
			g := &p.Groups[gi]
			walk(g.Caches)
			for ci := range g.Cores {
				walk(g.Cores[ci].Caches)
			}
		}
	}
	for k, size := range want {
		if got := sys.TotalCache(k.level, k.typ); got != size {
			t.Errorf("TotalCache(%d,%v)=%d, want %d", k.level, k.typ, got, size)
		}
	}
}

func TestCountsTopologyHybrid(t *testing.T) {
	defer Detect()

	tests := []struct {
		name         string
		physical     int
		logical      int
		tpc          int
		isHybrid     bool
		wantP        int
		wantE        int
		wantTotalTh  int
		wantSharedL3 int
	}{
		{
			name:         "Alder Lake i9 12900K (8P+8E)",
			physical:     16,
			logical:      24,
			tpc:          2,
			isHybrid:     true,
			wantP:        8,
			wantE:        8,
			wantTotalTh:  24,
			wantSharedL3: 24,
		},
		{
			name:         "Raptor Lake i9 13900K (8P+16E)",
			physical:     24,
			logical:      32,
			tpc:          2,
			isHybrid:     true,
			wantP:        8,
			wantE:        16,
			wantTotalTh:  32,
			wantSharedL3: 32,
		},
		{
			name:         "Alder Lake i5 12600K (6P+4E)",
			physical:     10,
			logical:      16,
			tpc:          2,
			isHybrid:     true,
			wantP:        6,
			wantE:        4,
			wantTotalTh:  16,
			wantSharedL3: 16,
		},
		{
			name:         "Alder Lake-N N100 (0P+8E)",
			physical:     8,
			logical:      8,
			tpc:          1,
			isHybrid:     true,
			wantP:        0,
			wantE:        8,
			wantTotalTh:  8,
			wantSharedL3: 8,
		},
		{
			name:         "Hybrid leaf 4 reporting max addressable IDs (64 cores, 24 logical)",
			physical:     64,
			logical:      24,
			tpc:          2,
			isHybrid:     true,
			wantP:        0,
			wantE:        24,
			wantTotalTh:  24,
			wantSharedL3: 24,
		},
		{
			name:         "Asymmetric SMT without HYBRID_CPU flag (12 physical, 16 logical, tpc=2)",
			physical:     12,
			logical:      16,
			tpc:          2,
			isHybrid:     false,
			wantP:        12,
			wantE:        0,
			wantTotalTh:  16,
			wantSharedL3: 16,
		},
		{
			name:         "Homogeneous SMT (8 physical, 16 logical, tpc=2)",
			physical:     8,
			logical:      16,
			tpc:          2,
			isHybrid:     false,
			wantP:        8,
			wantE:        0,
			wantTotalTh:  16,
			wantSharedL3: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CPU.PhysicalCores = tt.physical
			CPU.LogicalCores = tt.logical
			CPU.ThreadsPerCore = tt.tpc
			CPU.VendorString = "GenuineIntel"
			CPU.Cache.L1D = 32768
			CPU.Cache.L1I = 32768
			CPU.Cache.L2 = 524288
			CPU.Cache.L3 = 16777216

			if tt.isHybrid {
				CPU.Enable(HYBRID_CPU)
			} else {
				CPU.Disable(HYBRID_CPU)
			}

			sys := countsTopology()

			if sys.Online != tt.wantTotalTh {
				t.Errorf("sys.Online = %d, want %d", sys.Online, tt.wantTotalTh)
			}
			if got := sys.NumThreads(); got != tt.wantTotalTh {
				t.Errorf("NumThreads() = %d, want %d", got, tt.wantTotalTh)
			}

			seen := make(map[int]bool, tt.wantTotalTh)
			var allThreads []int
			pCores := 0
			eCores := 0
			for _, pkg := range sys.Packages {
				for _, grp := range pkg.Groups {
					for _, c := range grp.Cores {
						switch c.Type {
						case topology.Performance:
							pCores++
							if len(c.Threads) == 2 {
								for _, ca := range c.Caches {
									if ca.SharedThreads != 2 {
										t.Errorf("Core %d (Performance) cache SharedThreads = %d, want 2", c.ID, ca.SharedThreads)
									}
								}
							}
						case topology.Efficiency:
							eCores++
							if len(c.Threads) == 1 {
								for _, ca := range c.Caches {
									if ca.SharedThreads != 1 {
										t.Errorf("Core %d (Efficiency) cache SharedThreads = %d, want 1", c.ID, ca.SharedThreads)
									}
								}
							}
						}
						for _, th := range c.Threads {
							if th.ID < 0 || th.ID >= tt.wantTotalTh {
								t.Errorf("Thread ID %d out of bounds [0, %d)", th.ID, tt.wantTotalTh)
							}
							if seen[th.ID] {
								t.Errorf("Duplicate thread ID %d", th.ID)
							}
							seen[th.ID] = true
							allThreads = append(allThreads, th.ID)
						}
					}
					for _, ca := range grp.Caches {
						if ca.Level == 3 && ca.SharedThreads != tt.wantSharedL3 {
							t.Errorf("Group L3 SharedThreads = %d, want %d", ca.SharedThreads, tt.wantSharedL3)
						}
					}
				}
			}

			if len(allThreads) != tt.wantTotalTh {
				t.Errorf("Total threads collected = %d, want %d", len(allThreads), tt.wantTotalTh)
			}
			if pCores != tt.wantP {
				t.Errorf("Performance cores = %d, want %d", pCores, tt.wantP)
			}
			if eCores != tt.wantE {
				t.Errorf("Efficiency cores = %d, want %d", eCores, tt.wantE)
			}
		})
	}
}
