package main

import (
	"fmt"
	"strings"
)

func showLevelDetails() {
	n := len(ValueMap)

	fmt.Println(strings.Repeat("+----+  ", n))

	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Print("->")
		}

		fmt.Printf("|%s|", ValueMap[1<<uint(i)])
	}

	fmt.Println()

	fmt.Println(strings.Repeat("+----+  ", n))
}

func emptyValueMap() {
	ValueMap = map[int]string{}
}

func populateMapOfAlphabet() {
	emptyValueMap()

	ValueMap[1<<1] = "   A"
	ValueMap[1<<2] = "   B"
	ValueMap[1<<3] = "   C"
	ValueMap[1<<4] = "   D"
	ValueMap[1<<5] = "   E"
	ValueMap[1<<6] = "   F"
	ValueMap[1<<7] = "   G"
	ValueMap[1<<8] = "   H"
	ValueMap[1<<9] = "   I"

	ValueMap[1<<10] = "   J"
	ValueMap[1<<11] = "   K"
	ValueMap[1<<12] = "   L"
	ValueMap[1<<13] = "   M"
	ValueMap[1<<14] = "   N"
	ValueMap[1<<15] = "   O"
	ValueMap[1<<16] = "   P"
	ValueMap[1<<17] = "   Q"
	ValueMap[1<<18] = "   R"
	ValueMap[1<<19] = "   S"
	ValueMap[1<<20] = "   T"
	ValueMap[1<<21] = "   U"
	ValueMap[1<<22] = "   V"
	ValueMap[1<<23] = "   W"
	ValueMap[1<<24] = "   X"
	ValueMap[1<<25] = "   Y"
	ValueMap[1<<26] = "   Z"
}

func populateMapOfUltraman() {
	emptyValueMap()

	ValueMap[1<<1] = "初代"
	ValueMap[1<<2] = "赛文"
	ValueMap[1<<3] = "杰克"
	ValueMap[1<<4] = "艾斯"
}

func populateMapOfGloryOfKings() {
	emptyValueMap()

	ValueMap[1<<1] = "铁剑"
	ValueMap[1<<2] = "陨星"
	ValueMap[1<<3] = "战斧"
	ValueMap[1<<4] = "星锤"
	ValueMap[1<<5] = "破军"
}

func populateMapOfMilitary() {
	emptyValueMap()

	ValueMap[1<<1] = " 兵 "
	ValueMap[1<<2] = " 班 "
	ValueMap[1<<3] = " 排 "
	ValueMap[1<<4] = " 连 "
	ValueMap[1<<5] = " 营 "
	ValueMap[1<<6] = " 团 "
	ValueMap[1<<7] = " 旅 "
	ValueMap[1<<8] = " 师 "
	ValueMap[1<<9] = " 军 "
	ValueMap[1<<10] = "集团"
	ValueMap[1<<11] = "统帅"
}

func populateMapOfUniverse() {
	emptyValueMap()

	ValueMap[1<<1] = "核子"
	ValueMap[1<<2] = "原子"
	ValueMap[1<<3] = "分子"
	ValueMap[1<<4] = "石子"
	ValueMap[1<<5] = "流星"
	ValueMap[1<<6] = "卫星"
	ValueMap[1<<7] = "行星"
	ValueMap[1<<8] = "恒星"
	ValueMap[1<<9] = "星系"
	ValueMap[1<<10] = "黑洞"
}

func populateMapOfHuman() {
	emptyValueMap()

	ValueMap[1<<1] = "细胞"
	ValueMap[1<<2] = "组织"
	ValueMap[1<<3] = "器官"
	ValueMap[1<<4] = "系统"
	ValueMap[1<<5] = "人体"
}

func populateMapOfWorld() {
	emptyValueMap()

	ValueMap[1<<1] = " 人 "
	ValueMap[1<<2] = " 家 "
	ValueMap[1<<3] = " 村 "
	ValueMap[1<<4] = " 乡 "
	ValueMap[1<<5] = " 县 "
	ValueMap[1<<6] = " 市 "
	ValueMap[1<<7] = " 省 "
	ValueMap[1<<8] = " 国 "
	ValueMap[1<<9] = "世界"
}
