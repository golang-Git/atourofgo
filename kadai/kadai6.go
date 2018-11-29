package main

import (
	"fmt"
)

/*
🤔もんだい 6: バブルソート

  []int{2, 4, 5, 1, 3} をバブルソートアルゴリズムを実装して、[1 2 3 4 5] の結果を得よう。


　バブルソート
　バブルソートはリストにおいて隣り合うふたつの要素の値を比較して条件に応じた交換を行う整列アルゴリズムです。
　条件とは値の大小関係です。「値の大きい順(降順)」か「値の小さい順(昇順)」にリストを並び替えます。
　このソートを実行すると値の大きいまたは小さい要素が浮かびあがってくるように見えることから、バブル(bubble: 泡)ソートと呼ばれます。

　アルゴリズム (リストを昇順に整列させる手順)

　1. 先頭の要素'A'と隣り合う次の要素'B'の値を比較する
　2. 要素'A'が要素'B'より大きいなら、要素'A'と要素'B'の値を交換する
　3. 先頭の要素を'B'に移し、要素'B'と隣り合う要素'C'の値を比較/交換する
　4. 先頭の要素を'C','D','E'...と移動しながら比較/交換をリストの終端まで繰り返す
　5. 最も大きい値を持つ要素がリストの終端へ浮かびあがる
　6. リストの終端には最も大きな値が入っているので、リストの終端の位置をずらして(要素数をひとつ減らして)手順1〜6を繰り返す

　以上のように総当たりで比較を行い、条件に一致する交換を実行することで整列が完了します。

　図解
　https://www.codereading.com/algo_and_ds/algo/images/bubble-sort.png

*/

func main() {
	//arr := []int{2, 4, 5, 1, 3}
	arr := []int{2, 4, 5, 1, 3, 6, 8, 1, 2, 10, 8, 2, 4}
	//arr := []int{2, 4, 5, 1, 3}

	fmt.Println(BubbleSort(arr))
}

func BubbleSort(arr []int) []int {
	// 😤
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				work := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = work
			}
		}
	}
	return arr
}
