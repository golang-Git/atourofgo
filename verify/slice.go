package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ns := []int{10, 20, 30}
	ptr := unsafe.Pointer(&ns)
	s := *(*reflect.SliceHeader)(ptr)
	fmt.Printf("%#v\n", s)
}


package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func at(s reflect.SliceHeader, i int) unsafe.Pointer {
	// 先頭ポインタ + インデックス * int型のサイズ
	return unsafe.Pointer(s.Data + uintptr(i)*unsafe.Sizeof(int(0)))
}

func main() {
	a := [...]int{10, 20, 30}
	s := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&a[0])), Len: 2, Cap: len(a)}
	*(*int)(at(s, 2)) = 100 // unsafe.Pointerを*intに変換して代入している
	fmt.Println(a)
}


package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func at(s reflect.SliceHeader, i int) unsafe.Pointer {
	// 先頭ポインタ + インデックス * int型のサイズ
	return unsafe.Pointer(s.Data + uintptr(i)*unsafe.Sizeof(int(0)))
}

func myAppend(s reflect.SliceHeader, vs ...int) reflect.SliceHeader {
	// 新しい要素の追加
	for i := 0; i < len(vs); i++ {
		*((*int)(at(s, s.Len+i))) = vs[i]
	}
	return reflect.SliceHeader{Data: s.Data, Len: s.Len + len(vs), Cap: s.Cap}
}

func main() {
	a := [...]int{10, 20, 30}
	// s := a[0:2] -> [10 20]
	s := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&a[0])), Len: 2, Cap: len(a)}
	s = myAppend(s, 400)

	var ns []int
	*(*reflect.SliceHeader)(unsafe.Pointer(&ns)) = s
	fmt.Println(ns)
}


package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func at(s reflect.SliceHeader, i int) unsafe.Pointer {
	// 先頭ポインタ + インデックス * int型のサイズ
	return unsafe.Pointer(s.Data + uintptr(i)*unsafe.Sizeof(int(0)))
}

func myAppend(s reflect.SliceHeader, vs ...int) reflect.SliceHeader {
	// 容量が足りない場合
	if s.Len+len(vs) > s.Cap {
		fmt.Printf("len(vs)=%v\n", len(vs))
		s = growslice(s, s.Len+len(vs))
	}

	// 新しい要素の追加
	for i := 0; i < len(vs); i++ {
		*((*int)(at(s, s.Len+i))) = vs[i]
	}
	return reflect.SliceHeader{Data: s.Data, Len: s.Len + len(vs), Cap: s.Cap}
}

func growslice(old reflect.SliceHeader, cap int) reflect.SliceHeader {
	newcap := cap
	doublecap := old.Cap + old.Cap
	fmt.Printf("s.Len+len(vs)cap=%v s.Len=%v\n", cap, old.Cap)
	if cap < doublecap {
		newcap = doublecap
	}

	s := make([]int, old.Len, newcap)
	newslice := *(*reflect.SliceHeader)(unsafe.Pointer(&s))

	// 古いスライスから要素のコピー
	for i := 0; i < old.Len; i++ {
		*((*int)(at(newslice, i))) = *((*int)(at(old, i)))
	}

	return newslice
}

func main() {
	a := [...]int{10, 20, 30}
	// s := a[0:2] -> [10 20]
	s := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&a[0])), Len: 2, Cap: len(a)}
	s = myAppend(s, 400, 500) // 溢れる

	var ns []int
	*(*reflect.SliceHeader)(unsafe.Pointer(&ns)) = s
	fmt.Println(ns)
	fmt.Println(s.Cap)
}
*/

func main() {
	v := struct {
		a [5]int
		b [2]int
	}{
		a: [...]int{10, 20, 30, 40, 50},
		b: [...]int{100, 200},
	}

	// capがlen(a)より大きい
	s := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&v.a[1])), Len: 2, Cap: 6}
	var ns []int
	*(*reflect.SliceHeader)(unsafe.Pointer(&ns)) = s

	// aを超えてappend
	ns = append(ns, 60, 70, 80)
	fmt.Println(ns, v.a, v.b)
}
