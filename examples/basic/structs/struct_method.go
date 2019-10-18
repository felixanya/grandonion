package structs

type Computer struct {
    Brand       string
    Model       string
    ProductNo   string
    Remark      string
}

func NewComputer() *Computer {
    return nil
}

// receiver类型为指针的方法。非*Computer类型的变量不可以调用
func (c *Computer) AddRemark(remark string) error {

    return nil
}

// receiver类型为非指针的方法。
func (c Computer) GetProductNo() string {
    return "GetProductNo"
}
