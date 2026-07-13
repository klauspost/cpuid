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
