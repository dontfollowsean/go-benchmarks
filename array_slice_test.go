package main

import (
	"testing"
)

var globalSlice = make([]byte, 1024) // Global slice
var globalArray [1024]byte           // Global array

func BenchmarkSliceGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, value := range globalSlice {
			globalSlice[j]++
			globalSlice[j] = globalSlice[j] + value + 10
			globalSlice[j] += value
		}
	}
}

func BenchmarkArrayGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, value := range globalArray {
			globalArray[j]++
			globalArray[j] = globalArray[j] + value + 10
			globalArray[j] += value
		}
	}
}

func BenchmarkSliceLocal(b *testing.B) {
	var slice = make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		for j, value := range slice {
			slice[j]++
			slice[j] = slice[j] + value + 10
			slice[j] += value
		}
	}
}

func BenchmarkArrayLocal(b *testing.B) {
	var array [1024]byte
	for i := 0; i < b.N; i++ {
		for j, value := range array {
			array[j]++
			array[j] = array[j] + value + 10
			array[j] += value
		}
	}
}
