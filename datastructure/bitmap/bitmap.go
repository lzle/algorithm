package main

import "fmt"

type BitMap struct {
	nbits int
	bytes []int
}

func (bm *BitMap) Set(k int) {
	if k > bm.nbits {
		return
	}
	byteIndex := k / 8
	bitIndex := k % 8
	bm.bytes[byteIndex] =  bm.bytes[byteIndex] | (1 << bitIndex)
}

func (bm *BitMap) Get(k int) bool {
	if k > bm.nbits {
		return false
	}
	byteIndex := k / 8
	bitIndex := k % 8
	return (bm.bytes[byteIndex] & (1 << bitIndex)) != 0
}

func newBitMap(n int) *BitMap {
	bm := new(BitMap)
	bm.nbits = n
	bm.bytes = make([]int, n/8+1)
	return bm
}

// 位图
func main() {
	bm := newBitMap(100)
	fmt.Println(bm.bytes)
	bm.Set(10)
	bm.Set(33)
	bm.Set(60)
	fmt.Println(bm.bytes)
	fmt.Println(bm.Get(20))
}
