package bitmanip

import (
//"fmt"
)

var one byte = []byte("1")[0]
var zero byte = []byte("0")[0]

// Convert a string of 0 and 1s to an unsigned integer.
// Useful for controlling the base2 number easily.
func Base2(s string) uint64 {
	if len(s) > 64 {
		panic("String is too long")
	}

	var res uint64 = 0
	var index uint64 = 0
	for i := len(s) - 1; i >= 0; i-- {
		index++
		switch s[i] {
		case one:
			res |= (1 << uint64(index-1))
		case zero:
			continue
		default:
			panic("Should be 0 or 1")
		}
	}
	return res
}

//################################
// Parity: finding the parity of a bit word.
// e.g. count the number of bit set to 1 in the word.
// Complexity is O(sizeof(uint64)) -- note I could use unsafe.Sizeof
// aka my dumb version
func Parity(in uint64) uint64 {
	var res uint64 = 0
	for i := uint8(0); i < 64; i++ {
		res ^= (in & (1 << i)) >> i // = res XOR
	}
	return res
}

// For this case, we keep one and compare it with the lowest bit.
// we keep shifting "in" right to chek the next bit until the whole word is zero.
// O(sizeof(uint64)), but stops as soon as the last most-left 1 is counted-for.
func FasterParity(in uint64) uint64 {
	var res uint64 = 0
	for in != 0 {
		res ^= in & 1 // = res XOR in AND 00...001
		in = in >> 1
	}
	return res
}

// For this case, we use the trick: in AND (in - 1) which is equal to in with the lowest 1 flipped to 0
// We repeat the operation until we've exausted all the ones
// O(numberOfOnes(in)) -- so the complexity depends on the number of ones, regardless of the number of zeros
// or their position.
func EvenFasterParity(in uint64) uint64 {
	var res uint64 = 0
	for in != 0 {
		res ^= 1 // = res XOR 1
		in &= in - 1 // = in AND in -1 (technique to remove most right set bit)
	}
	return res
}

// In this case, we leverage a cache with the 16bit words parity already saved.
// Then we just need to XOR each part of the 64 word.
func CachedParity(in uint64, parityCache *[65536]uint8) uint64 {
	return uint64(
		parityCache[(in >> 48)] ^ //parity of left most 16
		parityCache[(in >> 32) & 0xFFFF ] ^ //parity of 2nd block of 16
		parityCache[(in >> 16) & 0xFFFF ] ^ //parity of 3rd block of 16
		parityCache[in & 0xFFFF] ) //parity of last block of 16
}

func ProduceCache() *[65536]uint8 {
	var cache [65536]uint8
	for i := 0; i < 65536; i++ {
		cache[i] = uint8(EvenFasterParity(uint64(i)))
	}
	return &cache
}


// In this case, we leverage the XOR ability of the language and divide and conquer
// O(log(sizeof(uint64)))
func XORParity(in uint64) uint64 {
	in ^= in >> 32
	in ^= in >> 16
	in ^= in >> 8
	in ^= in >> 4
	in ^= in >> 2
	in ^= in >> 1
	return in & 0x1
}


//################################
