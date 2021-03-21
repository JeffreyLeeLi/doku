package main

import (
	"fmt"
	"strings"
)

func main() {
	d := NewDoKuData(4, 5, 2)

	cmd := `h`
	for {
		fmt.Printf(`[%s] `, cmd)

		fmt.Scanln(&cmd)

		switch strings.ToLower(cmd) {
		case `n`, `new`:
			{
				d.Init()
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
			}
		}
	}
}

func displayKeymap() {
	fmt.Println(`[N]ew | W: up S: down A: left D: right | [H]elp [Q]uit`)
}
