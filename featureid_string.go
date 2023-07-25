// Code generated by "stringer -type=FeatureID,Vendor"; DO NOT EDIT.

package cpuid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ADX-1]
	_ = x[AESNI-2]
	_ = x[AMD3DNOW-3]
	_ = x[AMD3DNOWEXT-4]
	_ = x[AMXBF16-5]
	_ = x[AMXFP16-6]
	_ = x[AMXINT8-7]
	_ = x[AMXTILE-8]
	_ = x[APX_F-9]
	_ = x[AVX-10]
	_ = x[AVX10-11]
	_ = x[AVX10_128-12]
	_ = x[AVX10_256-13]
	_ = x[AVX10_512-14]
	_ = x[AVX2-15]
	_ = x[AVX512BF16-16]
	_ = x[AVX512BITALG-17]
	_ = x[AVX512BW-18]
	_ = x[AVX512CD-19]
	_ = x[AVX512DQ-20]
	_ = x[AVX512ER-21]
	_ = x[AVX512F-22]
	_ = x[AVX512FP16-23]
	_ = x[AVX512IFMA-24]
	_ = x[AVX512PF-25]
	_ = x[AVX512VBMI-26]
	_ = x[AVX512VBMI2-27]
	_ = x[AVX512VL-28]
	_ = x[AVX512VNNI-29]
	_ = x[AVX512VP2INTERSECT-30]
	_ = x[AVX512VPOPCNTDQ-31]
	_ = x[AVXIFMA-32]
	_ = x[AVXNECONVERT-33]
	_ = x[AVXSLOW-34]
	_ = x[AVXVNNI-35]
	_ = x[AVXVNNIINT8-36]
	_ = x[BHI_CTRL-37]
	_ = x[BMI1-38]
	_ = x[BMI2-39]
	_ = x[CETIBT-40]
	_ = x[CETSS-41]
	_ = x[CLDEMOTE-42]
	_ = x[CLMUL-43]
	_ = x[CLZERO-44]
	_ = x[CMOV-45]
	_ = x[CMPCCXADD-46]
	_ = x[CMPSB_SCADBS_SHORT-47]
	_ = x[CMPXCHG8-48]
	_ = x[CPBOOST-49]
	_ = x[CPPC-50]
	_ = x[CX16-51]
	_ = x[EFER_LMSLE_UNS-52]
	_ = x[ENQCMD-53]
	_ = x[ERMS-54]
	_ = x[F16C-55]
	_ = x[FLUSH_L1D-56]
	_ = x[FMA3-57]
	_ = x[FMA4-58]
	_ = x[FP128-59]
	_ = x[FP256-60]
	_ = x[FSRM-61]
	_ = x[FXSR-62]
	_ = x[FXSROPT-63]
	_ = x[GFNI-64]
	_ = x[HLE-65]
	_ = x[HRESET-66]
	_ = x[HTT-67]
	_ = x[HWA-68]
	_ = x[HYBRID_CPU-69]
	_ = x[HYPERVISOR-70]
	_ = x[IA32_ARCH_CAP-71]
	_ = x[IA32_CORE_CAP-72]
	_ = x[IBPB-73]
	_ = x[IBRS-74]
	_ = x[IBRS_PREFERRED-75]
	_ = x[IBRS_PROVIDES_SMP-76]
	_ = x[IBS-77]
	_ = x[IBSBRNTRGT-78]
	_ = x[IBSFETCHSAM-79]
	_ = x[IBSFFV-80]
	_ = x[IBSOPCNT-81]
	_ = x[IBSOPCNTEXT-82]
	_ = x[IBSOPSAM-83]
	_ = x[IBSRDWROPCNT-84]
	_ = x[IBSRIPINVALIDCHK-85]
	_ = x[IBS_FETCH_CTLX-86]
	_ = x[IBS_OPDATA4-87]
	_ = x[IBS_OPFUSE-88]
	_ = x[IBS_PREVENTHOST-89]
	_ = x[IBS_ZEN4-90]
	_ = x[IDPRED_CTRL-91]
	_ = x[INT_WBINVD-92]
	_ = x[INVLPGB-93]
	_ = x[KEYLOCKER-94]
	_ = x[KEYLOCKERW-95]
	_ = x[LAHF-96]
	_ = x[LAM-97]
	_ = x[LBRVIRT-98]
	_ = x[LZCNT-99]
	_ = x[MCAOVERFLOW-100]
	_ = x[MCDT_NO-101]
	_ = x[MCOMMIT-102]
	_ = x[MD_CLEAR-103]
	_ = x[MMX-104]
	_ = x[MMXEXT-105]
	_ = x[MOVBE-106]
	_ = x[MOVDIR64B-107]
	_ = x[MOVDIRI-108]
	_ = x[MOVSB_ZL-109]
	_ = x[MOVU-110]
	_ = x[MPX-111]
	_ = x[MSRIRC-112]
	_ = x[MSRLIST-113]
	_ = x[MSR_PAGEFLUSH-114]
	_ = x[NRIPS-115]
	_ = x[NX-116]
	_ = x[OSXSAVE-117]
	_ = x[PCONFIG-118]
	_ = x[POPCNT-119]
	_ = x[PPIN-120]
	_ = x[PREFETCHI-121]
	_ = x[PSFD-122]
	_ = x[RDPRU-123]
	_ = x[RDRAND-124]
	_ = x[RDSEED-125]
	_ = x[RDTSCP-126]
	_ = x[RRSBA_CTRL-127]
	_ = x[RTM-128]
	_ = x[RTM_ALWAYS_ABORT-129]
	_ = x[SERIALIZE-130]
	_ = x[SEV-131]
	_ = x[SEV_64BIT-132]
	_ = x[SEV_ALTERNATIVE-133]
	_ = x[SEV_DEBUGSWAP-134]
	_ = x[SEV_ES-135]
	_ = x[SEV_RESTRICTED-136]
	_ = x[SEV_SNP-137]
	_ = x[SGX-138]
	_ = x[SGXLC-139]
	_ = x[SHA-140]
	_ = x[SME-141]
	_ = x[SME_COHERENT-142]
	_ = x[SPEC_CTRL_SSBD-143]
	_ = x[SRBDS_CTRL-144]
	_ = x[SSE-145]
	_ = x[SSE2-146]
	_ = x[SSE3-147]
	_ = x[SSE4-148]
	_ = x[SSE42-149]
	_ = x[SSE4A-150]
	_ = x[SSSE3-151]
	_ = x[STIBP-152]
	_ = x[STIBP_ALWAYSON-153]
	_ = x[STOSB_SHORT-154]
	_ = x[SUCCOR-155]
	_ = x[SVM-156]
	_ = x[SVMDA-157]
	_ = x[SVMFBASID-158]
	_ = x[SVML-159]
	_ = x[SVMNP-160]
	_ = x[SVMPF-161]
	_ = x[SVMPFT-162]
	_ = x[SYSCALL-163]
	_ = x[SYSEE-164]
	_ = x[TBM-165]
	_ = x[TLB_FLUSH_NESTED-166]
	_ = x[TME-167]
	_ = x[TOPEXT-168]
	_ = x[TSCRATEMSR-169]
	_ = x[TSXLDTRK-170]
	_ = x[VAES-171]
	_ = x[VMCBCLEAN-172]
	_ = x[VMPL-173]
	_ = x[VMSA_REGPROT-174]
	_ = x[VMX-175]
	_ = x[VPCLMULQDQ-176]
	_ = x[VTE-177]
	_ = x[WAITPKG-178]
	_ = x[WBNOINVD-179]
	_ = x[WRMSRNS-180]
	_ = x[X87-181]
	_ = x[XGETBV1-182]
	_ = x[XOP-183]
	_ = x[XSAVE-184]
	_ = x[XSAVEC-185]
	_ = x[XSAVEOPT-186]
	_ = x[XSAVES-187]
	_ = x[AESARM-188]
	_ = x[ARMCPUID-189]
	_ = x[ASIMD-190]
	_ = x[ASIMDDP-191]
	_ = x[ASIMDHP-192]
	_ = x[ASIMDRDM-193]
	_ = x[ATOMICS-194]
	_ = x[CRC32-195]
	_ = x[DCPOP-196]
	_ = x[EVTSTRM-197]
	_ = x[FCMA-198]
	_ = x[FP-199]
	_ = x[FPHP-200]
	_ = x[GPA-201]
	_ = x[JSCVT-202]
	_ = x[LRCPC-203]
	_ = x[PMULL-204]
	_ = x[SHA1-205]
	_ = x[SHA2-206]
	_ = x[SHA3-207]
	_ = x[SHA512-208]
	_ = x[SM3-209]
	_ = x[SM4-210]
	_ = x[SVE-211]
	_ = x[lastID-212]
	_ = x[firstID-0]
}

const _FeatureID_name = "firstIDADXAESNIAMD3DNOWAMD3DNOWEXTAMXBF16AMXFP16AMXINT8AMXTILEAPX_FAVXAVX10AVX10_128AVX10_256AVX10_512AVX2AVX512BF16AVX512BITALGAVX512BWAVX512CDAVX512DQAVX512ERAVX512FAVX512FP16AVX512IFMAAVX512PFAVX512VBMIAVX512VBMI2AVX512VLAVX512VNNIAVX512VP2INTERSECTAVX512VPOPCNTDQAVXIFMAAVXNECONVERTAVXSLOWAVXVNNIAVXVNNIINT8BHI_CTRLBMI1BMI2CETIBTCETSSCLDEMOTECLMULCLZEROCMOVCMPCCXADDCMPSB_SCADBS_SHORTCMPXCHG8CPBOOSTCPPCCX16EFER_LMSLE_UNSENQCMDERMSF16CFLUSH_L1DFMA3FMA4FP128FP256FSRMFXSRFXSROPTGFNIHLEHRESETHTTHWAHYBRID_CPUHYPERVISORIA32_ARCH_CAPIA32_CORE_CAPIBPBIBRSIBRS_PREFERREDIBRS_PROVIDES_SMPIBSIBSBRNTRGTIBSFETCHSAMIBSFFVIBSOPCNTIBSOPCNTEXTIBSOPSAMIBSRDWROPCNTIBSRIPINVALIDCHKIBS_FETCH_CTLXIBS_OPDATA4IBS_OPFUSEIBS_PREVENTHOSTIBS_ZEN4IDPRED_CTRLINT_WBINVDINVLPGBKEYLOCKERKEYLOCKERWLAHFLAMLBRVIRTLZCNTMCAOVERFLOWMCDT_NOMCOMMITMD_CLEARMMXMMXEXTMOVBEMOVDIR64BMOVDIRIMOVSB_ZLMOVUMPXMSRIRCMSRLISTMSR_PAGEFLUSHNRIPSNXOSXSAVEPCONFIGPOPCNTPPINPREFETCHIPSFDRDPRURDRANDRDSEEDRDTSCPRRSBA_CTRLRTMRTM_ALWAYS_ABORTSERIALIZESEVSEV_64BITSEV_ALTERNATIVESEV_DEBUGSWAPSEV_ESSEV_RESTRICTEDSEV_SNPSGXSGXLCSHASMESME_COHERENTSPEC_CTRL_SSBDSRBDS_CTRLSSESSE2SSE3SSE4SSE42SSE4ASSSE3STIBPSTIBP_ALWAYSONSTOSB_SHORTSUCCORSVMSVMDASVMFBASIDSVMLSVMNPSVMPFSVMPFTSYSCALLSYSEETBMTLB_FLUSH_NESTEDTMETOPEXTTSCRATEMSRTSXLDTRKVAESVMCBCLEANVMPLVMSA_REGPROTVMXVPCLMULQDQVTEWAITPKGWBNOINVDWRMSRNSX87XGETBV1XOPXSAVEXSAVECXSAVEOPTXSAVESAESARMARMCPUIDASIMDASIMDDPASIMDHPASIMDRDMATOMICSCRC32DCPOPEVTSTRMFCMAFPFPHPGPAJSCVTLRCPCPMULLSHA1SHA2SHA3SHA512SM3SM4SVElastID"

var _FeatureID_index = [...]uint16{0, 7, 10, 15, 23, 34, 41, 48, 55, 62, 67, 70, 75, 84, 93, 102, 106, 116, 128, 136, 144, 152, 160, 167, 177, 187, 195, 205, 216, 224, 234, 252, 267, 274, 286, 293, 300, 311, 319, 323, 327, 333, 338, 346, 351, 357, 361, 370, 388, 396, 403, 407, 411, 425, 431, 435, 439, 448, 452, 456, 461, 466, 470, 474, 481, 485, 488, 494, 497, 500, 510, 520, 533, 546, 550, 554, 568, 585, 588, 598, 609, 615, 623, 634, 642, 654, 670, 684, 695, 705, 720, 728, 739, 749, 756, 765, 775, 779, 782, 789, 794, 805, 812, 819, 827, 830, 836, 841, 850, 857, 865, 869, 872, 878, 885, 898, 903, 905, 912, 919, 925, 929, 938, 942, 947, 953, 959, 965, 975, 978, 994, 1003, 1006, 1015, 1030, 1043, 1049, 1063, 1070, 1073, 1078, 1081, 1084, 1096, 1110, 1120, 1123, 1127, 1131, 1135, 1140, 1145, 1150, 1155, 1169, 1180, 1186, 1189, 1194, 1203, 1207, 1212, 1217, 1223, 1230, 1235, 1238, 1254, 1257, 1263, 1273, 1281, 1285, 1294, 1298, 1310, 1313, 1323, 1326, 1333, 1341, 1348, 1351, 1358, 1361, 1366, 1372, 1380, 1386, 1392, 1400, 1405, 1412, 1419, 1427, 1434, 1439, 1444, 1451, 1455, 1457, 1461, 1464, 1469, 1474, 1479, 1483, 1487, 1491, 1497, 1500, 1503, 1506, 1512}

func (i FeatureID) String() string {
	if i < 0 || i >= FeatureID(len(_FeatureID_index)-1) {
		return "FeatureID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureID_name[_FeatureID_index[i]:_FeatureID_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VendorUnknown-0]
	_ = x[Intel-1]
	_ = x[AMD-2]
	_ = x[VIA-3]
	_ = x[Transmeta-4]
	_ = x[NSC-5]
	_ = x[KVM-6]
	_ = x[MSVM-7]
	_ = x[VMware-8]
	_ = x[XenHVM-9]
	_ = x[Bhyve-10]
	_ = x[Hygon-11]
	_ = x[SiS-12]
	_ = x[RDC-13]
	_ = x[Ampere-14]
	_ = x[ARM-15]
	_ = x[Broadcom-16]
	_ = x[Cavium-17]
	_ = x[DEC-18]
	_ = x[Fujitsu-19]
	_ = x[Infineon-20]
	_ = x[Motorola-21]
	_ = x[NVIDIA-22]
	_ = x[AMCC-23]
	_ = x[Qualcomm-24]
	_ = x[Marvell-25]
	_ = x[lastVendor-26]
}

const _Vendor_name = "VendorUnknownIntelAMDVIATransmetaNSCKVMMSVMVMwareXenHVMBhyveHygonSiSRDCAmpereARMBroadcomCaviumDECFujitsuInfineonMotorolaNVIDIAAMCCQualcommMarvelllastVendor"

var _Vendor_index = [...]uint8{0, 13, 18, 21, 24, 33, 36, 39, 43, 49, 55, 60, 65, 68, 71, 77, 80, 88, 94, 97, 104, 112, 120, 126, 130, 138, 145, 155}

func (i Vendor) String() string {
	if i < 0 || i >= Vendor(len(_Vendor_index)-1) {
		return "Vendor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Vendor_name[_Vendor_index[i]:_Vendor_index[i+1]]
}
