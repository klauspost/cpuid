// Copyright (c) 2021 Klaus Post, released under MIT License. See LICENSE file.

// Package cpuid provides information about the CPU running the current program.
//
// CPU features are detected on startup, and kept for fast access through the life of the application.
// Currently x86 / x64 (AMD64), arm64, and riscv64 are supported.
//
// You can access the CPU information by accessing the shared CPU variable of the cpuid library.
//
// Package home: https://github.com/klauspost/cpuid
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/klauspost/cpuid/v2"
)

var js = flag.Bool("json", false, "Output as JSON")
var level = flag.Int("check-level", 0, "Check microarchitecture level. Exit code will be 0 if supported")

func main() {
	flag.Parse()
	if level != nil && *level > 0 {
		if *level < 1 || *level > 4 {
			log.Fatalln("Supply CPU level 1-4 to test as argument")
		}
		log.Println(cpuid.CPU.BrandName)
		if cpuid.CPU.X64Level() < *level {
			// Does os.Exit(1)
			log.Fatalf("Microarchitecture level %d not supported. Max level is %d.", *level, cpuid.CPU.X64Level())
		}
		log.Printf("Microarchitecture level %d is supported. Max level is %d.", *level, cpuid.CPU.X64Level())
		os.Exit(0)
	}
	if *js {
		info := struct {
			cpuid.CPUInfo
			Features  []string
			X64Level  int
			RVProfile int
			GOARM64   string `json:"GOARM64,omitempty"`
		}{
			CPUInfo:   cpuid.CPU,
			Features:  cpuid.CPU.FeatureSet(),
			X64Level:  cpuid.CPU.X64Level(),
			RVProfile: cpuid.CPU.RVProfile(),
			GOARM64:   cpuid.CPU.GOARM64(),
		}
		b, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		os.Exit(0)
	}

	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("Vendor String:", cpuid.CPU.VendorString)
	fmt.Println("Vendor ID:", cpuid.CPU.VendorID)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("Threads Per Core:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("Logical Cores:", cpuid.CPU.LogicalCores)
	fmt.Println("CPU Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model, "Stepping:", cpuid.CPU.Stepping)
	fmt.Println("Features:", strings.Join(cpuid.CPU.FeatureSet(), ","))
	if x := cpuid.CPU.X64Level(); x > 0 {
		fmt.Println("Microarchitecture level:", x)
	}
	if cpuid.CPU.AVX10Level > 0 {
		fmt.Println("AVX10 level:", cpuid.CPU.AVX10Level)
	}
	if rvp := cpuid.CPU.RVProfile(); rvp > 0 {
		fmt.Printf("RISC-V Profile: RVA%d\n", rvp)
	}
	if v := cpuid.CPU.GOARM64(); v != "" {
		fmt.Println("GOARM64:", v)
	}
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1I, "bytes")
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")
	if cpuid.CPU.Hz > 0 {
		fmt.Println("Frequency:", cpuid.CPU.Hz, "Hz")
	}
	if cpuid.CPU.BoostFreq > 0 {
		fmt.Println("Boost Frequency:", cpuid.CPU.BoostFreq, "Hz")
	}
	if cpuid.CPU.SGX.Available {
		fmt.Printf("SGX: %+v\n", cpuid.CPU.SGX)
	}
	if cpuid.CPU.AMDMemEncryption.Available {
		fmt.Printf("AMD Memory Encryption: %+v\n", cpuid.CPU.AMDMemEncryption)
	}
	if cpuid.CPU.PMU.VersionID != 0 {
		fmt.Println("PMU version:", cpuid.CPU.PMU.VersionID,
			"Fixed Counters:", cpuid.CPU.PMU.NumFixedPMC,
			"General Purpose Counters:", cpuid.CPU.PMU.NumGPCounters)
	}
}
