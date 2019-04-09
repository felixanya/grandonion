package factory

import "fmt"

type Shape interface {
	Draw()
}

// 根据传入参数生产一个shape对象
func NewShape(shape string) Shape {
	switch shape {
	case "":
		return nil
	case "CIRCLE":
		return &Circle{}
	case "RECTANGLE":
		return &Rectangle{}
	case "SQUARE":
		return &Square{}
	default:
		// shape type error
		return nil
	}
}

//------- Rectangle
type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}

//------- Square
type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("Inside Square  ::draw() method.")
}

//------- Circle
type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Inside Circle  ::draw() method.")
}

//------- factory
// 也可以定义一个Shape Factory结构体，通过结构体方法来生产shape对象
// 上面的NewShape少写一些代码，这种方法可以在factory上再进行一些自定义
type ShapeFactory struct {

}

func (sf ShapeFactory) GetShape(shape string) Shape {
	switch shape {
	case "":
		return nil
	case "CIRCLE":
		return &Circle{}
	case "RECTANGLE":
		return &Rectangle{}
	case "SQUARE":
		return &Square{}
	default:
		// shape type error
		return nil
	}
}
