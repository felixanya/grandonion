package main

import "fmt"

func MapData() []map[string]string {
	mpList := make([]map[string]string, 0)
	mp := make(map[string]string)
	mp["ni"] = "you"
	mp["hao"] = "good"
	mp["shijie"] = "world"

	mp2 := make(map[string]string)
	mp2["ni"] = "you2"
	mp2["hao"] = "good2"
	mp2["shijie"] = "world2"

	mpList = append(mpList, mp)
	mpList = append(mpList, mp2)

	for _, mpd := range mpList {
		delete(mpd, "hao")
		mpd["say"] = "hello"
	}

	return mpList
}

func main() {
	//ml := MapData()
	//fmt.Println(ml)

	mt := make(map[string]interface{})
	mt["ni"] = "you"
	mt["no"] = 100

	for k, v := range mt {
		fmt.Println(k)
		switch v.(type) {
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("unknown")
		}
	}
}
