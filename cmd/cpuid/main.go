// Copyright (c) 2021 Klaus Post, released under MIT License. See LICENSE file.

// Package cpuid provides information about the CPU running the current program.
//
// CPU features are detected on startup, and kept for fast access through the life of the application.
// Currently x86 / x64 (AMD64) as well as arm64 is supported.
//
// You can access the CPU information by accessing the shared CPU variable of the cpuid library.
//
// Package home: https://github.com/klauspost/cpuid
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/klauspost/cpuid/v2"
)

var js = flag.Bool("json", false, "Output as JSON")

func main() {
	flag.Parse()
	if *js {
		b, err := json.MarshalIndent(cpuid.CPU, "", "  ")
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
	fmt.Println("CPU Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model)
	fmt.Println("Features:", fmt.Sprintf(strings.Join(cpuid.CPU.FeatureSet(), ",")))
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
}
