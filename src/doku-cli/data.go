package main

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
