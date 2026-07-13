// Copyright (c) 2026 Klaus Post, released under MIT License. See LICENSE file.

package topology_test

import (
	"testing"

	"github.com/klauspost/cpuid/v2/internal/samplecpu"
	"github.com/klauspost/cpuid/v2/topology"
)

func TestAggregates(t *testing.T) {
	s := samplecpu.Topology()
	if got := s.NumPackages(); got != 1 {
		t.Errorf("NumPackages=%d want 1", got)
	}
	if got := s.NumCores(); got != 6 {
		t.Errorf("NumCores=%d want 6", got)
	}
	if got := s.NumThreads(); got != 8 {
		t.Errorf("NumThreads=%d want 8", got)
	}
	// Two L3 instances of 16MB, counted once each.
	if got := s.TotalCache(3, topology.Unified); got != 32<<20 {
		t.Errorf("TotalCache(L3)=%d want %d", got, 32<<20)
	}
	// Six per-core L2 instances of 1MB.
	if got := s.TotalCache(2, topology.Unified); got != 6<<20 {
		t.Errorf("TotalCache(L2)=%d want %d", got, 6<<20)
	}
	if bt := s.CoresByType(); bt[topology.Performance] != 2 || bt[topology.Efficiency] != 4 {
		t.Errorf("CoresByType=%v want P:2 E:4", bt)
	}
}

// TestSumBelowLevel checks that aggregates work at any level, so "total L3
// under this group" and "under the whole system" are both answerable.
func TestSumBelowLevel(t *testing.T) {
	s := samplecpu.Topology()
	g0 := &s.Packages[0].Groups[0]
	if got := g0.TotalCache(3, topology.Unified); got != 16<<20 {
		t.Errorf("group L3=%d want %d", got, 16<<20)
	}
	if got := g0.TotalCache(2, topology.Unified); got != 2<<20 {
		t.Errorf("group L2=%d want %d", got, 2<<20)
	}
	if got := g0.NumThreads(); got != 4 {
		t.Errorf("group threads=%d want 4", got)
	}
	if got := g0.NumCores(); got != 2 {
		t.Errorf("group cores=%d want 2", got)
	}
	if got := s.Packages[0].NumCores(); got != 6 {
		t.Errorf("package cores=%d want 6", got)
	}
}

// TestSampleCacheOnce verifies no cache instance is counted twice: summing the
// caches stored in the tree must equal the TotalCache aggregate.
func TestSampleCacheOnce(t *testing.T) {
	s := samplecpu.Topology()
	sum := 0
	for pi := range s.Packages {
		p := &s.Packages[pi]
		for _, c := range p.Caches {
			sum += c.Size
		}
		for gi := range p.Groups {
			g := &p.Groups[gi]
			for _, c := range g.Caches {
				sum += c.Size
			}
			for ci := range g.Cores {
				for _, c := range g.Cores[ci].Caches {
					sum += c.Size
				}
			}
		}
	}
	total := s.TotalCache(1, topology.Data) + s.TotalCache(1, topology.Instruction) +
		s.TotalCache(2, topology.Unified) + s.TotalCache(3, topology.Unified)
	if sum != total {
		t.Errorf("cache double-counted: tree sum %d != aggregate %d", sum, total)
	}
}

func TestStringNoPanic(t *testing.T) {
	if s := samplecpu.Topology().String(); len(s) == 0 {
		t.Error("empty String()")
	}
}

func TestEnumStrings(t *testing.T) {
	if topology.Unified.String() != "Unified" || topology.Data.String() != "Data" || topology.Instruction.String() != "Instruction" {
		t.Error("CacheType.String mismatch")
	}
	if topology.Performance.String() != "Performance" || topology.Efficiency.String() != "Efficiency" || topology.Unknown.String() != "Unknown" {
		t.Error("CoreType.String mismatch")
	}
}
