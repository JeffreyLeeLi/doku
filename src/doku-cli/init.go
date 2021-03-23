package main

import (
	"math/rand"
	"time"
)

var ValueMap map[int]string

func init() {
	rand.Seed(time.Now().Unix())

	ValueMap = make(map[int]string)

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
