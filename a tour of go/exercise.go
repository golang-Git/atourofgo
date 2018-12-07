package main

import "fmt"

type Teacher struct {
	Name string
}

type Student struct {
	Name   string
	Number int
	Grade  int
}

type Person interface {
	getAddress() string
}

func main() {
	s := Student{
		Name:   "example",
		Number: 1,
		Grade:  5,
	}

	t := Teacher{
		Name: "teacher",
	}

	fmt.Println(sendMail(s))
	fmt.Println(sendMail(t))
}

func (s Student) getAddress() string {
	return s.Name + "@student.ad.jp"
}

func (t Teacher) getAddress() string {
	return t.Name + "@teacher.ad.jp"
}

func sendMail(p Person) (context string) {
	from := p.getAddress()
	context = `
	送信元 ： ` + from + `
	これはテスト用のメールです。
	宜しくお願いします。
	`

	return context
}
