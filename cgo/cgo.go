// WARNING: This file has automatically been generated on Fri, 24 May 2024 14:49:00 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package cgo

/*
#cgo LDFLAGS: -L${SRCDIR} -lezkl_ffi
#cgo linux LDFLAGS: -lcrypto -ldl -lm -lrt -lssl -ludev
#cgo darwin LDFLAGS: -F/Library/Frameworks -framework Security -framework CoreServices -framework IOKit -framework IOSurface -framework AppKit -framework SystemConfiguration
#include "ezkl-ffi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// VerifyProof function as declared in cgo/ezkl-ffi.h:17
func VerifyProof(proofLength uint64, proof []byte, vkLength uint64, vk []byte, settingsLength uint64, settings []byte, srsLength uint64, srs []byte) bool {
	cproofLength, cproofLengthAllocMap := (C.size_t)(proofLength), cgoAllocsUnknown
	cproof, cproofAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&proof)))
	cvkLength, cvkLengthAllocMap := (C.size_t)(vkLength), cgoAllocsUnknown
	cvk, cvkAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&vk)))
	csettingsLength, csettingsLengthAllocMap := (C.size_t)(settingsLength), cgoAllocsUnknown
	csettings, csettingsAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&settings)))
	csrsLength, csrsLengthAllocMap := (C.size_t)(srsLength), cgoAllocsUnknown
	csrs, csrsAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&srs)))
	__ret := C.verify_proof(cproofLength, cproof, cvkLength, cvk, csettingsLength, csettings, csrsLength, csrs)
	runtime.KeepAlive(csrsAllocMap)
	runtime.KeepAlive(csrsLengthAllocMap)
	runtime.KeepAlive(csettingsAllocMap)
	runtime.KeepAlive(csettingsLengthAllocMap)
	runtime.KeepAlive(cvkAllocMap)
	runtime.KeepAlive(cvkLengthAllocMap)
	runtime.KeepAlive(cproofAllocMap)
	runtime.KeepAlive(cproofLengthAllocMap)
	__v := (bool)(__ret)
	return __v
}

// Prove function as declared in cgo/ezkl-ffi.h:26
func Prove(witnessLength uint64, witness []byte, pkLength uint64, pk []byte, cpmpiledCircuitLength uint64, compiledCircuit []byte, srsLength uint64, srs []byte) string {
	cwitnessLength, cwitnessLengthAllocMap := (C.size_t)(witnessLength), cgoAllocsUnknown
	cwitness, cwitnessAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&witness)))
	cpkLength, cpkLengthAllocMap := (C.size_t)(pkLength), cgoAllocsUnknown
	cpk, cpkAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&pk)))
	ccpmpiledCircuitLength, ccpmpiledCircuitLengthAllocMap := (C.size_t)(cpmpiledCircuitLength), cgoAllocsUnknown
	ccompiledCircuit, ccompiledCircuitAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&compiledCircuit)))
	csrsLength, csrsLengthAllocMap := (C.size_t)(srsLength), cgoAllocsUnknown
	csrs, csrsAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&srs)))
	__ret := C.prove(cwitnessLength, cwitness, cpkLength, cpk, ccpmpiledCircuitLength, ccompiledCircuit, csrsLength, csrs)
	runtime.KeepAlive(csrsAllocMap)
	runtime.KeepAlive(csrsLengthAllocMap)
	runtime.KeepAlive(ccompiledCircuitAllocMap)
	runtime.KeepAlive(ccpmpiledCircuitLengthAllocMap)
	runtime.KeepAlive(cpkAllocMap)
	runtime.KeepAlive(cpkLengthAllocMap)
	runtime.KeepAlive(cwitnessAllocMap)
	runtime.KeepAlive(cwitnessLengthAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// GenVk function as declared in cgo/ezkl-ffi.h:35
func GenVk(compiledCircuitLength uint64, compiledCircuit []byte, paramsSerializedLength uint64, paramsSerialized []byte, compressSelectors bool) {
	ccompiledCircuitLength, ccompiledCircuitLengthAllocMap := (C.size_t)(compiledCircuitLength), cgoAllocsUnknown
	ccompiledCircuit, ccompiledCircuitAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&compiledCircuit)))
	cparamsSerializedLength, cparamsSerializedLengthAllocMap := (C.size_t)(paramsSerializedLength), cgoAllocsUnknown
	cparamsSerialized, cparamsSerializedAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&paramsSerialized)))
	ccompressSelectors, ccompressSelectorsAllocMap := (C._Bool)(compressSelectors), cgoAllocsUnknown
	C.gen_vk(ccompiledCircuitLength, ccompiledCircuit, cparamsSerializedLength, cparamsSerialized, ccompressSelectors)
	runtime.KeepAlive(ccompressSelectorsAllocMap)
	runtime.KeepAlive(cparamsSerializedAllocMap)
	runtime.KeepAlive(cparamsSerializedLengthAllocMap)
	runtime.KeepAlive(ccompiledCircuitAllocMap)
	runtime.KeepAlive(ccompiledCircuitLengthAllocMap)
}

// GenPk function as declared in cgo/ezkl-ffi.h:41
func GenPk(vkLength uint64, vk []byte, circuitLength uint64, circuit []byte, paramsSerializedLength uint64, paramsSerialized []byte) {
	cvkLength, cvkLengthAllocMap := (C.size_t)(vkLength), cgoAllocsUnknown
	cvk, cvkAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&vk)))
	ccircuitLength, ccircuitLengthAllocMap := (C.size_t)(circuitLength), cgoAllocsUnknown
	ccircuit, ccircuitAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&circuit)))
	cparamsSerializedLength, cparamsSerializedLengthAllocMap := (C.size_t)(paramsSerializedLength), cgoAllocsUnknown
	cparamsSerialized, cparamsSerializedAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&paramsSerialized)))
	C.gen_pk(cvkLength, cvk, ccircuitLength, ccircuit, cparamsSerializedLength, cparamsSerialized)
	runtime.KeepAlive(cparamsSerializedAllocMap)
	runtime.KeepAlive(cparamsSerializedLengthAllocMap)
	runtime.KeepAlive(ccircuitAllocMap)
	runtime.KeepAlive(ccircuitLengthAllocMap)
	runtime.KeepAlive(cvkAllocMap)
	runtime.KeepAlive(cvkLengthAllocMap)
}
