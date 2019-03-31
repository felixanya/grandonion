package bpkg

//import "github.com/aaronize/grandonion/examples/cycle_import/apkg"
//
//func Add(bs string) string {
//	return bs + "---bpak.Add---"
//}
//
//func Goo(bs string) string {
//	return apkg.Minus(bs)
//}

type bt struct {
	key 	string
	value 	string
}

func NewBt(k, v string) *bt {
	return &bt{key: k, value: v}
}

func (b *bt) GetKey() string {
	return b.key
}
func (b *bt) SetKey(k string) {
	b.key = k
}

func (b *bt) GetVal() string {
	return b.value
}
func (b *bt) SetVal(v string) {
	b.value = v
}
