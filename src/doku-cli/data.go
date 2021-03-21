package main

import (
	"fmt"
	"math/rand"
)

type DoKuData struct {
	Data [][]int
	Base int
}

func NewDoKuData(n, l, b int) *DoKuData {
	data := make([][]int, n)

	for i := 0; i < n; i++ {
		data[i] = make([]int, l)
	}

	return &DoKuData{
		Data: data,
		Base: b,
	}
}

func (p *DoKuData) randPos() (row, col int) {
	row = rand.Intn(p.rowCount())
	col = rand.Intn(p.colCount(0))

	return
}

func (p *DoKuData) Init() {
	if p.hasEmptyRowOrCol() {
		return
	}

	p.Generate()
}

func (p *DoKuData) Generate() {
	i, j := p.randPos()

	for {
		if p.noneAt(i, j) {
			break
		}

		i, j = p.randPos()
	}

	p.set(i, j, p.Base)
}

func (p *DoKuData) noneAt(i, j int) bool {
	return (p.at(i, j) == 0)
}

func (p *DoKuData) at(i, j int) int {
	return p.Data[i][j]
}

func (p *DoKuData) set(i, j, v int) {
	p.Data[i][j] = v
}

func (p *DoKuData) rowCount() int {
	return len(p.Data)
}

func (p *DoKuData) colCount(i int) int {
	if p.rowCount() <= 0 {
		return 0
	}

	return len(p.Data[i])
}

func (p *DoKuData) hasEmptyRowOrCol() bool {
	n := p.rowCount()

	if n <= 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if p.colCount(0) <= 0 {
			return true
		}
	}

	return false
}

func (p *DoKuData) String() string {
	s := ""

	n := p.rowCount()

	for i := 0; i < n; i++ {
		if i > 0 {
			s += "\n"
		}

		l := p.colCount(i)

		for j := 0; j < l; j++ {
			if j > 0 {
				s += " "
			}

			s += fmt.Sprint(p.at(i, j))
		}
	}

	return s
}
