package main

import (
	"crypto/sha256"
	"reflect"
	"unsafe"
)

var pow256 = bigPow(2, 256)

type SingleHead struct {
	coinbase   uint8
	shardNum   uint32
	difficulty uint64
	merkleRoot [32]byte
	nonce      uint64
}

type MultiHead struct {
	headRoot   [32]byte
	shardCount uint32
	headPath   [32]byte
	nonce      uint64
}

func (singleHead *SingleHead) hash() [32]byte {

	var x reflect.SliceHeader
	len := unsafe.Sizeof(*singleHead)
	x.Len = int(len)
	x.Cap = x.Len
	x.Data = uintptr(unsafe.Pointer(singleHead))
	bytes := *(*[]byte)(unsafe.Pointer(&x))
	//glog.Infof("bytes=%v", bytes)

	h := sha256.New()

	h.Write(bytes)
	hash := h.Sum(nil)

	var hashArr [32]byte
	for i := 0; i < 32; i++ {
		hashArr[i] = hash[i]
	}

	return hashArr
}

func (multiHead *MultiHead) hash() [32]byte {

	var x reflect.SliceHeader
	len := unsafe.Sizeof(*multiHead)
	x.Len = int(len)
	x.Cap = x.Len
	x.Data = uintptr(unsafe.Pointer(multiHead))
	bytes := *(*[]byte)(unsafe.Pointer(&x))
	//glog.Infof("bytes=%v", bytes)

	h := sha256.New()

	h.Write(bytes)
	hash := h.Sum(nil)

	var hashArr [32]byte
	for i := 0; i < 32; i++ {
		hashArr[i] = hash[i]
	}

	return hashArr
}