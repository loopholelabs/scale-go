// Code generated by scale-signature v0.3.20, DO NOT EDIT.
// output: signature

package signature

import (
	"github.com/loopholelabs/polyglot"
	"unsafe"
)

const hash = "3a592aa345d412faa2e6285ee048ca2ab5aa64b0caa2f9ca67b2c1e0792101e5"

var (
	writeBuffer = polyglot.NewBuffer()
	readBuffer  []byte
)

// Write serializes the signature into the global writeBuffer and returns the pointer to the buffer and its size
//
// Users should not use this method.
func Write(ctx *ModelWithAllFieldTypes) (uint32, uint32) {
	writeBuffer.Reset()
	ctx.Encode(writeBuffer)
	underlying := writeBuffer.Bytes()
	ptr := &underlying[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(writeBuffer.Len())
}

// Read deserializes signature from the global readBuffer
//
// Users should not use this method.
func Read(ctx *ModelWithAllFieldTypes) (*ModelWithAllFieldTypes, error) {
	return DecodeModelWithAllFieldTypes(ctx, readBuffer)
}

// Error serializes an error into the global writeBuffer and returns a pointer to the buffer and its size
//
// Users should not use this method.
func Error(err error) (uint32, uint32) {
	writeBuffer.Reset()
	polyglot.Encoder(writeBuffer).Error(err)
	underlying := writeBuffer.Bytes()
	ptr := &underlying[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(writeBuffer.Len())
}

// Resize resizes the global readBuffer to the given size and returns the pointer to the buffer
//
// Users should not use this method.
func Resize(size uint32) uint32 {
	if uint32(cap(readBuffer)) < size {
		readBuffer = append(make([]byte, 0, uint32(len(readBuffer))+size), readBuffer...)
	}
	readBuffer = readBuffer[:size]
	return uint32(uintptr(unsafe.Pointer(&readBuffer[0])))
}

// Hash returns the hash of the Scale Signature
//
// Users should not use this method.
func Hash() (uint32, uint32) {
	writeBuffer.Reset()
	polyglot.Encoder(writeBuffer).String(hash)
	underlying := writeBuffer.Bytes()
	ptr := &underlying[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(writeBuffer.Len())
}

// Next calls the next function in the Scale Function Chain
func Next(ctx *ModelWithAllFieldTypes) (*ModelWithAllFieldTypes, error) {
	next(Write(ctx))
	return Read(ctx)
}

//go:export next
//go:linkname next
func next(offset uint32, length uint32)
