package main

import (
	"fmt"
	"strings"
)

func main() {
	d := NewDoKuData(4, 4, 2)

	cmd := `h`
	for {
		prevCmd := cmd

		fmt.Printf(`[%s] `, cmd)

		fmt.Scanln(&cmd)

		switch strings.ToLower(cmd) {
		case `n`, `new`:
			{
				d.Init()
				fmt.Println(d)
			}
		case `w`, `up`:
			{
				d.Up()
				fmt.Println(d)
			}
		case `s`, `down`:
			{
				d.Down()
				fmt.Println(d)
			}
		case `a`, `left`:
			{
				d.Left()
				fmt.Println(d)
			}
		case `d`, `right`:
			{
				d.Right()
				fmt.Println(d)
			}
		case `h`, `help`:
			{
				displayKeymap()
			}
		case `q`, `quit`:
			{
				fmt.Println("Seeya~")
				return
			}
		default:
			{
				fmt.Println("Invalid Command:", cmd)
				cmd = prevCmd
			}
		}
	}
}

func displayKeymap() {
	fmt.Println(`[N]ew | W: up S: down A: left D: right | [H]elp [Q]uit`)
}
