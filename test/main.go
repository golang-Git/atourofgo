package main

func main() {
	/*
		同フォルダのapp.goで定義した関数と変数。
		go runするときは、main.goと一緒にapp.goも実行しないとapp.goは読み込まれない
		また、インポートは独立しているので、そのファイル内でしか利用できない
	*/
	printMessage(Pi)

}
