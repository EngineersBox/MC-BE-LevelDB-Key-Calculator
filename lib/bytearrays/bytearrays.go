package lib

import (
	"reflect"
	"unsafe"
)

func IntToByteArray(num int32) []byte {
	// Get the size of the int in bytes
	size := int(unsafe.Sizeof(num))
	// Pre-fill a blank slice with null bytes to fit the int size
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		// Convert the current byte in the int to a byte
		// Formula: byte = pointer_of(ptr_to_int + curent_byte_offset)
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		// Append byte to slice
		arr[i] = byt
	}
	return arr
}

func ByteArrayToInt(arr []byte) int32 {
	// Pre-create an int32 to place bytes into
	val := int32(0)
	// Get the size of the byte array
	size := len(arr)
	for i := 0; i < size; i++ {
		// Place the current byte within the int at the offset position
		// Formula: pointer_of(ptr_to_int + curent_byte_offset) = byte
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}

func ReverseAny(s interface{}) {
	// Reflectively get the value of the slice size
	n := reflect.ValueOf(s).Len()
	// Create a swapper that will perform actions on the slice
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		// Swap each far-left value with the far-right value floor(len / 2) times
		// Example:
		// [1, 2, 3, 4, 5] <-- intitial
		// [5, 2, 3, 4, 5] <-- 1st pass
		// [5, 4, 3, 2, 5] <-- 2nd pass [DONE]
		swap(i, j)
	}
}
