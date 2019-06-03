package main

import "testing"

//func TestBabiBun_Generate(t *testing.T) {
//	var b IBun
//
//	bb := new(BabiBunShop)
//	b = bb.Generate("meat")
//	b.Create()
//
//	b2 := bb.Generate("vegetable")
//	b2.Create()
//
//	yh := new(YongheBunShop)
//	b3  := yh.Generate("meat")
//	b3.Create()
//
//	b4 := yh.Generate("vegetable")
//	b4.Create()
//}

func TestNewBunShop(t *testing.T) {
	bs1 := NewBunShop("Babi")
	bs1.Generate("meat").Create()

	bs2 := NewBunShop("Yonghe")
	bs2.Generate("vegetable").Create()
}
