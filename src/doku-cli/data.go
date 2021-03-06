package main

import (
	"fmt"
	"math/rand"
	"strings"
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
	col = rand.Intn(p.colCount())

	return
}

func (p *DoKuData) Init() {
	if p.hasEmptyRowOrCol() {
		return
	}

	for i := 0; i < p.rowCount(); i++ {
		for j := 0; j < p.colCount(); j++ {
			p.set(i, j, 0)
		}
	}

	p.generate()
}

func (p *DoKuData) generate() {
	i, j := p.randPos()

	for {
		if p.noneAt(i, j) {
			break
		}

		i, j = p.randPos()
	}

	p.set(i, j, p.Base)
}

func (p *DoKuData) swap(i, j, k, l int) {
	p.Data[i][j], p.Data[k][l] = p.Data[k][l], p.Data[i][j]
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

func (p *DoKuData) colCount() int {
	if p.rowCount() <= 0 {
		return 0
	}

	return len(p.Data[0])
}

func (p *DoKuData) hasEmptyRowOrCol() bool {
	n := p.rowCount()

	if n <= 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if p.colCount() <= 0 {
			return true
		}
	}

	return false
}

func (p *DoKuData) String() string {
	s := ""

	n := p.rowCount()

	for i := 0; i < n; i++ {
		s += fmt.Sprintf("+%s", strings.Repeat("----+", p.colCount()))
		s += "\n"

		l := p.colCount()

		for j := 0; j < l; j++ {
			s += "|"

			if p.noneAt(i, j) {
				s += strings.Repeat(" ", 4)
			} else {
				t := p.at(i, j)

				if v, ok := ValueMap[t]; ok {
					s += fmt.Sprintf("%v", v)
				} else {
					s += fmt.Sprintf("%4d", t)
				}
			}
		}

		s += "|"
		s += "\n"
	}

	s += fmt.Sprintf("+%s", strings.Repeat("----+", p.colCount()))

	return s
}
