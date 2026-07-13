// Copyright (c) 2026 Klaus Post, released under MIT License. See LICENSE file.

// Package samplecpu provides a synthetic, deterministic CPU topology for the
// cpuid module's own tests, examples and documentation. It is internal on
// purpose: it is a fixture, not part of the public API.
package samplecpu

import "github.com/klauspost/cpuid/v2/topology"

// Topology returns a synthetic topology: one package split into two groups (as
// a hybrid CPU's tiles or an AMD CPU's CCDs would be) — a performance group of
// two SMT cores and an efficiency group of four single-threaded cores, each
// core with private L1/L2 and each group with a shared L3. It is not read from
// any real CPU; use cpuid.ScanTopology or cpuid.CPU.Topology for that.
func Topology() *topology.System {
	core := func(id int, ct topology.CoreType, threads ...int) topology.Core {
		c := topology.Core{ID: id, Type: ct, Caches: []topology.Cache{
			{Level: 1, Type: topology.Data, Size: 32 << 10, LineSize: 64, SharedThreads: len(threads)},
			{Level: 1, Type: topology.Instruction, Size: 32 << 10, LineSize: 64, SharedThreads: len(threads)},
			{Level: 2, Type: topology.Unified, Size: 1 << 20, LineSize: 64, SharedThreads: len(threads)},
		}}
		for _, t := range threads {
			c.Threads = append(c.Threads, topology.Thread{APICID: uint32(t), ID: t})
		}
		return c
	}
	group := func(id, shared int, cores ...topology.Core) topology.Group {
		return topology.Group{ID: id, Cores: cores,
			Caches: []topology.Cache{{Level: 3, Type: topology.Unified, Size: 16 << 20, LineSize: 64, SharedThreads: shared}}}
	}
	return &topology.System{
		Vendor:  "GenuineIntel",
		Scanned: true,
		Online:  8,
		Packages: []topology.Package{{Groups: []topology.Group{
			group(0, 4, core(0, topology.Performance, 0, 1), core(1, topology.Performance, 2, 3)),
			group(1, 4,
				core(2, topology.Efficiency, 4), core(3, topology.Efficiency, 5),
				core(4, topology.Efficiency, 6), core(5, topology.Efficiency, 7)),
		}}},
	}
}
