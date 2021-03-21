package main

import "fmt"

type DoKuData struct {
	Data [][]int
}

func NewDoKuData(n, l int) *DoKuData {
	data := make([][]int, n)

	for i := 0; i < n; i++ {
		data[i] = make([]int, l)
	}

	return &DoKuData{
		Data: data,
	}
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

func (p *DoKuData) String() string {
	s := ""

	n := p.rowCount()

	for i := 0; i < n; i++ {
		if i > 0 {
			s += "\n"
		}

		t := p.Data[i]
		l := p.colCount(i)

		for j := 0; j < l; j++ {
			if j > 0 {
				s += " "
			}

			s += fmt.Sprint(t[j])
		}
	}

	return s
}
