package main

import "testing"

func TestBabiBun_Generate(t *testing.T) {
	var b IBun

	bb := new(BabiBun)
	b = bb.Generate("meat")
	b.Create()

	b2 := bb.Generate("vegetable")
	b2.Create()

	yh := new(YongHeBun)
	b3  := yh.Generate("meat")
	b3.Create()

	b4 := yh.Generate("vegetable")
	b4.Create()
}
