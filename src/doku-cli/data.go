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

func (p *DoKuData) String() string {
	s := ""

	n := len(p.Data)

	for i := 0; i < n; i++ {
		if i > 0 {
			s += "\n"
		}

		t := p.Data[i]
		l := len(t)

		for j := 0; j < l; j++ {
			if j > 0 {
				s += " "
			}

			s += fmt.Sprint(t[j])
		}
	}

	return s
}
