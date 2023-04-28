package main

import "fmt"

type Video struct {
	id int32
	title string
	url string
	thumbnail string
}

type Videos struct {
	videos []Video
}

func main () {
	Url := "https://youtube.com/"
	fmt.Println(Url)
}
