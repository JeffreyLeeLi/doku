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

func (p *DoKuData) Up() {
	swapped := false

	for j := 0; j < p.colCount(); j++ {
		for i := 0; i < p.rowCount(); i++ {
			if p.noneAt(i, j) {
				continue
			}

			for k := i; k > 0; k-- {
				if p.noneAt(k-1, j) {
					p.swap(k, j, k-1, j)
					swapped = true
				}
			}
		}
	}

	combined := false

	for j := 0; j < p.colCount(); j++ {
		for i := 0; i < p.rowCount()-1; i++ {
			if p.noneAt(i, j) || p.noneAt(i+1, j) {
				continue
			}

			one := p.at(i, j)
			other := p.at(i+1, j)

			if one == other {
				p.set(i, j, one+other)
				p.set(i+1, j, 0)

				for k := i + 1; k < p.rowCount()-1; k++ {
					p.swap(k, j, k+1, j)
				}

				combined = true
			}
		}
	}

	if swapped || combined {
		p.generate()
	}
}

func (p *DoKuData) Down() {
}

func (p *DoKuData) Left() {
}

func (p *DoKuData) Right() {
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
		if i > 0 {
			s += "\n"
		}

		l := p.colCount()

		for j := 0; j < l; j++ {
			if j > 0 {
				s += "|"
			}

			if p.noneAt(i, j) {
				s += strings.Repeat(" ", 4)
			} else {
				s += fmt.Sprintf("%4d", p.at(i, j))
			}
		}
	}

	return s
}
