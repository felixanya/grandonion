package common

type A interface {
	Minus(s string) string
}

type B interface {
	Add(s string) string
}

type C interface {
	Deadline() string

	Value(key int) int
}

type D interface {
	SetKV(k, v string)
	GetValue(k string) string
}