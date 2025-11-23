// https://go.dev/tour/generics/2
package generics

import (
	"fmt"
)

// 除了泛型函數，Go 也支持泛型類型（generic types）
// 類型可以用類型參數進行參數化，這對於實現泛型數據結構非常有用
//
// List 是一個單向鏈表的泛型實現
// - [T any] 是類型參數聲明
//   - T 是類型參數，可以是任何類型
//   - any 是類型約束，等同於 interface{}，表示 T 可以是任何類型
//
// - next *List[T] 是指向下一個節點的指針（遞歸使用泛型類型）
// - val T 是節點存儲的值，類型為 T
type List[T any] struct {
	next *List[T]
	val  T
}

// Push 在鏈表頭部添加一個新節點
// 返回新的頭節點
func (l *List[T]) Push(val T) *List[T] {
	return &List[T]{
		next: l,
		val:  val,
	}
}

// Append 在鏈表尾部添加一個新節點
// 如果鏈表為空，創建新節點作為頭節點
func (l *List[T]) Append(val T) *List[T] {
	if l == nil {
		return &List[T]{val: val}
	}

	// 找到鏈表尾部
	current := l
	for current.next != nil {
		current = current.next
	}

	// 在尾部添加新節點
	current.next = &List[T]{val: val}
	return l
}

// Length 返回鏈表的長度（節點數量）
func (l *List[T]) Length() int {
	if l == nil {
		return 0
	}

	count := 0
	current := l
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Print 打印鏈表的所有值
func (l *List[T]) Print() {
	if l == nil {
		fmt.Println("List is empty")
		return
	}

	current := l
	fmt.Print("List: ")
	for current != nil {
		fmt.Print(current.val)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println()
}

// Get 獲取指定索引位置的值
// 如果索引超出範圍，返回零值和 false
func (l *List[T]) Get(index int) (T, bool) {
	var zero T
	if l == nil {
		return zero, false
	}

	current := l
	for i := 0; i < index; i++ {
		if current.next == nil {
			return zero, false
		}
		current = current.next
	}
	return current.val, true
}

// Contains 檢查鏈表是否包含指定的值
// 需要 T 類型支持 == 操作符，所以使用 comparable 約束
func Contains[T comparable](l *List[T], val T) bool {
	current := l
	for current != nil {
		if current.val == val {
			return true
		}
		current = current.next
	}
	return false
}

func RunGenerics02() {
	// 創建一個整數類型的鏈表
	var intList *List[int]

	// 使用 Push 在頭部添加元素
	intList = intList.Push(3)
	intList = intList.Push(2)
	intList = intList.Push(1)

	fmt.Println("=== 整數鏈表 ===")
	intList.Print()                          // List: 1 -> 2 -> 3
	fmt.Printf("長度: %d\n", intList.Length()) // 長度: 3

	// 使用 Append 在尾部添加元素
	intList = intList.Append(4)
	intList = intList.Append(5)
	intList.Print() // List: 1 -> 2 -> 3 -> 4 -> 5

	// 獲取指定索引的值
	if val, ok := intList.Get(2); ok {
		fmt.Printf("索引 2 的值: %d\n", val) // 索引 2 的值: 3
	}

	// 檢查是否包含某個值
	fmt.Printf("包含 3: %v\n", Contains(intList, 3))   // 包含 3: true
	fmt.Printf("包含 10: %v\n", Contains(intList, 10)) // 包含 10: false

	fmt.Println()

	// 創建一個字符串類型的鏈表
	var strList *List[string]
	strList = strList.Push("world")
	strList = strList.Push("hello")

	fmt.Println("=== 字符串鏈表 ===")
	strList.Print()                          // List: hello -> world
	fmt.Printf("長度: %d\n", strList.Length()) // 長度: 2

	// 展示泛型類型的強大之處：同一個 List 類型可以處理不同類型的數據
}
