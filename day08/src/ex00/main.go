package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty array")
	}
	if idx < 0 {
		return 0, errors.New("index out of range (negative)")
	}
	if idx >= len(arr) {
		return 0, errors.New("index out of range")
	}

	firstElemPtr := unsafe.Pointer(&arr[0])
	elemSize := unsafe.Sizeof(arr[0])
	targetPtr := uintptr(firstElemPtr) + uintptr(idx)*elemSize
	element := *(*int)(unsafe.Pointer(targetPtr))

	return element, nil
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	idx := 4
	element, err := getElement(arr, idx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index", idx, "is", element)
	}
}
