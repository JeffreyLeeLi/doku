package main

import (
	"math/rand"
	"time"
)

var ValueMap map[int]string

func init() {
	rand.Seed(time.Now().Unix())

	ValueMap = make(map[int]string)
}
