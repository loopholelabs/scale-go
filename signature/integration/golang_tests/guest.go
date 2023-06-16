// Code generated by scale-signature-go v0.1.0, DO NOT EDIT.
// schema: MasterSchema:MasterSchemaTag
// output: golang_tests

package golang_tests

import (
	"github.com/loopholelabs/polyglot"
	"unsafe"
)

const identifier = "MasterSchema:MasterSchemaTag"

var (
	writeBuffer = polyglot.NewBuffer()
	readBuffer  []byte
)

// _Write serializes the signature into the global writeBuffer and returns the pointer to the buffer and its size
//
// Users should not use this method.
func _Write(ctx *ModelWithAllFieldTypes) (uint32, uint32) {
	writeBuffer.Reset()
	ctx.Encode(writeBuffer)
	underlying := writeBuffer.Bytes()
	ptr := &underlying[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(writeBuffer.Len())
}

// _Read deserializes signature from the global readBuffer
//
// Users should not use this method.
func _Read(ctx *ModelWithAllFieldTypes) (*ModelWithAllFieldTypes, error) {
	return DecodeModelWithAllFieldTypes(ctx, readBuffer)
}

// _Error serializes an error into the global writeBuffer and returns a pointer to the buffer and its size
//
// Users should not use this method.
func _Error(err error) (uint32, uint32) {
	writeBuffer.Reset()
	polyglot.Encoder(writeBuffer).Error(err)
	underlying := writeBuffer.Bytes()
	ptr := &underlying[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(writeBuffer.Len())
}

// _Resize resizes the global readBuffer to the given size and returns the pointer to the buffer
//
// Users should not use this method.
func _Resize(size uint32) uint32 {
	if uint32(cap(readBuffer)) < size {
		readBuffer = append(make([]byte, 0, uint32(len(readBuffer))+size), readBuffer...)
	}
	readBuffer = readBuffer[:size]
	return uint32(uintptr(unsafe.Pointer(&readBuffer[0])))
}

// _Identifier returns the identifier of the Scale Signature
//
// Users should not use this method.
func _Identifier() (uint32, uint32) {
	writeBuffer.Reset()
	polyglot.Encoder(writeBuffer).String(identifier)
	underlying := writeBuffer.Bytes()
	ptr := &underlying[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(writeBuffer.Len())
}

// Next calls the next function in the Scale Function Chain
func Next(ctx *ModelWithAllFieldTypes) (*ModelWithAllFieldTypes, error) {
	next(_Write(ctx))
	return _Read(ctx)
}

//export next
//go:linkname next
func next(offset uint32, length uint32)
