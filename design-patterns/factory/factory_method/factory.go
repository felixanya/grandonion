package main

import "fmt"

type IBunShop interface {
	Generate(string) IBun
}

func NewBunShop(name string) IBunShop {
	switch name {
	case "Babi":
		return &BabiBunShop{}
	case "Yonghe":
		return &YongheBunShop{}
	default:
		return nil
	}
}

type IBun interface {
	Create()
}

type BabiBunShop struct {

}

func (b *BabiBunShop) Generate(t string) IBun {
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

type YongheBunShop struct {

}

func (y *YongheBunShop) Generate(t string) IBun {
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
