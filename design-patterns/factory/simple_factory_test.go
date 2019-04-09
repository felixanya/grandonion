package factory

import "testing"

func TestShapeFactory_GetShape(t *testing.T) {
	factory := ShapeFactory{}
	factory.GetShape("CIRCLE").Draw()
	factory.GetShape("RECTANGLE").Draw()
	factory.GetShape("SQUARE").Draw()
}

func TestNewShape(t *testing.T) {
	shape := NewShape("CIRCLE")
	shape.Draw()
}
