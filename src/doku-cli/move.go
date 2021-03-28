package main

import (
	"fmt"
)

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
	} else {
		fmt.Println("Nothing Changed")
	}
}

func (p *DoKuData) Down() {
	swapped := false

	for j := 0; j < p.colCount(); j++ {
		for i := p.rowCount() - 1; i >= 0; i-- {
			if p.noneAt(i, j) {
				continue
			}

			for k := i; k < p.rowCount()-1; k++ {
				if p.noneAt(k+1, j) {
					p.swap(k, j, k+1, j)
					swapped = true
				}
			}
		}
	}

	combined := false

	for j := 0; j < p.colCount(); j++ {
		for i := p.rowCount() - 1; i > 0; i-- {
			if p.noneAt(i, j) || p.noneAt(i-1, j) {
				continue
			}

			one := p.at(i, j)
			other := p.at(i-1, j)

			if one == other {
				p.set(i, j, one+other)
				p.set(i-1, j, 0)

				for k := i - 1; k > 0; k-- {
					p.swap(k, j, k-1, j)
				}

				combined = true
			}
		}
	}

	if swapped || combined {
		p.generate()
	} else {
		fmt.Println("Nothing Changed")
	}
}

func (p *DoKuData) Left() {
	swapped := false

	for i := 0; i < p.rowCount(); i++ {
		for j := 0; j < p.colCount(); j++ {
			if p.noneAt(i, j) {
				continue
			}

			for l := j; l > 0; l-- {
				if p.noneAt(i, l-1) {
					p.swap(i, l, i, l-1)
					swapped = true
				}
			}
		}
	}

	combined := false

	for i := 0; i < p.rowCount(); i++ {
		for j := 0; j < p.colCount()-1; j++ {
			if p.noneAt(i, j) || p.noneAt(i, j+1) {
				continue
			}

			one := p.at(i, j)
			other := p.at(i, j+1)

			if one == other {
				p.set(i, j, one+other)
				p.set(i, j+1, 0)

				for l := j + 1; l < p.colCount()-1; l++ {
					p.swap(i, l, i, l+1)
				}

				combined = true
			}
		}
	}

	if swapped || combined {
		p.generate()
	} else {
		fmt.Println("Nothing Changed")
	}
}

func (p *DoKuData) Right() {
	swapped := false

	for i := 0; i < p.rowCount(); i++ {
		for j := p.colCount() - 1; j >= 0; j-- {
			if p.noneAt(i, j) {
				continue
			}

			for l := j; l < p.colCount()-1; l++ {
				if p.noneAt(i, l+1) {
					p.swap(i, l, i, l+1)
					swapped = true
				}
			}
		}
	}

	combined := false

	for i := 0; i < p.rowCount(); i++ {
		for j := p.colCount() - 1; j > 0; j-- {
			if p.noneAt(i, j) || p.noneAt(i, j-1) {
				continue
			}

			one := p.at(i, j)
			other := p.at(i, j-1)

			if one == other {
				p.set(i, j, one+other)
				p.set(i, j-1, 0)

				for l := j - 1; l > 0; l-- {
					p.swap(i, l, i, l-1)
				}

				combined = true
			}
		}
	}

	if swapped || combined {
		p.generate()
	} else {
		fmt.Println("Nothing Changed")
	}
}
