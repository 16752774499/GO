package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string   `json:"title"` //序列化时，将字段名替换成title
	Year  int      `json:"year"`
	Price int      `json:"rmb"` //序列化时，将字段名替换成rmb
	Actor []string `json:"actor"`
}

func main() {
	movie := Movie{
		Title: "战狼2",
		Year:  2016,
		Price: 98,
		Actor: []string{"刘德华", "张学友"},
	}
	//序列化 (结构体-----》json)
	data, err := json.Marshal(movie)
	if err != nil {
		fmt.Printf("序列化失败，错误：%v\n", err)
		return
	} else {
		fmt.Printf("序列化成功，数据：%s\n", data)
	}

	//反序列化 (json-----》结构体)
	var Mymovie Movie //用于接收反序列化后的数据
	err = json.Unmarshal(data, &Mymovie)
	if err != nil {
		fmt.Printf("反序列化失败，错误：%v\n", err)
		return
	} else {
		fmt.Printf("反序列化成功，数据：%+v\n", Mymovie)
	}

}
