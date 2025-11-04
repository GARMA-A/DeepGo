package main

type dynamicArray struct {
	data []int
	size int
	cap  int
}

func NewDynamicArray() *dynamicArray {
	return &dynamicArray{
		data: make([]int, 2),
		size: 0,
		cap:  2,
	}
}

func (da *dynamicArray) Append(value int) {
	if da.size == da.cap {
		newCap := da.cap * 2
		newData := make([]int, newCap)
		copy(newData, da.data)
		da.data = newData
		da.cap = newCap
	}
	da.data[da.size] = value
	da.size++
}

func (da *dynamicArray) pop() {
	if da.size == 0 {
		return
	}
	da.size--
}
