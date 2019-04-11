package builder

type IBuilder interface {
	SetName(string)
	SetColor(string)
	Build()
}

