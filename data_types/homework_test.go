package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian[T interface{ ~uint32 | ~uint16 | ~uint64 }](number T) T {
	size := unsafe.Sizeof(number)

	buf := make([]byte, size)

	for i := uintptr(0); i < size; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&number)) + i))
	}

	for i := 0; i < int(size)/2; i++ {
		j := int(size) - 1 - i
		buf[i], buf[j] = buf[j], buf[i]
	}

	return *(*T)(unsafe.Pointer(&buf[0]))
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
