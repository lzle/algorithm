package array

import (
	"fmt"
	"github.com/pkg/errors"
)

/*
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 * Author: Lian
 */

type Array struct {
	data     []int
	length   uint
	capacity uint
}

// 初始化数组
func NewArray(cap uint) *Array {
	if cap <= 0 {
		return nil
	}
	array := &Array{
		data:     make([]int, cap, cap),
		length:   0,
		capacity: cap,
	}
	return array
}

// 长度
func (array *Array) Len() uint{
	return array.length
}

// 容量
func (array *Array) Cap() uint{
	return array.capacity
}

// 越界
func (array *Array) isOutOfRange(index uint) bool{
	if index >= array.Cap() {
		return true
	}
	return false
}

// 插入
func (array *Array) Insert(index uint, value int) error {
	// 是否还有空间
	if array.Len() == array.Cap() {
		return errors.New("full array")
	}
	// 是否越界
	if array.isOutOfRange(index) {
		return errors.New("index out of range")
	}

	for i := array.Len(); i > index; i-- {
		array.data[i] = array.data[i-1]
	}
	array.data[index] = value
	array.length++
	return nil
}

// 删除
func (array *Array) Delete(index uint) error {
	// 是否越界
	if array.isOutOfRange(index) {
		return errors.New("index out of range")
	}

	for i := index; i < array.Len()-1; i++ {
		array.data[i] = array.data[i+1]
	}
	array.length--
	return nil
}

// 追加
func (array *Array) Append(value int) error {
	// 是否还有空间
	if array.Len() == array.Cap() {
		return errors.New("full array")
	}
	array.data[array.Len()] = value
	array.length++
	return nil
}

// 索引查找
func (array *Array) Find(index uint) (int,error){
	// 是否越界
	if array.isOutOfRange(index) {
		return 0,errors.New("index out of range")
	}
	return array.data[index],nil
}

// 打印数列
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
