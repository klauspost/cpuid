// Copyright (c) 2015 Klaus Post, released under MIT License. See LICENSE file.

// +build !amd64,!386

package cpuid

func cpuid(op uint32) (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}

func cpuidex(op, op2 uint32) (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}
func xgetbv(index uint32) (eax, edx uint32) {
	return 0, 0
}
