package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	row := make([]uint8, dx)
	for i := range res {
		for j := range row {
			row[j] = uint8((i + j) / 2)
		}
		res[i] = row
	}
	return res
}

func main() {
	pic.Show(Pic)
}
