//Pointers

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //&をつけることでiのポインタを代入できる（iと同じ場所を参照する）
	fmt.Println(*p) //参照したものは*をつけることでそのポインタの先の値を参照できる
	*p = 21         //pはiと同じ場所を参照するので*pに代入をするとiの値も変更する
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
