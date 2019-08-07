package main

import (
	"fmt"
	"github.com/klauspost/cpuid"
)

func main() {
	fmt.Println("IBS Enabled:", cpuid.CPU.IBS())
	fmt.Println("Extended IBS:", cpuid.CPU.ExtendedIBS())
}
