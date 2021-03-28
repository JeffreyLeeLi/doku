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
				fmt.Println("2->4->8->16->32->64->128->256->512->1024->2048->...")
			}
		case `1`:
			{
				populateMapOfAlphabet()
				showLevelDetails()
			}
		case `2`:
			{
				populateMapOfUltraman()
				showLevelDetails()
			}
		case `3`:
			{
				populateMapOfGloryOfKings()
				showLevelDetails()
			}
		case `4`:
			{
				populateMapOfMilitary()
				showLevelDetails()
			}
		case `5`:
			{
				populateMapOfUniverse()
				showLevelDetails()
			}
		case `6`:
			{
				populateMapOfHuman()
				showLevelDetails()
			}
		case `7`:
			{
				populateMapOfWorld()
				showLevelDetails()
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
	fmt.Println(`[N]ew | [W]: Up [S]: Down [A]: Left [D]: Right | [1]: Alphabet [2] Ultraman [3] Glory of Kings [4] Military [5] Universe [6] Human 【7】 World ... [0]: Number | [H]elp [Q]uit`)
}
