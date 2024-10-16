// WARNING: This file has automatically been generated on Fri, 24 May 2024 14:05:31 CEST.
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
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// copyPUint8TBytes copies the data from Go slice as *C.uint8_t.
func copyPUint8TBytes(slice *sliceHeader) (*C.uint8_t, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfUint8TValue) * slice.Len,
		Cap:  int(sizeOfUint8TValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.uint8_t)(mem0), allocs
}

// allocUint8TMemory allocates memory for type C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUint8TMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUint8TValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfUint8TValue = unsafe.Sizeof([1]C.uint8_t{})

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}
