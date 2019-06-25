package dpkg

type DImpl struct {
	Key 	string
	Value 	string
}

func NewD() *DImpl {
	return new(DImpl)
}

func (d *DImpl) SetKV(k, v string) {
	d.Key = k
	d.Value = v
}

func (d *DImpl) GetValue(k string) string {
	return d.Value
}
