package main

import (
	"fmt"
)

type MyCar struct {
	name  string
	speed int
}

type Car interface {
	run(int) interface{}
	stop()
}

func (c *MyCar) run(speed int) interface{} {
	c.speed = speed
	return fmt.Sprintf("時速%v kmで走ります。\n", c.speed)
}

func (c *MyCar) stop() {
	c.speed = 0
	fmt.Println("停止します。")
}

func main() {
	car := &MyCar{name: "GTR", speed: 120}
	//fmt.Print(car.run(car.speed))
	var str string
	str, _ = car.run(car.speed).(string)
	fmt.Print(str)
	car.stop()
}
