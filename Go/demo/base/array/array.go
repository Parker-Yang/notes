package array

import (
	"sync"
)

// 在golang中可变长数组由内部的slice实现
// slice的内部结构
/*type slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int // 数组的真实
	cap   int // 数组的容量
}*/
// 每次可以初始化一个固定容量的切片，切片内部维护一个固定大小的数组。当 append 新元素时，固定大小的数组不够时会自动扩容
/*// 创建一个长度为零，容量为二的可变长数组
arr := make([]int, 0, 2)
// 同过append方法可实现对数组的添加赋值
arr = append(arr, 1, 2)
t.Log(arr)
t.Log("-----------")
// 浅拷贝
ar := arr
ar[1] = 3
t.Log(arr)
t.Log(ar)
t.Log("-----------")
// 深拷贝
arrr := make([]int,len(arr))
copy(arrr, arr)
arrr[0] = 4
t.Log(arr)
t.Log(arrr)*/
type Array struct {
	array []int      // 数组
	len   int        // 数组的真实
	cap   int        // 数组的容量
	lock  sync.Mutex //为并发安全实现的锁
}

func Init(len, cap int) *Array {
	arr := new(Array)
	if len > cap {
		panic("len large than cap")
	}
	// 把切片当数组用
	array := make([]int, cap, cap)
	// 元数据
	arr.array = array
	arr.cap = cap
	arr.len = 0
	return arr
}

func (a *Array) Append(element int) {
	// 并发锁
	a.lock.Lock()
	// 使用完及时释放
	defer a.lock.Unlock()
	// 先判断是否满容量
	if a.len == a.cap {
		//slice的扩容策略为：
		//	若新入元素大小通过倍数增长能够hold住，那就根据旧容量是否超过1024决定是倍数增长还是1.25倍逐步增长；
		//	若新入元素大小超过了原有的容量，则新容量取两者相加计算出来的最小cap值。
		// 这里简单实现按照2倍扩容策略
		newCap := 2 * a.cap

		// 如果源cap为0
		if a.cap == 0 {
			newCap = 1
		}
		// 满容量cap为数组
		newArray := make([]int, newCap, newCap)
		// 数据迁移
		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = element
	a.len++
}

//批量添加
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// 获取指定下表的值
func (a *Array) Get(index int) int {
	if a.len == 0 || index > a.len {
		panic("out of array index")
	}
	return a.array[index]
}

// 获取长度
func (a *Array) Len() int {
	return a.len
}

// 获取容量
func (a *Array) Cap() int {
	return a.cap

}
