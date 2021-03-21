package main

import "fmt"

func main() {
	d := NewDoKuData(4, 5, 2)
	d.Init()
	fmt.Println(d)
}
