// WARNING: This file has automatically been generated on Sun, 12 May 2024 22:11:31 IST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package cgo

/*
#cgo LDFLAGS: -L${SRCDIR} -lezkl_ffi
#cgo linux LDFLAGS: -lcrypto -ldl -lm -lrt -lssl -ludev
#cgo darwin LDFLAGS: -F/Library/Frameworks -framework Security -framework CoreServices -framework IOKit -framework IOSurface -framework AppKit
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
func VerifyProof(proofLength uint64, proof []byte, vkLength uint64, vk []byte, settingsLength uint64, settings []byte, srsLength uint64, srs []byte) byte {
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
	__v := (byte)(__ret)
	return __v
}
