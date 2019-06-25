package interface_demo

import "fmt"

type demo interface {
	Say(string)
}

type person struct {
	id 	int
}

func (p *person) Say(sth string) {
	fmt.Printf("No.%d said，%s\n", p.id, sth)
}

func (p *person) Watch() {
	fmt.Printf("No.%d watch %s\n", p.id, "game")
}

func main() {
	var p demo
	p = &person{
		id: 2,
	}

	p.Say("nihao")
}
