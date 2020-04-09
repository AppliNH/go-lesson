package oop

import "fmt"

type socket interface {
	plug(d device)
}

type elecSocket []socket

type device struct {
	name        string
	powerNeeded int
	emoji       string
}

func (s elecSocket) plug(d device) {
	fmt.Println("Plugging", d.name, "...")
	d.draw()
}

func (d device) draw() {
	fmt.Println(d.name, "has been plugged ! Power is :", d.powerNeeded)
	fmt.Println(d.emoji)

}

func ExecLesson3() {

	var mySocket elecSocket

	blender := device{name: "Blender", powerNeeded: 30, emoji: `🥛`}
	coffeeMaker := device{name: "Coffee Maker", powerNeeded: 15, emoji: `☕`}
	hifiSystem := device{name: "Hi-fi System", powerNeeded: 40, emoji: `🎶`}

	mySocket.plug(blender)
	mySocket.plug(coffeeMaker)
	mySocket.plug(hifiSystem)


}
