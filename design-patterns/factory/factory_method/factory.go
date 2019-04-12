package main

import "fmt"

type IBunShop interface {
	Generate(string) IBun
}

type IBun interface {
	Create()
}

type BabiBun struct {

}

func (b *BabiBun) Generate(t string) IBun {
	switch t {
	case "meat":
		return &BabiMeatBun{}
	case "vegetable":
		return &BabiVegetableBun{}
	default:
		return nil
	}
}

type BabiMeatBun struct {

}

func (mb *BabiMeatBun) Create() {
	fmt.Println("create BabiMeatBun")
}

type BabiVegetableBun struct {

}

func (vb *BabiVegetableBun) Create() {
	fmt.Println("create BabiVegetableBun")
}

type YongHeBun struct {

}

func (y *YongHeBun) Generate(t string) IBun {
	switch t {
	case "meat":
		return &YongHeMeatBun{}
	case "vegetable":
		return &YongHeVegetableBun{}
	default:
		return nil
	}
}

type YongHeMeatBun struct {

}

func (mb *YongHeMeatBun) Create() {
	fmt.Println("create YongHeMeatBun")
}

type YongHeVegetableBun struct {

}

func (vb *YongHeVegetableBun) Create() {
	fmt.Println("create YongHeVegetableBun")
}
