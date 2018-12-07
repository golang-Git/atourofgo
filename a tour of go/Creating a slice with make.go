package main

import "./slice"

func main() {
	/*
		配列は他の変数に代入した場合、値を共有する。
		スライスにアペンドを行い、追加した値がcapの値を超えた場合、配列のcapは自動的に倍の値に拡張される
		値が拡張された場合、以降配列同氏は値を共有しなくなる
	*/
	a := make([]int, 5)
	slice.Printslice("a", a)

	b := make([]int, 0, 5)
	slice.Printslice("b", b)

	c := b[:2]
	slice.Printslice("c", c)

	d := c[2:5]
	slice.Printslice("d", d)

}
