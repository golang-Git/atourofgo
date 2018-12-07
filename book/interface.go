package main

import (
	"fmt"
)

type stringify interface {
	ToString() string
}

type person struct {
	Name        string
	Sex         string
	Bloodtype   string
	Phonenumber string
	Address     string
}

func (p *person) ToString() string {
	return fmt.Sprintf("name=%v sex=%v bloodtype=%v phonenumber=%v address=%v", p.Name, p.Sex, p.Bloodtype, p.Phonenumber, p.Address)
}

type info struct {
	Ipaddress string
	Password  string
}

func (i *info) ToString() string {
	return fmt.Sprintf("ipaddress=%v password=%v", i.Ipaddress, i.Password)
}

func Print(s stringify) {
	fmt.Print(s.ToString())
}

func main() {
	str := []stringify{
		&person{Name: "新田 美砂", Sex: "女", Bloodtype: "A", Phonenumber: "09028948024", Address: "石川県羽咋市神子原町6-3-3"},
		&info{Ipaddress: "248.15.162.182", Password: "BXLDyADB"},
	}
	for _, v := range str {
		Print(v)
	}
}
