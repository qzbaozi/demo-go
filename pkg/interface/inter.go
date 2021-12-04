package main

import "fmt"

type IBase interface {
	get() int
}

type Base struct {
	side int
}

func (b *Base) get() int {
	return b.side
}

func (b *Base) add(number int) int {
	b.side = number + b.side
	return b.side
}

func main() {
	b := new(Base)
	b.side = 1

	var baseIntf IBase
	baseIntf = b

	fmt.Println(baseIntf.get())
	fmt.Println(b.add(2))
	fmt.Println(b.get())
}
