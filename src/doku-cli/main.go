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
		case `0`:
			{
				emptyValueMap()
			}
		case `1`:
			{
				populateMapOfAlphabet()
			}
		case `2`:
			{
				populateMapOfUltraman()
			}
		case `3`:
			{
				populateMapOfGloryOfKings()
			}
		case `4`:
			{
				populateMapOfMilitary()
			}
		case `5`:
			{
				populateMapOfUniverse()
			}
		case `6`:
			{
				populateMapOfHuman()
			}
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
	fmt.Println(`[N]ew | [W]: Up [S]: Down [A]: Left [D]: Right | [0]: Number [1]: Alphabet [2] Ultraman [3] Glory of Kings [4] Military [5] Universe [6] Human | [H]elp [Q]uit`)
}
