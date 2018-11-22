package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int64 {
	/*
		関数の中で宣言された変数はローカル変数といい、呼び出しを終了する際に破棄される
		しかし、クロージャから呼び出されている変数はローカル変数であっても、破棄されずに保持される
		関数の中で呼び出されるローカル変数がすべて捕捉されるわけではなく、クロージャの中で参照されている変数のみが捕捉される
		goにはジェネレータ（自身の内部の状態を保持し、呼び出されるたびに現在の状態から導き出される処理結果を返す）の機能はないが、クロージャを用いることでジェネレータの役割を実現できる
		※クロージャを別に新しく生成した際は、現在の内部の状態は共有されずに新しく生成される
	*/
	var s int64
	var next int64
	//var b bool
	var work int64
	//var work2 int64
	return func() int64 {
		/*
			if (s + next) == 0 {
				next = 1
			} else if (s + next) == 1 {
				s = next
			} else {
				if b == true {
					if work == 0 {
						next = next + s
					} else {
						next = work
					}
					work = next + s
					s = next
				} else {
					b = true
				}
			}
		*/
		work = s     //ワークスペースに現在のフィボナッチ数の値を退避
		s = s + next //現在のフィボナッチ数にn-1番目の数値を足す
		next = work  //n-1番目のフィボナッチ数を次に足す値として代入
		if s == 0 {
			next += 1
			return s
		} else if s == 1 {
			return s
		} else {
			s -= 1
			fibonacci()
			s -= 1
			fibonacci()
			s += 2
		}
		return s
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
